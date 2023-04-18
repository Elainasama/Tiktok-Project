package common

import (
	"errors"
	"github.com/minio/minio-go/v6"
	"log"
	"strconv"
	"strings"
	"time"
)

type MinIO struct {
	Client      *minio.Client
	Endpoint    string
	Port        string
	VideoBucket string
	ImageBucket string
}

var m MinIO

func GetMinIO() *MinIO {
	return &m
}

func InitMinIO() {
	conf := GetConfig()
	ip := conf.Minio.Host
	port := conf.Minio.Port

	endpoint := ip + ":" + port

	accessKeyId := conf.Minio.AccessKeyID
	secretAccessKey := conf.Minio.SecretAccessKey
	useSSL := false

	minioClient, err := minio.New(endpoint, accessKeyId, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln("Create MinIO Failed", err)
		return
	}
	CreateBucket(minioClient, conf.Minio.VideoBucket)
	CreateBucket(minioClient, conf.Minio.imageBucket)
	m = MinIO{
		Client:      minioClient,
		Endpoint:    endpoint,
		Port:        port,
		VideoBucket: conf.Minio.VideoBucket,
		ImageBucket: conf.Minio.imageBucket,
	}
}

func CreateBucket(m *minio.Client, bucketName string) {
	found, err := m.BucketExists(bucketName)
	if err != nil {
		log.Println("Check BucketName Exist Error")
	}
	if !found {
		err = m.MakeBucket(bucketName, "guangdong")
		if err != nil {
			log.Fatalln("Create BucketName Failed")
		}
	}
	//设置桶策略
	policy := `{"Version": "2012-10-17",
				"Statement": 
					[{
						"Action":["s3:GetObject"],
						"Effect": "Allow",
						"Principal": {"AWS": ["*"]},
						"Resource": ["arn:aws:s3:::` + bucketName + `/*"],
						"Sid": ""
					}]
				}`
	err = m.SetBucketPolicy(bucketName, policy)
	if err != nil {
		log.Printf("SetBucketPolicy %s  err:%s\n", bucketName, err.Error())
	}
}

// UploadFile 时间戳命名
// 命名规则 userid + '_' + current time + suffix
func (m *MinIO) UploadFile(file, userid, bucket string) (string, error) {
	var filename strings.Builder
	var contentType, suffix string
	if bucket == "video" {
		contentType = "video/mp4"
		suffix = ".mp4"
	} else if bucket == "image" {
		contentType = "image/jpeg"
		suffix = ".jpg"
	} else {
		return "", errors.New("不支持的上传类型")
	}

	filename.WriteString(userid)
	filename.WriteString("_")
	filename.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	filename.WriteString(suffix)

	n, err := m.Client.FPutObject(bucket, filename.String(), file, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Println("Upload File Failed", err)
		return "", err
	}
	log.Printf("Upload File %d bytes Successfully,filename : %s\n", n, filename.String())
	url := "http://" + m.Endpoint + "/" + bucket + "/" + filename.String()
	return url, nil
}

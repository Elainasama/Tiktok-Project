package common

import (
	"github.com/spf13/viper"
	"log"
)

type Configs struct {
	Minio MinioConfig
	Path  Path
}

type MinioConfig struct {
	Host            string
	Port            string
	AccessKeyID     string
	SecretAccessKey string
	VideoBucket     string
	imageBucket     string
}
type Path struct {
	VideoPath string
	ImagePath string
}

var Config Configs

func InitConfig() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	minio := MinioConfig{
		Host:            viper.GetString("minio.host"),
		Port:            viper.GetString("minio.port"),
		AccessKeyID:     viper.GetString("minio.accessKeyID"),
		SecretAccessKey: viper.GetString("minio.secretAccessKey"),
		VideoBucket:     viper.GetString("minio.videoBucket"),
		imageBucket:     viper.GetString("minio.imageBucket"),
	}
	path := Path{
		VideoPath: viper.GetString("path.videoPath"),
		ImagePath: viper.GetString("path.imagePath"),
	}
	Config = Configs{
		Minio: minio,
		Path:  path,
	}
}

func GetConfig() *Configs {
	return &Config
}

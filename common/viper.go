package common

import (
	"github.com/spf13/viper"
	"log"
)

type Configs struct {
	Minio MinioConfig
	Path  Path
	Redis RedisConfig
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
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
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
	redis := RedisConfig{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}
	Config = Configs{
		Minio: minio,
		Path:  path,
		Redis: redis,
	}
}

func GetConfig() *Configs {
	return &Config
}

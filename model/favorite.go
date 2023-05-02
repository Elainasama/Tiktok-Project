package model

type Favorite struct {
	FavoriteId int64 `gorm:"primary_key;type:bigint(20) auto_increment;not null;column:favorite_id"`
	UserId     int64 `gorm:"not null;column:user_id"`
	VideoId    int64 `gorm:"not null;column:video_id"`
}

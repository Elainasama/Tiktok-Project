package model

// User
// 重构todo
type User struct {
	UserId          int64  `gorm:"primary_key;type:bigint(20) auto_increment;not null;column:user_id"`
	UserName        string `gorm:"not null;column:user_name"`
	Password        string `gorm:"not null;type:varchar(500);column:password"`
	FollowCount     int64  `gorm:"column:follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count"`
	TotalFavorite   int64  `gorm:"column:total_favorite"`
	FavoritedCount  int64  `gorm:"column:favorited_count"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
}

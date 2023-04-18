package model

type Message struct {
	MessageId  int64  `gorm:"primary_key;type:bigint(20) auto_increment;not null;column:message_id"`
	ToUserId   int64  `gorm:"not null;column:to_user_id"`
	FromUserId int64  `gorm:"column:from_user_id;"`
	Content    string `gorm:"column:content;"`
	CreatTime  int64  `gorm:"column:creat_time;"`
}

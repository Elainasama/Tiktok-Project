package model

type Relation struct {
	RelationId int64 `gorm:"primary_key;type:bigint(20) auto_increment;column:relation_id;not null"`
	// 关注的人的id
	FansId int64 `gorm:"not null;column:fans_id"`
	// 被关注者的id
	ToUserId int64 `gorm:"not null;column:to_user_id"`
}

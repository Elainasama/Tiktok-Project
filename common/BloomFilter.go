package common

import (
	"TikTokLite/model"
	"github.com/steakknife/bloomfilter"
	"hash"
	"hash/fnv"
	"strconv"
)

// 重构点 布隆过滤器
// 用比较小的代价防止redis内存击穿

const (
	// 位图长度
	maxElements = 1000000
	// 可接受的错误率
	probCollide = 0.01
)

var bf *bloomfilter.Filter

func InitFilter() {
	var err error
	bf, err = bloomfilter.NewOptimal(maxElements, probCollide)
	if err != nil {
		panic(err)
	}
	addToFilter()
}

func GetFilter() *bloomfilter.Filter {
	return bf
}

func ConvertHash64(data string) hash.Hash64 {
	h := fnv.New64()
	h.Write([]byte(data))
	return h
}

// addToFilter 重新启动时把数据库所有id都加入到过滤器中
func addToFilter() {
	db := GetDB()
	var users []model.User
	err := db.Table("users").Select("user_id").Find(&users).Error
	if err != nil {
		panic(err)
	}
	for _, v := range users {
		bf.Add(ConvertHash64(strconv.FormatInt(v.UserId, 10)))
	}
}

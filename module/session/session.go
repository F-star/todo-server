package session

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"
	"todo/redisapp"
)

var prefix = "todo:session:"

func createRandomSID() (sid string) {
	ts := time.Now().Unix()
	tsStr := strconv.FormatInt(ts, 10)

	src := rand.NewSource(ts + 42)
	r := rand.New(src)
	randHex := strconv.FormatInt(r.Int63(), 16)
	return tsStr + randHex
}

func NewSession(uid string) (sid string) {
	// return "session:" + uid
	ctx := context.Background()
	sid = createRandomSID()
	err := redisapp.Rdb.Set(ctx, prefix+sid, uid, time.Hour).Err() // 默认 1 小时
	if err != nil {
		log.Println(err)
	}
	return sid
}

func getUIDBySID(sid string) string {
	ctx := context.Background()
	uid, err := redisapp.Rdb.Get(ctx, prefix+sid).Result()
	if err != nil {
		// TODO: err handle
		log.Println(err)
	}
	log.Println("获取的uid为（检测如果找不到是否返回空字符串）：", uid)
	return uid
}

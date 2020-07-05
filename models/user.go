package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"todo/config"
)

type User struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Name      string         `json:"name" gorm:"type:char(24);not null"`
	Avator    sql.NullString `json:"avator" gorm:"type:char(50)"`
	Password  string         `json:"password" gorm:"type:char(40)"`     // sha1 password: 40
	Salt      string         `json:"salt" gorm:"type:char(8);not null"` // password salt，8 char
	CreatedAt time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (b *User) TableName() string {
	return "user"
}

func GetUserInfoById() {

}

func getRandStr(size int) (salt string) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	randInt := r.Int()
	salt = strconv.FormatInt(int64(randInt), 16)
	salt = (salt + strings.Repeat("0", size-1))[:size]
	return salt
}

func getHashWithPwdAndSalt(password string, salt string) string {
	mid := len(salt) / 2
	s := sha1.New()
	s.Write([]byte(salt[:mid] + password + salt[mid:]))
	h := s.Sum(nil)
	return fmt.Sprintf("%x", h)
}

func CreateUser(name string, password string) error {
	// hash password + salt
	salt := getRandStr(8)
	password = getHashWithPwdAndSalt(password, salt)

	user := User{
		Name:     name,
		Password: password,
		Salt:     salt,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

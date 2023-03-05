package model

import (
	"dyx_xy/server/common"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name" gorm:"uniqueIndex;size:256"`
	Password  string `json:"password"`
	Authority int    `json:"authority"` // 权限，999代表管理员用户

	SecretKey string `gorm:"-" json:"secret_key"`
}

// BeforeUpdate : hook before a user is updated
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("before update")
	fmt.Println(u.Password)

	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		tx.Statement.SetColumn("Password", hash)
	}

	fmt.Println("after update")
	fmt.Println(u.Password)
	return
}

func (u *User) Verify() (bool, error) {
	u2 := User{Name: u.Name}
	tx := DB.Find(&u2)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, tx.Error
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(u2.Password), []byte(u.Password))
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (u *User) Rigerster() error {
	if u.Authority == 0 {
		u.Authority = 1
	}

	if u.Name == "" || u.Password == "" {
		return fmt.Errorf("the name or password format is incorrect")
	}

	if u.Authority != 1 && u.SecretKey != common.TryGetConfig("secret_key", "209_secret_key_test") {
		return fmt.Errorf("failed to verify secret key")
	}

	if DB.Find(&User{Name: u.Name}).Error != nil {
		return fmt.Errorf("repeated user name")
	}

	return DB.Create(u).Error
}

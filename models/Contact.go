package models

import (
	"log"

	"github.com/jyu/utils"
	"gorm.io/gorm"
)

type Contacts struct {
	gorm.Model
	OwnerId     string
	TargetId    string
	Description string
	ChatType    int64
}

func (table *Contacts) TableNanme() string {
	return "contact"
}

func SearchFriend(account_friend string) ([]UserBasic, error) {
	contacts := make([]Contacts, 0)
	objIds := make([]string, 0)
	err := utils.DB_MySQL.Where("owner_id = ?", account_friend).Find(&contacts).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(contacts)
	for _, v := range contacts {
		log.Println(v)
		objIds = append(objIds, string(v.TargetId))
	}
	log.Println(objIds)
	users := make([]UserBasic, 0)
	err = utils.DB_MySQL.Where("account in ?", objIds).Find(&users).Error
	return users, err
}

func AddFriend(owner_account string, friend_account string) (bool, error) {
	if friend_account == "" {
		return	false, nil
	} else {
		friend_user, err := FindUserByAccount(friend_account)
		if friend_user.Account == "" {
			return false, nil
		}
		if err != nil {
			log.Println(err)
			return false, err
		}
		tx := utils.DB_MySQL.Begin()
		utils.DB_MySQL.Create(&Contacts{
			OwnerId:     owner_account,
            TargetId:    friend_account,
            Description:  "好友",
            ChatType:    1,
		})
		utils.DB_MySQL.Create(&Contacts{
			OwnerId:     friend_account,
            TargetId:    owner_account,
            Description:  "好友",
            ChatType:    1,
		})
		tx.Commit()
		return true, nil
	}
}	

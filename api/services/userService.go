package services

import (
	"errors"
	"log"
	"time"

	"github.com/QuocBao92/go-sample/api/models"
	"github.com/jinzhu/gorm"
)

func SaveUser(db *gorm.DB, u *models.User) (*models.User, error) {
	// var err error
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &models.User{}, err
	}

	return u, nil
}

func FindAllUsers(db *gorm.DB, u *models.User) (*[]models.User, error) {
	users := []models.User{}
	err := db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]models.User{}, err
	}

	return &users, err
}

func FindUserByID(db *gorm.DB, u *models.User, uid uint32) (*models.User, error) {
	var err error
	err = db.Debug().Model(models.User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &models.User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &models.User{}, errors.New("User Not Found")
	}
	return u, err
}

func UpdateAUser(db *gorm.DB, u *models.User, uid uint32) (*models.User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.Nickname,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &models.User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &models.User{}, err
	}
	return u, nil
}

func DeleteAUser(db *gorm.DB, u *models.User, uid uint32) (int64, error) {

	db = db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

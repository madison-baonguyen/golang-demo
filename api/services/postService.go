package services

import (
	"errors"
	"time"

	"github.com/QuocBao92/go-sample/api/models"
	"github.com/jinzhu/gorm"
)

func SavePost(db *gorm.DB, p *models.Post) (*models.Post, error) {
	var err error
	err = db.Debug().Model(&models.Post{}).Create(&p).Error
	if err != nil {
		return &models.Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&models.User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &models.Post{}, err
		}
	}
	return p, nil
}

func FindAllPosts(db *gorm.DB, p *models.Post) (*[]models.Post, error) {
	var err error
	posts := []models.Post{}
	err = db.Debug().Model(&models.Post{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]models.Post{}, err
	}
	if len(posts) > 0 {
		for i, _ := range posts {
			err := db.Debug().Model(&models.User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]models.Post{}, err
			}
		}
	}
	return &posts, nil
}

func FindPostByID(db *gorm.DB, p *models.Post, pid uint64) (*models.Post, error) {
	var err error
	err = db.Debug().Model(&models.Post{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &models.Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&models.User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &models.Post{}, err
		}
	}
	return p, nil
}

func UpdateAPost(db *gorm.DB, p *models.Post) (*models.Post, error) {

	var err error

	err = db.Debug().Model(&models.Post{}).Where("id = ?", p.ID).Updates(models.Post{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &models.Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&models.User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &models.Post{}, err
		}
	}
	return p, nil
}

func DeleteAPost(db *gorm.DB, p *models.Post, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&models.Post{}).Where("id = ? and author_id = ?", pid, uid).Take(&models.Post{}).Delete(&models.Post{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

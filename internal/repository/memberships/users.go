package memberships

import (
	"github.com/Fairuzzzzz/music-catalog/internal/models/memberships"
	"gorm.io/gorm"
)

func (r *repository) CreateUser(model memberships.User) error {
	return r.db.Create(&model).Error
}

func (r *repository) GetUser(email, username string, id uint) (*memberships.User, error) {
	user := memberships.User{}
	res := r.db.Where("email = ?", email).Or("username = ?", username).Or("id = ?", id).Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

package repository

import (
	"basic-gin/entity"

	"gorm.io/gorm"
)

type UserRepository struct{
	db         *gorm.DB
}
func NewUserRepository(db *gorm.DB) UserRepository{
	return UserRepository{db}
}

// Membuat User
func (r *UserRepository) CreateUser(user *entity.User) (error) {
	// Menyimpan user ke database
	err := r.db.Create(user).Error
	return err
}

func (r *UserRepository) FindByUsername(username string) (entity.User, error){
	user := entity.User{}
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepository) GetUserById(id uint) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

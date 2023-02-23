package repository

import (
	"basic-gin/entity"
	"basic-gin/model"
	"basic-gin/sdk/crypto"

	"gorm.io/gorm"
)

type UserRepository struct{
	db         *gorm.DB
}

// Membuat User
func (r *UserRepository) CreateUser( model model.RegisterUser) (*entity.User, error) {
	// Ingat, sebelum menyimpan data user ke database, sebaiknya lakukan hashing password terlebih dahulu
	hashPassword, err := crypto.HashValue(model.Password)
	// Pengecekan error
	if err != nil {
		return nil, err
	}
	// Membuat user
	var user entity.User = entity.User{
		Name:     model.Name,
		Username: model.Username,
		Password: hashPassword,
	}
	// Menyimpan user ke database
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (entity.User, error){
	user := entity.User{}
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepository) GetUserById( id string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("id = ?", id).Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

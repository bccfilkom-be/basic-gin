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
func NewUserRepository(db *gorm.DB) UserRepository{
	return UserRepository{db}
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

func (r *UserRepository) GetUserById(id uint) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

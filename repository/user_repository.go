package repository

import (
	"basic-gin/entity"
	"basic-gin/model"
	"basic-gin/sdk/crypto"
	sdk_jwt "basic-gin/sdk/jwt"

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

func (r *UserRepository) LoginUser( model model.LoginUser) (map[string]any, error) {
	var user entity.User
	/*
		Kita cari terlebih dahulu buat tahu apakah beneran user ada di database atau tidak
		berdasarkan username user
	*/
	result := r.db.Where("username = ?", model.Username).Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	// Misalkan beneran ada, berarti kita coba validasi, apakah password mereka sama(Ingat, password di database adalah password yang sudah dihash)
	err := crypto.ValidateHash(model.Password, user.Password)
	if err != nil {
		return nil, err
	}
	// Karena sama, maka kita bisa generate token Jwt yang membuktikan bahwa user ini beneran dia sendiri
	tokenJwt, err := sdk_jwt.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	res := map[string]any{
		"status":  "success",
		"message": "user berhasil login",
		"data":    user,
		"jwt":     tokenJwt,
	}
	return res, nil
}

func (r *UserRepository) GetUserById( id string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("id = ?", id).Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

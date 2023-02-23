package crypto

import (
	"basic-gin/entity"
	"basic-gin/model"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(payload entity.User) (string, error) {
	/*
		Jwt Token adalah salah satu jenis token yang digunakan untuk melihat apakah user
		terautorisasi untuk mendapatkan akses ke endpoint yang sifatnya sensitif atau privasi.

		Contoh :
		Kita asumsikan terdapat 3 endpoint
		1. /getuser -> get user adalah endpoint yang sifatnya public atau bisa diakses semua user
		2. /createuser -> create user adalah endpoint yang sifatnya privasi atau sensitif
		3. /updateuser -> update user adalah endpoint yang sifatnya privasi atau sensitif

		Untuk mengakses get user, tentu saja kalian tinggal hit /getuser di aplikasi seperti Postman
		atau Insomnia kalian, tetapi bagaimana dengan endpoint /createuser atau /updateuser?
		Karena 2 endpoint tersebut adalah endpoint yang sifatnya privasi atau sensitif, apakah
		masuk akal jika ketika kalian hit endpoint /createuser atau /updateuser, maka akan
		langsung memberikan output kembalian sukses? Tentu saja jika kita langsung hit 2 endpoint
		tersebut dan langsung mendapatkan kembalian, maka tidak bisa dikatakan endpoint
		yang sensitif atau privasi.

		Endpoint yang sensitif atau privasi butuh sesuatu yang membuktikan bahwa benar yang mengakses
		endpoint tersebut adalah kalian dan BUKAN ORANG LAIN. Disini lah datang Jwt atau
		JSON Web Token. Inti dari JSON Web Token adalah membuktikan kalau kalian adalah orang yang benar-benar
		mengakses endpoint tersebut.

		JWT dibagi jadi 3 bagian,
		header -> kurang lebih hanya berisi informasi tentang JWT kalian
		payload -> data diri kalian disimpan di dalam sini(JANGAN MASUKAN DATA YANG SENSITIF SEPERTI PASSWORD!)
		signature -> ini adalah hasil algoritma crypthographic yang ada, biasa HS256
	*/

	// get jwt expire from env
	expStr := os.Getenv("JWT_EXP")
	var exp time.Duration
	exp, err := time.ParseDuration(expStr)
	if expStr == "" || err != nil {
		exp = time.Hour * 1
	}


	tokenJwtSementara := jwt.NewWithClaims(jwt.SigningMethodHS256, model.NewUserClaims(payload.ID, exp))
	// secret_key sama seperti namanya adalah kunci rahasia yang digunakan untuk token jwt kalian.
	// secret_key HANYA BOLEH DIKETAHUI SAMA KALIAN SENDIRI dan PASTIKAN TIDAK DIKETAHUI ORANG LAIN!
	tokenJwt, err := tokenJwtSementara.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return "", err
	}
	return tokenJwt, nil
}

func DecodeToken(signedToken string, ptrClaims jwt.Claims, KEY string) (error) {

	token, err := jwt.ParseWithClaims(signedToken, ptrClaims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // method used to sign the token 
		if !ok {
			// wrong signing method
			return "", errors.New("wrong signing method")
		}
		return []byte(KEY), nil
	})

	if err != nil {
		// parse failed
		return fmt.Errorf("token has been tampered with")
	}

	if !token.Valid{
		// token is not valid
		return fmt.Errorf("invalid token")
	}

	return nil
}

package crypto

import (
	"basic-gin/entity"
	"os"

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

	tokenJwtSementara := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// Claims = Payload
		"id":        payload.ID,
		"username":  payload.Username,
		"createdAt": payload.CreatedAt.Unix(),
	})

	// secret_key sama seperti namanya adalah kunci rahasia yang digunakan untuk token jwt kalian.
	// secret_key HANYA BOLEH DIKETAHUI SAMA KALIAN SENDIRI dan PASTIKAN TIDAK DIKETAHUI ORANG LAIN!
	tokenJwt, err := tokenJwtSementara.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return "", err
	}
	return tokenJwt, nil
}

package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func HashValue(rawValue string) (string, error) {
	/*
		Hash adalah kemampuan untuk mengubah bentuk sebuah value menjadi kode-kode unik yang sulit
		untuk dipecahkan. Dengan adanya hash, maka akan sangat membantu menjaga keamanan data
		seseorang yang sifatnya PRIVASI. Contoh nya adalah password, adalah ide yang buruk
		apabila kita menyimpan password di database secara mentah(atau dengan kata lain
		password nya kelihatan di database). Akan sangat baik jika kita bisa membuat password
		tersebut dibentuk menjadi sebuah nilai yang bisa diakses oleh sembarang orang

		Contoh simpel :
		Password : ABCDEFGHIJ
		HashedPassword : $2iasd1202345VWkals..............................(Ini hanya contoh saja, bentuknya lebih kompleks)
	*/

	// Membuat hash password
	password, err := bcrypt.GenerateFromPassword([]byte(rawValue), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashPassword := string(password)
	return hashPassword, nil
}

/*
Fungsi ini digunakan untuk mengecek apakah value yang dimiliki dengan value hash nilainya
sama atau tidak
*/
func ValidateHash(rawValue, hashValue string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashValue), []byte(rawValue))
	return err
}

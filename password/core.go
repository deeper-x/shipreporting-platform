package password

import "golang.org/x/crypto/bcrypt"

import "log"

// HashPassword create hash from string
func HashPassword(passwd string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(passwd), 14)

	if err != nil {
		log.Println(err)
	}

	return string(b), err
}

// CheckHash given input password, check matching
func CheckHash(passwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))

	return err == nil
}

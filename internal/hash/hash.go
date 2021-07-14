package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(userPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
}

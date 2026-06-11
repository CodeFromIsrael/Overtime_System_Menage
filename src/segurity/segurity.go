package segurity

import "golang.org/x/crypto/bcrypt"

func Hash(Password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
}

func CheckPass(passHash, passString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(passString))
}

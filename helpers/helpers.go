package helpers

import (
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// generating uuid
func GenUuid() string {
	return uuid.New().String()
}

// hashing the password using the bcrypt package
func HashPass(pass string) (string, error) {

	password := []byte(pass)

	//setting the cost as default to for balancing between security and time taken to decrypt the hash
	hashedpass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	return string(hashedpass), err

}

// comparing wheather the encryped password and password is equal
func MatchPass(encryptedPass, pass string) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(pass)); err != nil {
		return false
	}

	return true
}

// helps to find the offset and limit for pagination
func FindLimitandOffset(pageNo, limit int) (int, int) {
	if limit == 0 {
		limit = 5
	}
	if pageNo == 0 {
		pageNo = 1
	}

	offset := (pageNo - 1) * limit

	return offset, limit
}

// converts the string to integer
func StrtoInt(s string) (int, error) {
	return strconv.Atoi(s)
}

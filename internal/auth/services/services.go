package services

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"micro_auth/internal/auth/models"
	"micro_auth/internal/database"
	"strings"
	"time"
)

func CreateUser(username, email, password string) (models.User, error) {
	db := database.DB
	hashedPassword, _ := setPassword(password)
	guid, _ := uuid.NewRandom()
	timestamp := time.Now().UTC()
	var user models.User
	err := db.QueryRowx(
		"INSERT INTO users (id, username, email, password, created_at) VALUES($1, $2, $3, $4, $5) RETURNING *",
		guid, username, NormalizeEmail(email), hashedPassword, timestamp,
	).StructScan(&user)
	return user, err
}

func NormalizeEmail(email string) string {
	trimmed := strings.TrimSpace(email)
	split := strings.Split(trimmed, "@")
	emailDomain := split[len(split)-1]
	emailDomain = strings.ToLower(emailDomain)
	split = split[:len(split)-1]
	emailName := strings.Join(split, "@")
	return emailName + "@" + emailDomain
}

func MakeRandomPassword(length int, allowedCharacters string) string {
	if allowedCharacters == "" {
		allowedCharacters =
			"abcdefghjklmnpqrstuvwxyz" +
				"ABCDEFGHJKLMNPQRSTUVWXYZ" +
				"23456789"
	}

	return getRandomString(length, allowedCharacters)
}

// Return a securely generated random string.
// The default length of 12 with a-z, A-Z, 0-9 character set returns
// a 71-bit value. log_2((26+26+10)^12) =~ 71 bits.
func getRandomString(length int, allowedCharacters string) string {
	if length == 0 {
		length = 12
	}
	if allowedCharacters == "" {
		allowedCharacters =
			"abcdefghijklmnopqrstuvwxyz" +
				"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
				"0123456789"
	}

	b := make([]rune, length)
	for i := range b {
		b[i] = []rune(allowedCharacters)[rand.Intn(len(allowedCharacters))]
	}

	return string(b)
}

func setPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func GetByUsername(username string) models.User {
	db := database.DB
	var user models.User
	_ = db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	return user
}

func GetByEmail(email string) models.User {
	db := database.DB
	var user models.User
	_ = db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	return user
}

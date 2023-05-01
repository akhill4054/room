package models

import (
	"errors"

	"github.com/akhill4054/room-backend/pkg/util"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement:true"`
	Username string `gorm:"unique;index;not null;default:null"`
	Email    string `gorm:"unique;not null;default:null"`
}

func (User) TableName() string {
	return "roomusers"
}

type Password struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement:true;"`
	UID      int    `gorm:"not null;unique;index;default:null;"`
	Password string `gorm:"not null;default:null;"`
	User     User   `gorm:"foreignKey:UID;constraint:OnDelete:SET NULL;not null;"`
}

func GetUser(id int) (*User, error) {
	user := User{}

	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		panic(err)
	}

	return &user, nil
}

func GetUserWithClaims(id int, username string) (*User, error) {
	user := User{}

	err := db.Where("id = ? AND username = ?", id, username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		panic(err)
	}

	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	user := User{}

	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		panic(err)
	}

	return &user, nil
}

func CreateUser(username string, email string, password string) (*User, error) {
	user := User{
		Username: username,
		Email:    email,
	}

	if len(password) < 8 {
		return nil, ErrUserPasswordValidation
	}

	var userCount int64

	db.Where("username = ?", username).Model(&User{}).Count(&userCount)
	if userCount != 0 {
		return nil, ErrUsernameAlreadyExists
	}

	db.Where("email = ?", email).Model(&User{}).Count(&userCount)
	if userCount != 0 {
		return nil, ErrEmailIsAlreadyInUse
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&user)
		if result.Error != nil {
			return result.Error
		}

		hashedPassword, err := util.GetPasswordHash(password)
		if err != nil {
			panic(err)
		}

		passwordRecord := Password{
			UID:      user.ID,
			Password: string(hashedPassword),
		}
		result = tx.Create(&passwordRecord)

		return result.Error
	}); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAuthenticationToken(uid int, username string, password string) (string, error) {
	passwordRecord := Password{}

	err := db.Where("uid = ?", uid).First(&passwordRecord).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", ErrInvalidUsernameOrPassword
		}
		panic(err)
	}

	isMatch := util.ComparePasswords(password, passwordRecord.Password)
	if !isMatch {
		return "", ErrInvalidUsernameOrPassword
	}

	token, err := util.GenerateToken(uid, username)
	if err != nil {
		panic(err)
	}

	return token, err
}

var (
	ErrUsernameAlreadyExists = errors.New("Username already exists")
	ErrEmailIsAlreadyInUse   = errors.New("Provided email address is already in use")
	ErrUserNotFound          = errors.New("User not found")

	ErrUserPasswordValidation    = errors.New("Password must be at least 8 characters long")
	ErrInvalidUsernameOrPassword = errors.New("Wrong username or password")
)

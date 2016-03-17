package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	// UserUsernameMinLength is the minimum length of the username.
	UserUsernameMinLength = "3"

	// UserUsernameMaxLength is the maximum length of the username.
	UserUsernameMaxLength = "255"

	// UserPasswordMinLength is the minimum length of the password.
	UserPasswordMinLength = "3"

	// UserPasswordMaxLength is the maximum length of the password.
	UserPasswordMaxLength = "255"
)

// UserDefaultOrder is the default ordering for user listings.
func UserDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Order(
		"users.username ASC",
	)
}

// Users is simply a collection of user structs.
type Users []*User

// User represents a user model definition.
type User struct {
	ID         uint        `json:"id" gorm:"primary_key"`
	Permission *Permission `json:"permission"`
	Slug       string      `json:"slug" sql:"unique_index"`
	Username   string      `json:"username" sql:"unique_index"`
	Email      string      `json:"email" sql:"unique_index"`
	Password   string      `json:"password,omitempty" sql:"-"`
	Hashword   string      `json:"-"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Mods       Mods        `json:"mods" gorm:"many2many:user_mods;"`
}

// BeforeSave invokes required actions before persisting.
func (u *User) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				u.Slug = slugify.Slugify(u.Username)
			} else {
				u.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", u.Username, i),
				)
			}

			notFound := db.Where(
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&User{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	if u.Email != "" {
		email, err := govalidator.NormalizeEmail(
			u.Email,
		)

		if err != nil {
			return fmt.Errorf("Failed to normalize email")
		}

		u.Email = email
	}

	if u.Password != "" {
		encrypt, err := bcrypt.GenerateFromPassword(
			[]byte(u.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return fmt.Errorf("Failed to encrypt password")
		}

		u.Hashword = string(encrypt)
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *User) Validate(db *gorm.DB) {
	if !govalidator.StringLength(u.Username, UserUsernameMinLength, UserUsernameMaxLength) {
		db.AddError(fmt.Errorf(
			"Username should be longer than %s and shorter than %s",
			UserUsernameMinLength,
			UserUsernameMaxLength,
		))
	}

	if u.Username != "" {
		notFound := db.Where("username = ?", u.Username).Not("id", u.ID).First(&User{}).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Username is already present"))
		}
	}

	if !govalidator.IsEmail(u.Email) {
		db.AddError(fmt.Errorf(
			"Email must be a valid email address",
		))
	}

	if u.Email != "" {
		normalized, _ := govalidator.NormalizeEmail(
			u.Email,
		)

		notFound := db.Where("email = ?", normalized).Not("id", u.ID).First(&User{}).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Email is already present"))
		}
	}

	if db.NewRecord(u) {
		if !govalidator.StringLength(u.Password, UserPasswordMinLength, UserPasswordMaxLength) {
			db.AddError(fmt.Errorf(
				"Password should be longer than %s and shorter than %s",
				UserPasswordMinLength,
				UserPasswordMaxLength,
			))
		}
	}
}

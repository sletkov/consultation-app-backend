package models

import "golang.org/x/crypto/bcrypt"

const (
	UserStudent = "student"
	UserTeacher = "teacher"
	UserAdmin   = "admin"
)

type UserRole string

type User struct {
	ID                string   `json:"id"`
	FullName          string   `json:"full_name"`
	Role              UserRole `json:"role"`
	Email             string   `json:"email"`
	Password          string   `json:"password,omitempty"`
	EncryptedPassword string   `json:"-"`
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) EncryptPassword() error {
	if len(u.Password) > 0 {
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 4)
		if err != nil {
			return err
		}

		u.EncryptedPassword = string(encryptedPassword)
	}

	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func (u *User) RemoveSensitiveFields() {
	u.Password = ""
	u.EncryptedPassword = ""
}

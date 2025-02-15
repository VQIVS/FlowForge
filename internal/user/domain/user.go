package domain

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"

	"github.com/VQIVS/FlowForge/pkg/conv"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

type (
	UserID uint
	Phone  string
)

func (p Phone) IsValid() bool {
	// TODO: Regexp validation
	return true
}

type User struct {
	ID        UserID
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Password  string
	Phone     Phone
}

func (u *User) Validate() error {
	if !u.Phone.IsValid() {
		return errors.New("invalid phone number")
	}
	return nil
}

func (u *User) PasswordIsCorrect(password string) bool {
	return NewPassword(password) == u.Password
}

func NewPassword(password string) string {
	h := sha256.New()
	h.Write(conv.ToBytes(password))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

type UserFilter struct {
	ID    UserID
	Phone string
}

func (f *UserFilter) IsValid() bool {
	f.Phone = strings.TrimSpace(f.Phone)
	return f.ID > 0 || len(f.Phone) > 0
}

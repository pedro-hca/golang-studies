package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type User struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
}

func Init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (u *User) Validator() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	return nil
}

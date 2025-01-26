package api

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"`
	Firstname string    `json:"firstname" binding:"required"`
	Lastname  string    `json:"lastname" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Age       uint      `json:"age" binding:"required"`
	Created   time.Time `json:"created"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Created = time.Now()
	u.ID = uuid.New()
	return
}

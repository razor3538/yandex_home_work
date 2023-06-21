package domain

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;" json:"id"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}

// BeforeCreate will set a UUID rather than numeric ID
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	if base.ID.IsNil() {
		uuidv4, _ := uuid.NewV4()
		return scope.SetColumn("ID", uuidv4)
	}
	return nil
}

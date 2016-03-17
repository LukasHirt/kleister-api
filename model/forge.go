package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// ForgeDefaultOrder is the default ordering for forge listings.
func ForgeDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Order(
		"forges.minecraft DESC",
	).Order(
		"forges.name DESC",
	)
}

// Forges is simply a collection of forge structs.
type Forges []*Forge

// Forge represents a forge model definition.
type Forge struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"version" sql:"unique_index"`
	Minecraft string    `json:"minecraft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Builds    Builds    `json:"builds"`
}

// BeforeSave invokes required actions before persisting.
func (u *Forge) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		for {
			u.Slug = uuid.NewV4().String()

			notFound := db.Where(
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&Forge{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

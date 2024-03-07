package minecraft

import (
	"context"
	"errors"

	"github.com/kleister/go-minecraft/version"
	"github.com/kleister/kleister-api/pkg/model"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle *gorm.DB
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(handle *gorm.DB) *GormService {
	return &GormService{
		handle: handle,
	}
}

// Search implements the Store interface for database persistence.
func (s *GormService) Search(ctx context.Context, search string) ([]*model.Minecraft, error) {
	records := make([]*model.Minecraft, 0)
	q := s.query(ctx)

	if search != "" {
		q = q.Or(
			"name LIKE ?",
			"%"+search+"%",
		).Or(
			"type LIKE ?",
			"%"+search+"%",
		)
	}

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Show implements the Service interface for database persistence.
func (s *GormService) Show(ctx context.Context, name string) (*model.Minecraft, error) {
	record := &model.Minecraft{}

	err := s.query(ctx).Where(
		"id = ?",
		name,
	).Or(
		"name = ?",
		name,
	).First(
		record,
	).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return record, ErrNotFound
	}

	return record, err
}

// Sync implements the Store interface for database persistence.
func (s *GormService) Sync(ctx context.Context, versions version.Versions) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	for _, v := range versions {
		if err := tx.Where(
			&model.Minecraft{
				Name: v.ID,
			},
		).Attrs(
			&model.Minecraft{
				Type: "release",
			},
		).FirstOrCreate(
			&model.Minecraft{},
		).Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Order(
		"name ASC",
	).Model(
		&model.Minecraft{},
	)
}

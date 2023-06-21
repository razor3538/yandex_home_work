package migrate

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
	"server/config"
	"server/internal/domain"
)

// Migrate run migration for all entities and add constraints for them
func Migrate() {
	db := config.DB
	id, _ := uuid.NewV4()

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: id.String(),
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&domain.User{}).Error
				if err != nil {
					return err
				}

				err = tx.AutoMigrate(&domain.Password{}).Error
				if err != nil {
					return err
				}

				err = tx.AutoMigrate(&domain.Text{}).Error
				if err != nil {
					return err
				}

				err = tx.AutoMigrate(&domain.Cards{}).Error
				if err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				err := tx.DropTable("users").Error
				if err != nil {
					return err
				}

				err = tx.DropTable("passwords").Error
				if err != nil {
					return err
				}

				err = tx.DropTable("texts").Error
				if err != nil {
					return err
				}

				err = tx.DropTable("cards").Error
				if err != nil {
					return err
				}
				return nil
			},
		},
	})

	err := m.Migrate()
	if err != nil {
		panic(err)
	}

	if err == nil {
		println("Migration did run successfully")
	} else {
		println("Could not migrate: %v", err)
	}
}

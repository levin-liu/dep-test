package util

import (
	"context"
	"sync"

	"gorm.io/gorm"
)

var migrateOnce sync.Once

func Migrate(ctx context.Context, db *gorm.DB, tables ...interface{}) (err error) {
	migrateOnce.Do(func() {
		for _, table := range tables {

			if err = db.WithContext(ctx).AutoMigrate(table); err != nil {

				return
			}

			if migratable, ok := table.(Migratable); ok {

				if err = migratable.AfterMigrate(db); err != nil {
					return
				}
			}
		}
	})
	return err
}

type Migratable interface {
	AfterMigrate(db *gorm.DB) error
}

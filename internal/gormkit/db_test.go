package gormkit

import (
	"context"
	"testing"

	"github.com/levin-liu/dep-test/internal/util"
)

func TestDBMigrate(t *testing.T) {
	dsn := "host=localhost user=db password=123 dbname=mcd_ketchup port=5432 sslmode=disable TimeZone=Asia/Tokyo"

	db := GetLocalKetchup(dsn)
	ctx := context.Background()
	if err := util.Migrate(ctx, db, User{}); err != nil {
		t.Fatal(err)
	}
}

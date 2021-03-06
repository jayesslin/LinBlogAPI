package db

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	DBConsul = NewGormConsul()
)

func Init() error {
	if DBConsul.Init() {
		log.Info("DB Connection success !!!")
	} else {
		log.Warn("DB Connection failed")
		return errors.New("DB Connection Error !!! <------ ")
	}
	return nil
}

func ReadDB(ctx context.Context) *gorm.DB {
	return DBConsul.ReadDB(ctx)
}

func WriteDB(ctx context.Context) *gorm.DB {
	return DBConsul.WriteDB(ctx)
}

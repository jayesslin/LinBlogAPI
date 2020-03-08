package db

import (
	"context"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GormConsul 通过Consul发现MySQL读写实例, 绑定到gorm对象
type GormConsul struct {
	ctx       context.Context
	writeGorm *gorm.DB
	readGorm  *gorm.DB
}

func NewGormConsul() *GormConsul {
	return &GormConsul{}
}

func (s *GormConsul) Init() bool {
	p := s.initGormPool()
	return p
}

func (s *GormConsul) initGormPool() bool {
	var writeErr, readErr error
	s.writeGorm, writeErr = s.newGormInstance(DatabaseMaster)
	s.readGorm, readErr = s.newGormInstance(DatabaseMaster)
	return writeErr == nil && readErr == nil
}

func (s *GormConsul) newGormInstance(serviceName string) (*gorm.DB, error) {
	var connStr string
	connStr = DatabaseMaster

	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Warn(err.Error())
		return nil, err
	}

	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(20)
	return db, nil
}

// WriteDB 获取可读写Gorm实例
func (s *GormConsul) WriteDB(ctx context.Context) *gorm.DB {
	return s.writeGorm
}

// ReadDB 获取只读Gorm实例
func (s *GormConsul) ReadDB(ctx context.Context) *gorm.DB {
	return s.readGorm
}

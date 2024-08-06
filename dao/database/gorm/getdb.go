package gorm

import (
	"context"
	loggeri "github.com/hopeio/utils/dao/database/gorm/logger"
	"github.com/hopeio/utils/log"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBWithLogger(db *gorm.DB, log *log.Logger, conf *logger.Config) *gorm.DB {
	return db.Session(&gorm.Session{
		Logger: &loggeri.Logger{Logger: log.Logger,
			Config: conf,
		}})
}

func NewDBWithContext(db *gorm.DB, ctx context.Context) *gorm.DB {
	return db.Session(&gorm.Session{Context: ctx})
}

func NewTraceDB(db *gorm.DB, ctx context.Context, traceId string) *gorm.DB {
	return db.Session(&gorm.Session{Context: loggeri.SetTranceId(ctx, traceId), NewDB: true})
}

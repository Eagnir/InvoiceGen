package repository

import (
	"InvoiceGen/infrastructure/repository/exception"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var newLogger logger.Interface = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second,   // Slow SQL threshold
		LogLevel:                  logger.Silent, // Log level
		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,         // Disable color
	},
)

var SQLiteFilePath string = "test.db"
var DefaultConfig *gorm.Config = &gorm.Config{
	DisableForeignKeyConstraintWhenMigrating: false,
	Logger:                                   newLogger,
}

type DBContext struct {
	Context *gorm.DB
}

func NewDBContext() *DBContext {
	return &DBContext{}
}

//NewUserMySQL create new repository
func (ctx *DBContext) OpenContext() error {
	if ctx.Context != nil {
		return nil // Connection Already Open
		//err := ctx.CloseContext()
		//if err != nil {
		//	return err
		//}
	}
	db, err := gorm.Open(sqlite.Open(SQLiteFilePath), DefaultConfig)
	if err != nil {
		return err
	}
	ctx.Context = db
	return nil
}

func (ctx *DBContext) CloseContext() error {
	if ctx.Context == nil {
		return exception.GORM_ContextDoesNotExist
	}
	db, err := ctx.Context.DB()
	if err != nil {
		return err
	}
	ctx.Context = nil
	db.Close()
	return nil
}

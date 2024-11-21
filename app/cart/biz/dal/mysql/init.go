package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/YiD11/gomall/app/cart/biz/model"
	"gorm.io/plugin/opentelemetry/tracing"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(os.Getenv("MYSQL_DSN"), os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		log.Panicln(err)
	}
	// isFirstCreateTable := DB.Migrator().HasTable(&model.Cart{})
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		DB.AutoMigrate(&model.Cart{})
	}
}

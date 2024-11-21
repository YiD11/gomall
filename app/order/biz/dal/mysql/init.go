package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/YiD11/gomall/app/order/biz/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
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

	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "online" {
		if err = DB.AutoMigrate(&model.Order{}); err != nil {
			log.Panicln(err)
		}

		if err = DB.AutoMigrate(&model.OrderItem{}); err != nil {
			log.Panicln(err)
		}
	}
}

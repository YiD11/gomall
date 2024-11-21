package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/YiD11/gomall/app/product/biz/model"

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
	needDemoData := !DB.Migrator().HasTable(&model.Product{})

	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		log.Panicln(err)
	}

	if os.Getenv("GO_ENV") != "online" && needDemoData {
		err := DB.AutoMigrate(
			&model.Product{},
			&model.Category{},
		)
		if err != nil {
			log.Panicln(err)
		}
		if needDemoData {
			// DB.Model(&model.Category{}).Create(&model.Category{Name: "T-shirt", Description: "T-shirt"})
			// DB.Model(&model.Category{}).Create(&model.Category{Name: "Sticker", Description: "Sticker"})

			// DB.Model(&model.Product{}).Create(&model.Product{
			// 	Name: "T-shirt1",
			// 	Description: "This is T-shirt1",
			// 	Picture: "/static/image/t-shirt1.jpeg",
			// 	Price: 100.5,
			// 	Categories: []model.Category{},
			// })
			// DB.Model(&model.Product{}).Create(&model.Product{Name: "T-shirt2", Description: "This is T-shirt2", Picture: "/static/image/t-shirt2.jpg", Price: 200.1})
			// DB.Model(&model.Product{}).Create(&model.Product{Name: "T-shirt3", Description: "This is T-shirt3", Picture: "/static/image/t-shirt2.jpg", Price: 300.1})

			DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06', NULL,'T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06', NULL,'Sticker','Sticker')")
			
			DB.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', NULL, 'golang', 'A golang sticker', '/static/image/go-docker.png', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', NULL, 'Mouse-Pad', 'The mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mousepad.jpg', 8.80 ), ( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt.jpeg', 6.60 ), ( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.jpeg', 2.20 ), ( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', NULL, 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/sweatshirt.jpg', 1.10 ), ( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.jpeg', 1.80 ), ( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80 )")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")
		}
	}
}

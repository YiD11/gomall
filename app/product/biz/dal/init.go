package dal

import (
	"github.com/YiD11/gomall/app/product/biz/dal/mysql"
	"github.com/YiD11/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

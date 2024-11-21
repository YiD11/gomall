package dal

import (
	"github.com/YiD11/gomall/app/checkout/biz/dal/mysql"
	"github.com/YiD11/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

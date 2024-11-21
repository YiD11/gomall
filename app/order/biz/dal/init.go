package dal

import (
	"github.com/YiD11/gomall/app/order/biz/dal/mysql"
	"github.com/YiD11/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

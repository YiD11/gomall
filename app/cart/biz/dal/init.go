package dal

import (
	"github.com/YiD11/gomall/app/cart/biz/dal/mysql"
	"github.com/YiD11/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

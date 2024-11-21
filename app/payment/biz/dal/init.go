package dal

import (
	"github.com/YiD11/gomall/app/payment/biz/dal/mysql"
	"github.com/YiD11/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

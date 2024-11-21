package dal

import (
	"github.com/YiD11/gomall/app/email/biz/dal/mysql"
	"github.com/YiD11/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

package dal

import (
	"github.com/YiD11/gomall/app/user/biz/dal/mysql"
	"github.com/YiD11/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

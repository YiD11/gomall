package dal

import (
	"github.com/YiD11/gomall/app/frontend/biz/dal/mysql"
	"github.com/YiD11/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

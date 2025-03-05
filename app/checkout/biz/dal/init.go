package dal

import (
	"github.com/qitian118/gomall/app/checkout/biz/dal/mysql"
	"github.com/qitian118/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

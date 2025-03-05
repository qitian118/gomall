package dal

import (
	"github.com/qitian118/gomall/app/order/biz/dal/mysql"
	"github.com/qitian118/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

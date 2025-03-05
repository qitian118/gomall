package dal

import (
	"github.com/qitian118/gomall/app/cart/biz/dal/mysql"
	"github.com/qitian118/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

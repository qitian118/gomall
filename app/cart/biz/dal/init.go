package dal

import (
	"github.com/gomall_project/app/cart/biz/dal/mysql"
	"github.com/gomall_project/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

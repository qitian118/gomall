package dal

import (
	"github.com/qitian118/gomall/app/auth/biz/dal/mysql"
	"github.com/qitian118/gomall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

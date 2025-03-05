package dal

import (
	"github.com/qitian118/gomall/app/user/biz/dal/mysql"
	"github.com/qitian118/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

package dal

import (
	"github.com/qitian118/gomall/app/frontend/biz/dal/mysql"
	"github.com/qitian118/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

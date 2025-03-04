package dal

import (
	"github.com/gomall_project/app/frontend/biz/dal/mysql"
	"github.com/gomall_project/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

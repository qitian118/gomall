package dal

import (
	"github.com/gomall_project/app/user/biz/dal/mysql"
	"github.com/gomall_project/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

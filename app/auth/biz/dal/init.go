package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/auth/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

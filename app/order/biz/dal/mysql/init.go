package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	"github.com/cloudwego/biz-demo/gomall/app/order/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	_ = godotenv.Load()
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		if err := DB.AutoMigrate(&model.Order{}, &model.OrderItem{}); err != nil {
			klog.Error(err)
		}
	}
}

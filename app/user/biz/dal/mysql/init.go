package mysql

import (
	"fmt"
	"os"
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gomall_project/app/user/biz/model"
	"github.com/gomall_project/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	klog.Infof("MySQL DSN信息: %s", strings.Replace(dsn, os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_PASSWORD"), 1))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}

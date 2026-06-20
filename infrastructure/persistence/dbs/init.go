package dbs

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	config "github.com/Y1le/godolist/conf"
)

var DB *gorm.DB

func MySQLInit() {
	conf := config.Conf.MySQL
	conn := strings.Join([]string{conf.UserName, ":", conf.Password, "@tcp(", conf.Host, ":", conf.Port, ")/", conf.Database, "?charset=utf8&parseTime=true"}, "")
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,  // DSN data source name
		DefaultStringSize:         256,   // string 绫诲瀷瀛楁鐨勯粯璁ら暱搴?
		DisableDatetimePrecision:  true,  // 绂佺敤 datetime 绮惧害锛孧ySQL 5.6 涔嬪墠鐨勬暟鎹簱涓嶆敮鎸?
		DontSupportRenameIndex:    true,  // 閲嶅懡鍚嶇储寮曟椂閲囩敤鍒犻櫎骞舵柊寤虹殑鏂瑰紡锛孧ySQL 5.7 涔嬪墠鐨勬暟鎹簱鍜?MariaDB 涓嶆敮鎸侀噸鍛藉悕绱㈠紩
		DontSupportRenameColumn:   true,  // 鐢?`change` 閲嶅懡鍚嶅垪锛孧ySQL 8 涔嬪墠鐨勬暟鎹簱鍜?MariaDB 涓嶆敮鎸侀噸鍛藉悕鍒?
		SkipInitializeWithVersion: false, // 鏍规嵁鐗堟湰鑷姩閰嶇疆
	}), &gorm.Config{
		Logger: ormLogger, // 鎵撳嵃鏃ュ織
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 琛ㄦ槑涓嶅姞s
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 璁剧疆杩炴帴姹狅紝绌洪棽
	sqlDB.SetMaxOpenConns(100) // 鎵撳紑
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}

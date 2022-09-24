package core

import (
	"fmt"
	"github.com/go-ini/ini"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var (
	Cfg           *ini.File
	JwtSecret     string
	READ_TIMEOUT  time.Duration
	WRITE_TIMEOUT time.Duration
	Addr          string

	RunTimeMode string
	DB          *gorm.DB
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/conf.ini")
	if err != nil {

	}
	app, err := Cfg.GetSection("app")
	JwtSecret = app.Key("JWT_SECRET").String()
	fmt.Println(JwtSecret)

	loadDb(Cfg)
	loadRuntimeModel(Cfg)
	loadServer(Cfg)

}

func loadDb(Cfg *ini.File) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	section, _ := Cfg.GetSection("database")
	user := section.Key("USER").String()
	port, _ := section.Key("PORT").Int()
	host := section.Key("HOST").String()
	pw := section.Key("PW").String()
	db_name := section.Key("DB_NAME").String()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, user, pw, db_name, port)
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // 使用单数表名

	},
		Logger:      newLogger,
		QueryFields: true})
}

func loadRuntimeModel(Cfg *ini.File) {
	RunTimeMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer(Cfg *ini.File) {
	section := Cfg.Section("server")
	Addr = section.Key("Addr").String()
	WRITE_TIMEOUT = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	READ_TIMEOUT = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second

}

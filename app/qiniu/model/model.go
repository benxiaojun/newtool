package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/olebedev/config"
	"github.com/sillydong/goczd/gofile"
	"os"
	"path"
)

var (
	x         *xorm.Engine
	HasEngine bool
	DbCfg     struct {
		Host, DbName, UserName, PassWord, LogDir string
	}
)

func LoadConfig(conf string) error {
	if gofile.FileExists(conf) {
		cfg, err := config.ParseJsonFile(conf)
		if err != nil {
			return err
		} else {
			DbCfg.Host = cfg.UString("host", "localhost:3306")
			DbCfg.DbName = cfg.UString("dbname", "")
			DbCfg.UserName = cfg.UString("username", "")
			DbCfg.PassWord = cfg.UString("password", "")
			DbCfg.LogDir = cfg.UString("logdir", "")

			return nil
		}
	} else {
		return fmt.Errorf("file not exists: %s\n", conf)
	}
}

func InitEngine() (err error) {
	cnnstr := ""
	if DbCfg.Host[0] == '/' {
		// looks like a unix socket
		cnnstr = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=utf8",
			DbCfg.UserName, DbCfg.PassWord, DbCfg.Host, DbCfg.DbName)
	} else {
		cnnstr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			DbCfg.UserName, DbCfg.PassWord, DbCfg.Host, DbCfg.DbName)
	}

	x, err = xorm.NewEngine("mysql", cnnstr)
	if err != nil {
		return err
	}

	x.SetDefaultCacher(xorm.NewLRUCacher2(xorm.NewMemoryStore(), 3600, 1000))
	x.SetMapper(core.NewCacheMapper(core.GonicMapper{}))

	//x.SetLogger(goxorm.NewXormLogger(DbCfg.LogDir, "db.log", core.LOG_WARNING))
	logPath := path.Join(DbCfg.LogDir, "db.log")
	os.MkdirAll(path.Dir(logPath), os.ModePerm)

	f, err := os.Create(logPath)
	if err != nil {
		return fmt.Errorf("fail to create db.log: %v\n", err)
	}
	x.SetLogger(xorm.NewSimpleLogger(f))

	x.ShowSQL(true)

	if err = Ping(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
		return err
	}

	HasEngine = true
	return nil
}

func Ping() error {
	return x.Ping()
}

func DumpDatabase(filePath string) error {
	return x.DumpAllToFile(filePath)
}

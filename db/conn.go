package db

import (
	"blog/config"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var MasterDB *xorm.Engine
var dns string

func init() {
	fillDns()

	// 启动时就打开数据库连接
	if err := initEngine(); err != nil {
		panic(err)
	}
}

var (
	ConnectDBErr = errors.New("connect db error")
	UseDBErr     = errors.New("use db error")
)

// TestDB 测试数据库
func TestDB() error {
	fillDns()
	egnine, err := xorm.NewEngine("mysql", dns)
	if err != nil {
		fmt.Println("new engine error:", err)
		return err
	}
	defer egnine.Close()

	// 测试数据库连接是否 OK
	if err = egnine.Ping(); err != nil {
		fmt.Println("ping db error:", err)
		return ConnectDBErr
	}

	_, err = egnine.Exec("use " + "test1")
	if err != nil {
		fmt.Println("use db error:", err)
		_, err = egnine.Exec("CREATE DATABASE " + "test1" + " DEFAULT CHARACTER SET " + "utf8")
		if err != nil {
			fmt.Println("create database error:", err)

			return UseDBErr
		}

		fmt.Println("create database successfully!")
	}

	// 初始化 MasterDB
	Init()

	return nil
}

func Init() error {
	fillDns()

	// 启动时就打开数据库连接
	if err := initEngine(); err != nil {
		fmt.Println("mysql is not open:", err)
		return err
	}

	return nil
}

func fillDns() {
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.MysqlUser,
		config.MysqlPassword,
		config.MysqlHost,
		config.MysqlPort,
		config.MysqlDbname,
		config.MysqlCharset)
}

func initEngine() error {
	var err error

	MasterDB, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		return err
	}

	maxIdle := 2
	maxConn := 10

	MasterDB.SetMaxIdleConns(maxIdle)
	MasterDB.SetMaxOpenConns(maxConn)

	showSQL := false
	logLevel := 1

	MasterDB.ShowSQL(showSQL)
	MasterDB.Logger().SetLevel(core.LogLevel(logLevel))

	// 启用缓存
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// MasterDB.SetDefaultCacher(cacher)

	return nil
}

func StdMasterDB() *sql.DB {
	return MasterDB.DB().DB
}

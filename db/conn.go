package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"golangschool/common/tools"
	"golangschool/config"
)

var MasterDB *xorm.Engine
var dns string
var (
	ConnectDBErr = errors.New("connect db error")
	UseDBErr     = errors.New("use db error")
)

func init() {
	fillDns()
	// 启动时就打开数据库连接
	if err := initEngine(); err != nil {
		panic(err)
	}
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


	MasterDB.ShowSQL(false)

	// 启用缓存
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// MasterDB.SetDefaultCacher(cacher)

	return nil
}

func StdMasterDB() *sql.DB {
	return MasterDB.DB().DB
}

// TestDB 测试数据库
func TestDB() error {
	fillDns()
	eg, err := xorm.NewEngine("mysql", dns)
	if err != nil {
		fmt.Println("new engine error:", err)
		return err
	}
	defer tools.CheckErr(eg.Close())

	// 测试数据库连接是否 OK
	if err = eg.Ping(); err != nil {
		fmt.Println("ping db error:", err)
		return ConnectDBErr
	}

	_, err = eg.Exec("use " + "test1")
	if err != nil {
		fmt.Println("use db error:", err)
		_, err = eg.Exec("CREATE DATABASE " + "test1" + " DEFAULT CHARACTER SET " + "utf8")
		if err != nil {
			fmt.Println("create database error:", err)

			return UseDBErr
		}

		fmt.Println("create database successfully!")
	}

	// 初始化 MasterDB
	//controllers.CheckErr(Init())
	tools.CheckErr(Init())
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

package models

import (
	"github.com/Unknwon/com" //通用函数包
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" //只引用初始化函数，不调用包内其他函数;仅仅是一个驱动使用
	"os"
	"path"
	"time"
)

//定义常量，设置数据库的名称，驱动的名称
const (
	_DB_NAME        = "date/beeblog.db" //输入路径，嵌入式数据库
	_SQLITE3_DRIVER = "sqlite3"         //
)

type Category struct { //注意字段开头大写，会反射字段和标签
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Content          string `orm:"size(5000)"`
	Attachment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int64     `orm:"index"`
	Author           string
	ReplyTime        time.Time `orm:"index"`
	ReplyCount       int64
	RepleyLastUserId int64
}

func RegisterDB() { //注册驱动路径
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
		//orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
	}

	orm.RegisterModel(new(Category), new(Topic))                   //orm 正确使用步骤1、注册模型
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)              //
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10) //注册默认数据库 数据库名 数据库 数据库路径 最大连接数
	//必须有一个名为 “default”
	//orm.RegisterDataBase("'default'", _SQLITE3_DRIVER, _DB_NAME, 12)

}

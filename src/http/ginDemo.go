package main

import (
	"database/sql"
	"fmt"
	"log"
	rand2 "math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
	USER_NAME = "root"
	PASS_WORD = "gx921016"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "mySql"
	CHARSET   = "utf8"
)

// 初始化链接
func dataBaseInit() {

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

	// 打开连接失败
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}

}

func daoMysql() {
	// 打开数据库,格式是⽤户名:密码@协议/数据库名称？编码⽅式
	dataBaseInit()
}

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/ping", func(context *gin.Context) {
		h := gin.H{
			"daibi": "老马",
		}
		fmt.Println("ping")
		if v, e := context.Get("requestId"); e {
			h["requestId"] = v
		}

		context.JSON(200, h)
	})

	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

var (
	g errgroup.Group
)

func main() {

	daoMysql()
	//server01 := &http.Server{
	//	Addr:         ":8080",
	//	Handler:      router01(),
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//}
	//
	//server02 := &http.Server{
	//	Addr:         ":8081",
	//	Handler:      router02(),
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//}

	//server01.ListenAndServe()
	//server02.ListenAndServe()
	//fmt.Println("aaaaa")
	//daoMysql()
	//g.Go(func() error {
	//	return server01.ListenAndServe()
	//})
	//
	//g.Go(func() error {
	//	return server02.ListenAndServe()
	//})
	//
	//if err := g.Wait(); err != nil {
	//	log.Fatal(err)
	//}

}

func ginTest() {
	r := gin.Default()
	logger, e := zap.NewProduction()
	if e != nil {
		panic(e)
	}

	r.Use(func(context *gin.Context) {
		s := time.Now()
		fmt.Println("Now")
		context.Next()
		if v, e := context.Get("requestId"); e {
			logger.Info("incoming request",
				zap.String("path", context.Request.URL.Path),
				zap.Int("status", context.Writer.Status()),
				zap.Duration("time", time.Now().Sub(s)),
				zap.Reflect("requestId", v))

		} else {
			logger.Info("incoming request",
				zap.String("path", context.Request.URL.Path),
				zap.Int("status", context.Writer.Status()),
				zap.Duration("time", time.Now().Sub(s)))
		}

	}, func(context *gin.Context) {
		fmt.Println("Set")
		context.Set("requestId", rand2.Int())
		context.Next()
	})

	r.GET("/ping", func(context *gin.Context) {
		h := gin.H{
			"daibi": "老马",
		}
		fmt.Println("ping")
		if v, e := context.Get("requestId"); e {
			h["requestId"] = v
		}

		context.JSON(200, h)
	})

	r.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello")
	})
	r.Run(":9090")

}

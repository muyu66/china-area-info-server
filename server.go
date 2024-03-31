package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"gonum.org/v1/gonum/graph/simple"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func loadData() []Area {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bad-comment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Infoln("已开启数据库连接")

	sqlDB, _ := db.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal(err)
		} else {
			log.Infoln("已关闭数据库连接")
		}

	}(sqlDB)

	var areas []Area
	db.Find(&areas)

	return areas
}

func makeGraph(areas *[]Area) *simple.DirectedGraph {
	dg := simple.NewDirectedGraph()
	dg.AddNode(simple.Node(0)) // 根节点

	// TODO: 双遍历有无改进空间
	for _, area := range *areas {
		node := UnitNode{
			Node: simple.Node(area.Id),
			Name: area.Name,
			Deep: area.Deep,
		}
		dg.AddNode(node)
	}

	for _, area := range *areas {
		edge := simple.Edge{F: dg.Node(int64(area.Id)), T: dg.Node(int64(area.Pid))}
		dg.SetEdge(edge)
	}
	return dg
}

func main() {
	log.SetLevel(log.ErrorLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(log.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")
			return nil
		},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10, // 4 KB
		DisableStackAll:   true,
		DisablePrintStack: false,
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			log.Error(err.Error())
			fmt.Println(string(stack))
			return err
		},
		DisableErrorHandler: false,
	}))

	var areas = loadData()
	var dg = makeGraph(&areas)

	ctl := &controller{dg: dg}
	router(e, ctl)

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	if err := e.StartH2CServer(":3000", s); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

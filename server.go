package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"gonum.org/v1/gonum/graph/simple"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	e := echo.New()
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	var areas = loadData()
	var dg = makeGraph(&areas)

	ctl := &controller{dg: dg}
	router(e, ctl)
	e.Logger.Fatal(e.Start(":3000"))
}

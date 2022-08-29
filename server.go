package main

import (
	"myapp/handler"
	"myapp/mdw"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	// register model
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	err := orm.RegisterDataBase("default", "mysql", "root:@/_imdbtop2500?charset=utf8")
	if err != nil {
		glog.Fatalf("Failed to register database: %v", err)
	}
}
func main() {
	server := echo.New()
	server.Use(middleware.Logger())

	isLogedIn := middleware.JWT([]byte("mysecretkey"))
	isAdmin := mdw.IsAdminMdv

	server.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))
	server.GET("/", handler.Hello, isLogedIn)

	groupUser := server.Group("/api/movie")
	groupUser.GET("/get", handler.GetMovie)
	groupUser.GET("/get_all", handler.GetAllMovies)
	groupUser.GET("/get_random", handler.GetRandomMovies)
	groupUser.POST("/update", handler.UpdateMovie)

	groupAdmin := server.Group("/api/movie", isLogedIn, isAdmin)
	groupAdmin.GET("/admin", handler.GetAllMovies)

	server.Logger.Fatal(server.Start(":8888"))
}

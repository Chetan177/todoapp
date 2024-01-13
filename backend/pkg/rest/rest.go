package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"todo/pkg/db"
)

type Rest struct {
	e  *echo.Echo
	db *db.DB
}

func NewRestServer() *Rest {
	e := echo.New()
	database := db.NewDB()
	return &Rest{
		e:  e,
		db: database,
	}
}

func (r *Rest) StartServer() {
	r.loadRoutes(r.e)

	log.Info("Starting Server")
	go func() {
		r.e.Logger.Fatal(r.e.Start(":1323"))
	}()
}

package main

import (
	"github.com/Conty111/AvitoTech/logging"
	"github.com/Conty111/AvitoTech/storage"
	"github.com/Conty111/AvitoTech/web"
)

func main() {
	logger := logging.NewLogger()
	db := storage.New(DBaddr, DBpassword, DBport, DBnum, logger)
	app := web.NewApp(db, logger, AppPort)
	app.Start()
}

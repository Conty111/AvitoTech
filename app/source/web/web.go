package web

import (
	"github.com/Conty111/AvitoTech/storage"
	"github.com/Conty111/AvitoTech/web/gorrila_mux"
	"go.uber.org/zap"
	"net/http"
)

type WebAPI interface {
	Start()
	GetRequest(w http.ResponseWriter, r *http.Request)
	SetRequest(w http.ResponseWriter, r *http.Request)
	DelRequest(w http.ResponseWriter, r *http.Request)
}

func NewApp(db storage.Storage, logger *zap.Logger, port int) WebAPI {
	api := gorrila_mux.NewGorillaAPI(db, logger, port)
	return api
}

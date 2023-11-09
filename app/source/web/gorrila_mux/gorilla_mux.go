package gorrila_mux

import (
	"fmt"
	"github.com/Conty111/AvitoTech/storage"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type GorillaApp struct {
	router http.Handler
	Port   int
	logger *zap.Logger
	db     storage.Storage
}

// Глобальная переменная, через которую будем обращаться к логгеру и БД
var g GorillaApp

// Запускаем сервер
func (g *GorillaApp) Start() {
	http.Handle("/", g.router)
	g.logger.Info("Сервер запущен", zap.Int("port", g.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", g.Port), g.router)
	if err != nil {
		g.logger.Fatal("Server starting error", zap.Error(err))
		return
	}
}

// Инициализирует глобальную переменную g GorillaApp и возвращает ссылку на нее
func NewGorillaAPI(db storage.Storage, logger *zap.Logger, port int) *GorillaApp {
	r := mux.NewRouter()
	r.Use(loggingMiddleware(logger))

	// Определяем маршруты
	r.NotFoundHandler = handler403{}
	r.HandleFunc("/get_key", g.GetRequest)
	r.HandleFunc("/set_key", g.SetRequest)
	r.HandleFunc("/del_key", g.DelRequest)

	g.router = r
	g.db = db
	g.logger = logger
	g.Port = port
	return &g
}

// Middleware для логирования запросов и ответов
func loggingMiddleware(logger *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			// Завершаем логирование после того, как запрос выполнен
			duration := time.Since(start)
			logger.Info("HTTP запрос",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Duration("duration", duration),
			)
		})
	}
}

// Обработчик для всех остальных URI
type handler403 struct{}

func (h handler403) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(403)
}

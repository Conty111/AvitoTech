package gorrila_mux

import (
	"database/sql"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"io"
	"net/http"
)

// Структура тела запроса в json (если передаются неизвестные
// пары ключ-значение - записываются в pairs)
type RBody struct {
	Key  string `json:"key"`
	Pair map[string]string
}

func (g *GorillaApp) GetRequest(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		g.logger.Warn("Get request with nil params")
		return
	}
	val, err := g.db.GetValue(key)
	if errors.Is(err, sql.ErrNoRows) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		g.logger.Error("Getting error", zap.Error(err))
	}
	_, err = w.Write([]byte(val))
	if err != nil {
		g.logger.Error("Cannot write response", zap.Error(err))
	}
}

func (g *GorillaApp) DelRequest(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			g.logger.Error("Body close error", zap.Error(err))
			return
		}
	}(r.Body)
	body, err := io.ReadAll(r.Body)
	var b RBody
	err = json.Unmarshal(body, &b)
	if err != nil {
		g.logger.Error("Cannot unmarshal request body", zap.Error(err))
	}
	if b.Key == "" {
		g.logger.Warn("Body key params is nil")
	}

	err = g.db.Delete([]string{b.Key})
	if err != nil {
		g.logger.Error("Error while deleting key-value", zap.String("key", b.Key), zap.Error(err))
	}
	w.WriteHeader(http.StatusOK)
}

func (g *GorillaApp) SetRequest(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			g.logger.Error("Body close error", zap.Error(err))
			return
		}
	}(r.Body)
	body, err := io.ReadAll(r.Body)
	var b RBody
	err = json.Unmarshal(body, &b.Pair)
	if err != nil {
		g.logger.Error("Cannot unmarshal request body", zap.Error(err))
	}

	for key, val := range b.Pair {
		err = g.db.SetValue(key, val)
		if err != nil {
			g.logger.Error("Error while setting key-value", zap.String("key", b.Key), zap.Error(err))
		}
	}
	w.WriteHeader(http.StatusOK)
}

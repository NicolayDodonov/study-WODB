package handler

import (
	"net/http"
	"study-WODB/internal/logger"
	"study-WODB/internal/storage/postgres"
)

type RestHandler struct {
}

func NewRestHandler(p *postgres.Storage, l *logger.Logger) *RestHandler {
	return &RestHandler{}
}

// Add отвечает за добавление нового ресторана в систему.
func (h *RestHandler) Add(w http.ResponseWriter, r *http.Request) {

}

// Get отвечает за получение всех ресторанов или конкретного ресторана по его ID.
func (h *RestHandler) Get(w http.ResponseWriter, r *http.Request) {

}

// Del отвечает за удаление ресторана из системы по его ID.
func (h *RestHandler) Del(w http.ResponseWriter, r *http.Request) {

}

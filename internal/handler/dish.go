package handler

import (
	"net/http"
	"study-WODB/internal/logger"
	"study-WODB/internal/storage/postgres"
)

type DishHandler struct {
}

func NewDishHandler(p *postgres.Storage, l *logger.Logger) *DishHandler {
	return &DishHandler{}
}

func (h *DishHandler) Add(w http.ResponseWriter, r *http.Request) {

}

func (h *DishHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *DishHandler) Del(w http.ResponseWriter, r *http.Request) {

}

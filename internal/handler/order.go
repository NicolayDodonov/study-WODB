package handler

import (
	"net/http"
	"study-WODB/internal/logger"
	"study-WODB/internal/storage/mongo"
	"study-WODB/internal/storage/postgres"
)

type OrderHandler struct {
	postgres *postgres.Storage
	mongo    *mongo.Storage
}

func NewOrderHandler(p *postgres.Storage, m *mongo.Storage, l *logger.Logger) *OrderHandler {
	return &OrderHandler{
		postgres: p,
		mongo:    m,
	}
}

// Make отвечает за создание нового заказа в системе.
// Формирует полноценную транзакцию на весь заказ, в ходе которой
// проверяет есть ли такой товар в ресторане на данный момент.
func (h *OrderHandler) Make(w http.ResponseWriter, r *http.Request) {
	// Проверить JWT на актуальность

	// Считать заказ в структуру

	// Создать транзакцию

	// Циклично проверять каждый элемент меню на наличие в указаном ресторане

	// Составить сумму заказа

	// Добавить в таблицу заказа

	// Перенаправить на страницу оплаты
}

// Close отвечает за отмену или закрытие заказа в системе.
// Получает информацию о пользователе и заказе что надо отменить.
func (h *OrderHandler) Close(w http.ResponseWriter, r *http.Request) {
	// Проверить JWT на актуальность

	// Получить информацию о действии

	// Проверить если ли такой ID в системе

	// Удалить заказа пользователя
}

// Pay отвечает за оплату заказа. В учебных целях только ИМИТИРУЕТ систему оплаты
// через приём информации об отправленной сумме рублей.
func (h *OrderHandler) Pay(w http.ResponseWriter, r *http.Request) {
	// Проверить JWT на актуальность

	// Получить информацию о заплаченной сумме

	// Проверить сумму заказа
}

// Feedback отвечает за создание нового отзыва в системе
func (h *OrderHandler) Feedback(w http.ResponseWriter, r *http.Request) {
	// Проверить JWT на актуальность

	// Получить информацию об отзыве

	// Записать отзыв в систему
}

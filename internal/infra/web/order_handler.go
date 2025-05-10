package web

import (
	"encoding/json"
	"net/http"

	"github.com/JMKobayashi/Clean-Architeture-Challenge/internal/entity"
	"github.com/JMKobayashi/Clean-Architeture-Challenge/internal/usecase"
	"github.com/JMKobayashi/Clean-Architeture-Challenge/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher    events.EventDispatcherInterface
	OrderRepository    entity.OrderRepositoryInterface
	OrderCreatedEvent  events.EventInterface
	CreateOrderUseCase *usecase.CreateOrderUseCase
	ListOrdersUseCase  *usecase.ListOrdersUseCase
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
	CreateOrderUseCase *usecase.CreateOrderUseCase,
	ListOrdersUseCase *usecase.ListOrdersUseCase,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:    EventDispatcher,
		OrderRepository:    OrderRepository,
		OrderCreatedEvent:  OrderCreatedEvent,
		CreateOrderUseCase: CreateOrderUseCase,
		ListOrdersUseCase:  ListOrdersUseCase,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.CreateOrderUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListOrdersUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

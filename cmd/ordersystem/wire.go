//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/JMKobayashi/Clean-Architeture-Challenge/internal/entity"
	"github.com/JMKobayashi/Clean-Architeture-Challenge/internal/event"
	"github.com/JMKobayashi/Clean-Architeture-Challenge/internal/infra/database"
	"github.com/JMKobayashi/Clean-Architeture-Challenge/internal/infra/web"
	"github.com/JMKobayashi/Clean-Architeture-Challenge/internal/usecase"
	"github.com/JMKobayashi/Clean-Architeture-Challenge/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
		usecase.NewListOrdersUseCase,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

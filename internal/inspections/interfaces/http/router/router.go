package router

import (
	"proj/internal/common/interfaces/http/router"
	"proj/internal/inspections/application/command"
	"proj/internal/inspections/application/query"
	"proj/internal/inspections/domain/repository"
	"proj/internal/inspections/interfaces/http/handler"
)

type ItemsRouter struct {
	itemsRepo repository.InspectionItemsRepository
}

func NewItemsRouter(itemsRepo repository.InspectionItemsRepository) *ItemsRouter {
	return &ItemsRouter{
		itemsRepo: itemsRepo,
	}
}

func (r *ItemsRouter) RegisterRoutes(group *router.Group) {
	cmdHandler := command.NewInspectionItemsCommandHandler(r.itemsRepo)
	queryHandler := query.NewInspectionItemsQueryHandler(r.itemsRepo)

	itemsHandler := handler.NewItemsHandler(cmdHandler, queryHandler)

	group.HandleFunc("POST /inspection/items", itemsHandler.CreateItem)
}

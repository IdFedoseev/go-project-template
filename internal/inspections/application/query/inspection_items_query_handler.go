package query

import "proj/internal/inspections/domain/repository"

type InspectionItemsQueryHandler struct {
	itemsRepository repository.InspectionItemsRepository
}

func NewInspectionItemsQueryHandler(itemsRepository repository.InspectionItemsRepository) *InspectionItemsQueryHandler {
	return &InspectionItemsQueryHandler{
		itemsRepository: itemsRepository,
	}
}

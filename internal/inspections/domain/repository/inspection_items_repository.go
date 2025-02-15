package repository

import (
	"context"
	"proj/internal/inspections/domain/entity"
)

type InspectionItemsRepository interface {
	Create(ctx context.Context, item *entity.InspectionItem) error
	GetByID(ctx context.Context, id entity.ID) (*entity.InspectionItem, error)
	//GetByInspectionID(ctx context.Context, inspectionID string) ([]*entity.InspectionItem, error)
	Update(ctx context.Context, item *entity.InspectionItem) error
	Delete(ctx context.Context, id entity.ID) error
	List(ctx context.Context, limit, offset int) ([]*entity.InspectionItem, error)
}

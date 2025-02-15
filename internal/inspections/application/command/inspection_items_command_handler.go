package command

import (
	"context"
	domainCommand "proj/internal/inspections/domain/command"
	"proj/internal/inspections/domain/entity"
	"proj/internal/inspections/domain/repository"
	"proj/pkg/logger"
	"proj/pkg/tracer"
	"time"
)

type InspectionItemsCommandHandler struct {
	itemsRepo repository.InspectionItemsRepository
}

func NewInspectionItemsCommandHandler(itemsRepo repository.InspectionItemsRepository) *InspectionItemsCommandHandler {
	return &InspectionItemsCommandHandler{
		itemsRepo: itemsRepo,
	}
}

func (h *InspectionItemsCommandHandler) HandleCreate(cmd domainCommand.CreateInspectionItemCommand) (*entity.InspectionItem, error) {
	ctx, span := tracer.StartSpan(context.Background(), "InspectionItemsCommandHandler.HandleCreate")
	defer span.End()
	logger.Info("InspectionItemsCommandHandler.HandleCreate")
	item := &entity.InspectionItem{
		Question:  cmd.Question,
		Answer:    cmd.Answer,
		Score:     cmd.Score,
		Comment:   cmd.Comment,
		PhotoURLs: cmd.PhotoUrls,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := h.itemsRepo.Create(ctx, item)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return item, nil
}

//TODO: upd

package handler

import (
	"encoding/json"
	"net/http"
	appCommand "proj/internal/inspections/application/command"
	"proj/internal/inspections/application/dto"
	appQuery "proj/internal/inspections/application/query"
	domainCommand "proj/internal/inspections/domain/command"
	"proj/pkg/validator"
)

type InspectionItemsHandler struct {
	commandHandler *appCommand.InspectionItemsCommandHandler
	queryHandler   *appQuery.InspectionItemsQueryHandler
}

func NewItemsHandler(commandHandler *appCommand.InspectionItemsCommandHandler, queryHandler *appQuery.InspectionItemsQueryHandler) *InspectionItemsHandler {
	return &InspectionItemsHandler{
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
	}
}

func (h *InspectionItemsHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateInspectionItemsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validator.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := domainCommand.CreateInspectionItemCommand{
		Ctx:       r.Context(),
		Question:  req.Question,
		Answer:    req.Answer,
		PhotoUrls: req.PhotoURLs,
		Score:     req.Score,
		Comment:   req.Comment,
	}
	item, err := h.commandHandler.HandleCreate(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

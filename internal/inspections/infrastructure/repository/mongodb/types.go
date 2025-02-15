package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"proj/internal/inspections/domain/entity"
	"time"
)

type MongoInspection struct {
	ID          primitive.ObjectID    `bson:"_id,omitempty"`
	Name        string                `bson:"name"`
	Description string                `bson:"description"`
	Status      string                `bson:"status"`
	CreatedBy   primitive.ObjectID    `bson:"created_by"`
	Items       []MongoInspectionItem `bson:"items"`
	CreatedAt   time.Time             `bson:"created_at"`
	UpdatedAt   time.Time             `bson:"updated_at"`
}

type MongoInspectionItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Question  string             `bson:"question"`
	Answer    string             `bson:"answer"`
	PhotoUrls []string           `bson:"photo_urls"`
	Score     int                `bson:"score"`
	Comment   string             `bson:"comment"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (item *MongoInspectionItem) ToEntity() *entity.InspectionItem {
	return &entity.InspectionItem{
		ID:        item.ID,
		Question:  item.Question,
		Answer:    item.Answer,
		PhotoURLs: item.PhotoUrls,
		Score:     item.Score,
		Comment:   item.Comment,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func (inspection *MongoInspection) ToEntity() *entity.Inspection {
	items := make([]entity.InspectionItem, len(inspection.Items))
	for i, item := range inspection.Items {
		items[i] = *item.ToEntity()
	}
	return &entity.Inspection{
		ID:          inspection.ID,
		Name:        inspection.Name,
		Description: inspection.Description,
		Status:      inspection.Status,
		CreatedBy:   inspection.CreatedBy,
		Items:       items,
		CreatedAt:   inspection.CreatedAt,
		UpdatedAt:   inspection.UpdatedAt,
	}
}

func ItemFromEntity(item *entity.InspectionItem) *MongoInspectionItem {
	id, _ := item.ID.(primitive.ObjectID)
	return &MongoInspectionItem{
		ID:        id,
		Question:  item.Question,
		Answer:    item.Answer,
		PhotoUrls: item.PhotoURLs,
		Score:     item.Score,
		Comment:   item.Comment,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

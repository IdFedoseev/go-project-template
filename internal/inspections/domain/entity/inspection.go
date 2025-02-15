package entity

import (
	"proj/internal/user/domain/entity"
	"time"
)

type ID interface{}

type Inspection struct {
	ID          entity.ID        `json:"id" bson:"_id,omitempty"`
	Name        string           `json:"name" bson:"name,omitempty"`
	Description string           `json:"description" bson:"description,omitempty"`
	Status      string           `json:"status" bson:"status,omitempty"`
	CreatedBy   entity.ID        `json:"created_by" bson:"created_by,omitempty"`
	Items       []InspectionItem `json:"items" bson:"items,omitempty"`
	CreatedAt   time.Time        `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" bson:"updated_at"`
}

type InspectionItem struct {
	ID        entity.ID `json:"id" bson:"_id,omitempty"`
	Question  string    `json:"question" bson:"question,omitempty"`
	Answer    string    `json:"answer" bson:"answer,omitempty"`
	PhotoURLs []string  `json:"photo_urls" bson:"photo_urls,omitempty"`
	Score     int       `json:"score" bson:"score,omitempty"`
	Comment   string    `json:"comment" bson:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"proj/internal/inspections/domain/entity"
	"proj/pkg/tracer"
	"time"
)

type InspectionItemsRepository struct {
	collection *mongo.Collection
}

func NewInspectionItemsRepository(db *mongo.Database) *InspectionItemsRepository {
	return &InspectionItemsRepository{
		collection: db.Collection("inspection_items"),
	}
}

func (r *InspectionItemsRepository) Create(ctx context.Context, item *entity.InspectionItem) error {
	ctx, span := tracer.StartSpan(ctx, "mongodb_inspection_items_repository_Create")
	defer span.End()
	mongoItem := ItemFromEntity(item)
	result, err := r.collection.InsertOne(ctx, mongoItem)
	if err != nil {
		return err
	}
	item.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *InspectionItemsRepository) GetByID(ctx context.Context, id entity.ID) (*entity.InspectionItem, error) {
	ctx, span := tracer.StartSpan(ctx, "mongodb_inspection_items_repository_GetByID")
	defer span.End()
	objID, ok := id.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid ID type")
	}
	var mongoItem MongoInspectionItem
	err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&mongoItem)
	if err != nil {
		return nil, err
	}
	return mongoItem.ToEntity(), nil
}

func (r *InspectionItemsRepository) Update(ctx context.Context, item *entity.InspectionItem) error {
	ctx, span := tracer.StartSpan(ctx, "mongodb_inspection_items_repository_Update")
	defer span.End()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"question":  item.Question,
			"answer":    item.Answer,
			"photoUrls": item.PhotoURLs,
			"score":     item.Score,
			"comment":   item.Comment,
			"updatedAt": time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": item.ID}, update)
	return err
}

func (r *InspectionItemsRepository) Delete(ctx context.Context, id entity.ID) error {
	ctx, span := tracer.StartSpan(ctx, "mongodb_inspection_items_repository_Delete")
	defer span.End()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectID, ok := id.(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("invalid ID type")
	}
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

func (r *InspectionItemsRepository) List(ctx context.Context, limit, offset int) ([]*entity.InspectionItem, error) {
	ctx, span := tracer.StartSpan(ctx, "mongodb_inspection_items_repository_List")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var mongoItems []MongoInspectionItem

	if err = cursor.All(ctx, &mongoItems); err != nil {
		return nil, err
	}
	items := make([]*entity.InspectionItem, len(mongoItems))
	for i, item := range mongoItems {
		items[i] = item.ToEntity()
	}
	return items, nil
}

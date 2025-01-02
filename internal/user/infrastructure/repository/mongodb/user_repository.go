package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"proj/internal/user/domain/entity"
	"proj/pkg/tracer"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	ctx, span := tracer.StartSpan(ctx, "mongodb_create_user")
	defer span.End()

	mongoUser := FromEntity(user)
	result, err := r.collection.InsertOne(ctx, mongoUser)
	if err != nil {
		return err
	}

	user.ID = result.InsertedID
	return nil
}

func (r *UserRepository) GetByID(id entity.ID) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, ok := id.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid ID type")
	}

	var mongoUser MongoUser
	err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&mongoUser)
	if err != nil {
		return nil, err
	}

	return mongoUser.ToEntity(), nil
}

func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var mongoUser MongoUser
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&mongoUser)
	if err != nil {
		return nil, err
	}
	return mongoUser.ToEntity(), nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	ctx, span := tracer.StartSpan(ctx, "mongodb_create_user")
	defer span.End()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"username":  user.Username,
			"email":     user.Email,
			"password":  user.Password,
			"updatedAt": user.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		update,
	)
	return err
}

func (r *UserRepository) Delete(id entity.ID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, ok := id.(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("invalid ID type")
	}

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

func (r *UserRepository) List() ([]*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var mongoUsers []MongoUser
	if err = cursor.All(ctx, &mongoUsers); err != nil {
		return nil, err
	}

	users := make([]*entity.User, len(mongoUsers))
	for i, mu := range mongoUsers {
		users[i] = mu.ToEntity()
	}
	return users, nil
}

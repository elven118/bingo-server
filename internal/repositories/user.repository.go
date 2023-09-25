package repositories

import (
	"bingo/graph/model"
	"bingo/internal/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

// Decode as graphql model.User
func (ur *UserRepository) FindAllPlayersWithName() ([]*model.User, error) {
	var users []*model.User
	filter := bson.D{{Key: "name", Value: bson.D{{Key: "$ne", Value: nil}}}, {Key: "role", Value: models.RolePlayer}}
	options := options.Find().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "name", Value: 1}, {Key: "role", Value: 1}})
	cursor, err := ur.collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) FindUserByID(id string) (*models.User, error) {
	var user models.User
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	err := ur.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindUserByName(name string) (*models.User, error) {
	var user models.User
	filter := bson.M{"name": name}
	err := ur.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindUserByCode(code string) (*models.User, error) {
	var user models.User
	filter := bson.M{"code": code}
	err := ur.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Update(user *models.User) error {
	userReplace, err := ur.collection.ReplaceOne(context.Background(), bson.D{{Key: "_id", Value: user.ID}}, user)
	if err != nil {
		return err
	}
	log.Printf("User replace result: %v", userReplace)
	return nil
}

// func (r *MongoRepository) Add(user *models.User) error {

// }

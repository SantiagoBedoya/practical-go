package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/SantiagoBedoya/reddit_oauth/internal/models"
	"github.com/SantiagoBedoya/reddit_oauth/internal/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepo struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func newMongoClient(dsn string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return client, nil
}

func setupIndexes(collection *mongo.Collection) error {
	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func NewMongoRepo(dsn string, db, collection string) (service.Repository, error) {
	client, err := newMongoClient(dsn)
	if err != nil {
		return nil, err
	}
	database := client.Database(db)
	coll := database.Collection(collection)
	if err := setupIndexes(coll); err != nil {
		return nil, err
	}
	return &mongoRepo{
		client:     client,
		db:         database,
		collection: coll,
	}, nil
}

func (r mongoRepo) Create(ctx context.Context, doc *models.User) error {
	if user, _ := r.FindByEmail(ctx, doc.Email); user != nil {
		return service.ErrEmailInUse
	}
	result, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	doc.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}
func (r mongoRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.D{{
		Key: "email", Value: email,
	}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, service.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

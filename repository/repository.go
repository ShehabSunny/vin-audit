package repository

import (
	"context"
	"time"

	"github.com/shehabsunny/vin-audit/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// CarRepository struct
type CarRepository struct {
	client         *mongo.Client
	CarIdentifier  string
	UserIdentifier string
	config         config.Config
}

func (cr CarRepository) CreateCarMake(makeID string, name string, logoID string) (CarMake, error) {
	carMake := CarMake{
		MakeID: makeID, // uuid.NewV4().String(),
		Name:   name,
		LogoID: logoID,
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := cr.client.Database(cr.config.Database).Collection(cr.config.CarMakeColl)
	_, err := collection.InsertOne(ctx, carMake)
	if err != nil {
		return CarMake{}, err
	}
	return carMake, nil
}

func (cr CarRepository) CreateCarModel(modelID string, name string, logoID string, makeID string) (CarModel, error) {
	carModel := CarModel{
		ModelID: modelID, // uuid.NewV4().String(),
		MakeID:  makeID,
		Name:    name,
		LogoID:  logoID,
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := cr.client.Database(cr.config.Database).Collection(cr.config.CarModelColl)
	_, err := collection.InsertOne(ctx, carModel)
	if err != nil {
		return CarModel{}, err
	}
	return carModel, nil
}

func (cr CarRepository) CreateCarBodyTrim(bodyTrimID string, name string, logoID string, modelID string) (CarBodyTrim, error) {
	carBodyTrim := CarBodyTrim{
		ModelID:    modelID,
		BodyTrimID: bodyTrimID,
		Name:       name,
		LogoID:     logoID,
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := cr.client.Database(cr.config.Database).Collection(cr.config.CarBodyTrimColl)
	_, err := collection.InsertOne(ctx, carBodyTrim)
	if err != nil {
		return CarBodyTrim{}, err
	}
	return carBodyTrim, nil
}

func (cr CarRepository) CreateCarStyle(styleID string, name string, logoID string, bodyTrimID string) (CarStyle, error) {
	carStyle := CarStyle{
		StyleID:    styleID,
		BodyTrimID: bodyTrimID,
		Name:       name,
		LogoID:     logoID,
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := cr.client.Database(cr.config.Database).Collection(cr.config.CarStyleColl)
	_, err := collection.InsertOne(ctx, carStyle)
	if err != nil {
		return CarStyle{}, err
	}
	return carStyle, nil
}

// NewMongoRepository returns a new car repo for mongo database
func NewMongoRepository(dbClient *mongo.Client) CarRepository {
	return CarRepository{
		client:         dbClient,
		config:         config.New(),
		CarIdentifier:  "id",
		UserIdentifier: "userid",
	}
}

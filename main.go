package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/shehabsunny/vin-audit/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Printf("Coul not read file: %v", err)
		return
	}

	var data repository.Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Printf("Could not unmarshal: %v", err)
		return
	}

	cs := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cs))
	if err != nil {
		log.Printf("Could not connect to mongo: %v", err)
		return
	}

	repo := repository.NewMongoRepository(client)

	for _, carMake := range data.Selections.Makes {
		// Add car make
		log.Println("Adding: ", carMake.ID, carMake.Name, "")
		_, err := repo.CreateCarMake(carMake.ID, carMake.Name, "")
		if err != nil {
			log.Printf("Could not add car make: %v", err)
		}

		// Add car model
		for _, carModel := range carMake.Models {
			log.Println("Adding: ", carModel.ID, carMake.Name, "", carMake.ID)
			_, err := repo.CreateCarModel(carModel.ID, carMake.Name, "", carMake.ID)
			if err != nil {
				log.Printf("Could not add car model: %v", err)
			}

			// Add car trim
			for _, carTrim := range carModel.Trims {
				log.Println("Adding: ", carTrim.ID, carTrim.Name, "", carModel.ID)
				_, err := repo.CreateCarBodyTrim(carTrim.ID, carTrim.Name, "", carModel.ID)
				if err != nil {
					log.Printf("Could not add car trim: %v", err)
				}

				// Add car style
				// Add N/A if empty style
				if len(carTrim.Styles) == 0 {
					log.Println("Adding: ", "na", "N/A", "", carTrim.ID)
					_, err := repo.CreateCarStyle("na", "N/A", "", carTrim.ID)
					if err != nil {
						log.Printf("Could not add car style: %v", err)
					}
				} else {
					// Add all styles
					for _, carStyle := range carTrim.Styles {
						log.Println("Adding: ", carStyle.ID, carStyle.Name, "", carTrim.ID)
						_, err := repo.CreateCarStyle(carStyle.ID, carStyle.Name, "", carTrim.ID)
						if err != nil {
							log.Printf("Could not add car style: %v", err)
						}
					}
				}
			}
		}
	}
}

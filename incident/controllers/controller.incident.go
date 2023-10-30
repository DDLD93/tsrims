package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/ddld93/incident/db"
	"github.com/ddld93/incident/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)


type incidenceController struct{}

var dbCollection *mongo.Collection

func init() {
	db := db.StartMongo().Database("testdb")
	dbCollection = db.Collection("incident")
}

func NewIncidentController() *incidenceController {
	return &incidenceController{}
}

func (*incidenceController) AddIncident(incident *models.Incident) (*models.Incident, error) {
	incident.CreatedAt = time.Now()
	result, err := dbCollection.InsertOne(context.Background(), incident)
	if err != nil {
		return nil, err
	}
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	incident.ID = insertedID
	return incident, nil
}

func (*incidenceController) GetAllIncident() ([]*models.Incident, error) {
	cursor, err := dbCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var incidents []*models.Incident
	if err := cursor.All(context.Background(), &incidents); err != nil {
		return nil, err
	}

	return incidents, nil
}

func (uc *incidenceController) GetIncidentByID(incidentID string) (*models.Incident, error) {
	var incident models.Incident
	id,_:= primitive.ObjectIDFromHex(incidentID)
	err := dbCollection.FindOne(context.Background(), bson.M{"_id":id}).Decode(&incident)
	if err != nil {
		return nil, err
	}

	return &incident, nil
}

func (uc *incidenceController) UpdateIncident(incidentID string, updatedIncident *models.Incident) (*models.Incident, error) {
    // Parse the incidentID into an ObjectID
    id, err := primitive.ObjectIDFromHex(incidentID)
    if err != nil {
        return nil, err
    }

    // Define an update operation to replace the existing incident with the updated incident
    update := bson.M{"$set": updatedIncident}

    // Find the incident by ID and update it
    filter := bson.M{"_id": id};
    result, err := dbCollection.UpdateOne(context.Background(), filter, update)
	
	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no documents were updated")
	}
    if err != nil {
        return nil, err
	}
	return updatedIncident ,nil
}

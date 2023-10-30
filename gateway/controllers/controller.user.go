package controllers

import (
	"context"
	"fmt"

	"github.com/ddld93/tsrims/auth/db"
	"github.com/ddld93/tsrims/auth/models"
	"github.com/ddld93/tsrims/auth/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type userController struct{}

var dbCollection *mongo.Collection

func init() {
	db := db.StartMongo().Database("testdb")
	dbCollection = db.Collection("users")
}

func NewCustomerController() *userController {
	return &userController{}
}

func (*userController) Register(user *models.User) (*models.User, error) {

	existingUser := &models.User{}
	err := dbCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(existingUser)
	if err == nil {
		return nil, fmt.Errorf("user with this email already exists")
	} else if err != mongo.ErrNoDocuments {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	result, err := dbCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	user.ID = insertedID
	return user, nil
}

func (*userController) Login(email, password string) (*models.User, string, error) {
	user := models.User{}
	err := dbCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", err
	}

	token, err := utils.GenerateJWTToken(&user)
	if err != nil {
		return nil, "", err
	}

	// Create a response object with both user and token

	return &user, token, nil
}

func (*userController) GetAllUsers() ([]*models.User, error) {
	cursor, err := dbCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []*models.User
	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *userController) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	id,_:= primitive.ObjectIDFromHex(userID)
	err := dbCollection.FindOne(context.Background(), bson.M{"_id":id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

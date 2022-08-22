package service

import (
	"context"
	"encoding/json"
	"grocerylist/database"
	"grocerylist/service/assistants"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Data structure to recieve householdID
type DataForList struct {
	HouseholdId string `json:"HouseholdId"`
}

// Data obejct to handle the mongo drivers generation of primitive objectIds
type MyObjectID string

// data structure to populate and insert into the mongo db
type CreateList struct {
	ID          MyObjectID              `bson:"_id"`
	HouseholdId string                  `bson:"HouseholdId, omitempty"`
	Items       []assistants.CreateItem `bson:"Items"`
}

// CreateList has to be a post recieving a json object with HouseholdId, the house hold must now have a list beforehand.
// The function returns a json object with ListId, which can be used to retrieve data from the list in the future.
func (s *Server) CreateList(w http.ResponseWriter, r *http.Request) {

	//In any case return a json format
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//If the method is not post, return bad requst
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "Wrong method"}`)
		return
	}

	//Get the json body of the post and populate the dataforlist structure
	var cl DataForList
	err := json.NewDecoder(r.Body).Decode(&cl)
	if err != nil {
		str := `{"Error": "Bad request"}`
		w.WriteHeader(400)
		io.WriteString(w, str)
		return
	}

	if cl.HouseholdId == "" {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "No household provided"}`)
		return
	}

	//Instantiate a connection to mongo
	var client database.MongClient
	client.DbConnect(database.ConstGroceryListCollection)

	//Check if the household allready has a grocery list
	lookfor := cl.HouseholdId
	filter := bson.D{primitive.E{Key: "HouseholdId", Value: lookfor}}
	var results bson.M

	//Checking if there is any matches on the house hold id, if so return 400
	_ = client.Connection.FindOne(context.TODO(), filter).Decode(&results)
	if results != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Household already has a list"}`)
		return
	}

	//Household doesnt have a list create one
	newId := primitive.NewObjectID()

	cList := CreateList{MyObjectID(newId.Hex()), cl.HouseholdId, nil}

	_, err = client.Connection.InsertOne(context.TODO(), cList)
	if err != nil {
		io.WriteString(w, `{"Error": "Failed creating a list"}`)
		w.WriteHeader(400)
		return
	}
	defer client.DbDisconnect()

	//Format the return data and serve as json.
	type ReturnData struct {
		Succes string     `json:"Succes"`
		ListId MyObjectID `json:"ListId"`
	}
	returnDataFormat := ReturnData{"List Created", cList.ID}
	json.NewEncoder(w).Encode(returnDataFormat)
}

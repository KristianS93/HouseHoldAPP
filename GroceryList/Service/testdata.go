package service

import (
	"context"
	"fmt"
	"grocerylist/database"
	"grocerylist/service/assistants"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var TestItems = []assistants.CreateItem{}

var TestItemID string

func AddTestData() {
	//Create test DB items collection.
	var ItemClient database.MongClient
	ItemClient.DbConnect(database.ConstGroceryItemsCollection)

	//Create test DB items collection.
	var ListClient database.MongClient
	ListClient.DbConnect(database.ConstGroceryListCollection)

	// To test getList we will need a list
	// and a items.

	cList := CreateList{"62fa8c527abec12155c907c3", "TestHouse", nil}

	_, err := ListClient.Connection.InsertOne(context.TODO(), cList)
	if err != nil {
		fmt.Println("failed inserting")
		return
	}
	//insert items to test db.

	obj1 := assistants.CreateItem{ID: "", ListId: "62fa8c527abec12155c907c3", ItemName: "Test item1", Quantity: "4", Unit: "pakker"}
	obj2 := assistants.CreateItem{ID: "", ListId: "62fa8c527abec12155c907c3", ItemName: "Test item2", Quantity: "5", Unit: "stk"}

	TestItems = append(TestItems, obj1)
	TestItems = append(TestItems, obj2)

	var itemInsertFormat []assistants.CreateItem

	for _, v := range TestItems {
		newId := primitive.NewObjectID()
		insertObj := assistants.CreateItem{ID: string(newId.Hex()), ListId: v.ListId, ItemName: v.ItemName, Quantity: v.Quantity, Unit: v.Unit}
		itemInsertFormat = append(itemInsertFormat, insertObj)
	}

	var insertItemQuery []interface{}
	for _, v := range itemInsertFormat {
		insertItemQuery = append(insertItemQuery, v)
	}

	_, err = ItemClient.Connection.InsertMany(context.TODO(), insertItemQuery)
	if err != nil {
		fmt.Println("error inserting  test items")
		return
	}

	//Getting item id for test
	filter := bson.D{primitive.E{Key: "ListId", Value: "62fa8c527abec12155c907c3"}}

	res, err := ItemClient.Connection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("No items on the list")
		return
	}
	defer ListClient.DbDisconnect()

	var itemsList []ItemList
	if err = res.All(context.TODO(), &itemsList); err != nil {
		fmt.Println("Could not retrieve list")
	}
	TestItemID = itemsList[0].ID
}

func DeleteTestData() {
	//Create test DB items collection.
	var ListClient database.MongClient
	ListClient.DbConnect(database.ConstGroceryListCollection)

	filter := bson.D{{Key: "HouseholdId", Value: "TestId"}}
	_, err := ListClient.Connection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("error deleting test list")
		return
	}
	defer ListClient.DbDisconnect()

}

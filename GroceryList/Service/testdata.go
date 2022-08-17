package service

import (
	"context"
	"fmt"
	"grocerylist/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var TestItems = []CreateItem{}

func AddTestData() {
	//Create test DB items collection.
	var ItemClient database.MongClient
	ItemClient.DbConnect(database.ConstGroceryItemsCollection)

	//Create test DB items collection.
	var ListClient database.MongClient
	ListClient.DbConnect(database.ConstGroceryListCollection)
	// ListClient.DbDisconnect()

	// To test getList we will need a list
	// and a items.

	cList := CreateList{"62fa8c527abec12155c907c3", "TestHouse", nil}

	_, err := ListClient.Connection.InsertOne(context.TODO(), cList)
	if err != nil {
		fmt.Println("failed inserting")
		return
	}

	//insert items to test db.

	obj1 := CreateItem{"", "62fa8c527abec12155c907c3", "Test item1", "4", "pakker"}
	obj2 := CreateItem{"", "62fa8c527abec12155c907c3", "Test item2", "5", "stk"}

	TestItems = append(TestItems, obj1)
	TestItems = append(TestItems, obj2)

	var itemInsertFormat []CreateItem

	for _, v := range TestItems {
		newId := primitive.NewObjectID()
		insertObj := CreateItem{string(newId.Hex()), v.ListId, v.ItemName, v.Quantity, v.Unit}
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

}

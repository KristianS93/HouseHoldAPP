package service

import (
	"context"
	"fmt"
	"grocerylist/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

	cList := CreateList{"62fa8c527abec12155c907c3", "testhouse", nil}

	_, err := ListClient.Connection.InsertOne(context.TODO(), cList)
	if err != nil {
		fmt.Println("failed inserting")
		return
	}

	//insert items to test db.

	var itemformat []CreateItem
	obj1 := CreateItem{"", "62fa8c527abec12155c907c3", "Test item1", "4", "pakker"}
	obj2 := CreateItem{"", "62fa8c527abec12155c907c3", "Test item2", "5", "stk"}

	itemformat = append(itemformat, obj1)
	itemformat = append(itemformat, obj2)

	var itemInsertFormat []CreateItem

	for _, v := range itemformat {
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
		fmt.Println("error inserting items")
		return
	}
}

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"grocerylist/database"
	"grocerylist/service/assistants"
	"net/http"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ItemClient = database.MongClient{}

//additem
//changeitem
//deleteitem

var ServerIns = Server{}

func TestMain(m *testing.M) {

	//Initialize server settings¢
	ServerIns.Init()

	//Db connection
	//Create test DB items collection.
	//addint test data
	AddTestData()

	code := m.Run()
	DeleteTestData()
	//clearData()
	os.Exit(code)
}

func TestAddItem(t *testing.T) {

	type AddItemTests teststruct

	testcasesStatusCode := []AddItemTests{
		{"Add single item correct request, output 200", ExtraData{"POST", "/AddItem", `[{"ListId":"62fa8c527abec12155c907c3", "ItemName":"TestItem1", "Quantity":"3", "Unit":"Pcs"}]`}, http.StatusOK},
		{"Add multiple item correct request, output 200", ExtraData{"POST", "/AddItem", `[{"ListId":"62fa8c527abec12155c907c3", "ItemName":"TestItem1", "Quantity":"3", "Unit":"Pcs"}, {"ListId":"62fa8c527abec12155c907c3", "ItemName":"TestItem2", "Quantity":"5", "Unit":"Pcs"}]`}, http.StatusOK},
		{"Wrong method in request, output 405", ExtraData{"GET", "/AddItem", `[{"ListId":"62fa8c527abec12155c907c3", "ItemName":"TestItem1", "Quantity":"3", "Unit":"Pcs"}]`}, http.StatusMethodNotAllowed},
		{"Wrong json in request, output 400", ExtraData{"POST", "/AddItem", `{"ListId":"62fa8c527abec12155c907c3", "ItemName":"TestItem1", "Quantity":"3", "Unit":"Pcs"}`}, http.StatusBadRequest},
		{"No json, output 400", ExtraData{"POST", "/AddItem", `{}`}, http.StatusBadRequest},
		{"Missing ListId, output 400", ExtraData{"POST", "/AddItem", `[{"ListId":"", "ItemName":"TestItem1", "Quantity":"3", "Unit":"Pcs"}]`}, http.StatusBadRequest},
		{"missing ItemName, output 400", ExtraData{"POST", "/AddItem", `[{"ListId":"62fa8c527abec12155c907c3", "ItemName":"", "Quantity":"3", "Unit":"Pcs"}]`}, http.StatusBadRequest},
	}

	for _, v := range testcasesStatusCode {
		req, _ := http.NewRequest(v.Input.method, v.Input.url, CreateReader(v.Input.body))
		res := ExecuteRequest(req)
		CheckResCode(t, v.Output, res.Code, v.Name)
	}

	fmt.Println("Passed Add item tests")

}

func TestChangeItem(t *testing.T) {

	type ChangeItemTests teststruct
	te := assistants.ItemData{Id: TestItemID, ItemName: "Changed name", Quantity: "3", Unit: "stk"}
	wr := assistants.ItemData{Id: "", ItemName: "Changed name", Quantity: "3", Unit: "stk"}
	tr := assistants.ItemData{Id: "423424323", ItemName: "Changed name", Quantity: "3", Unit: "stk"}

	genJson, _ := json.Marshal(te)
	genjson2, _ := json.Marshal(wr)
	genJson3, _ := json.Marshal(tr)

	testcasesStatusCode := []ChangeItemTests{
		{"Change item correct request, output 200", ExtraData{"PATCH", "/ChangeItem", string(genJson)}, http.StatusOK},
		{"Wrong method request, output 405", ExtraData{"GET", "/ChangeItem", string(genJson)}, http.StatusMethodNotAllowed},
		{"Json fail request, output 400", ExtraData{"PATCH", "/ChangeItem", ``}, http.StatusBadRequest},
		{"No item id, output 400", ExtraData{"PATCH", "/ChangeItem", string(genjson2)}, http.StatusBadRequest},
		{"Non existing item id, output 400", ExtraData{"PATCH", "/ChangeItem", string(genJson3)}, http.StatusBadRequest},
	}

	for _, v := range testcasesStatusCode {
		req, _ := http.NewRequest(v.Input.method, v.Input.url, CreateReader(v.Input.body))
		res := ExecuteRequest(req)
		CheckResCode(t, v.Output, res.Code, v.Name)
	}

	var ItemClient database.MongClient
	ItemClient.DbConnect(database.ConstGroceryItemsCollection)

	filter := bson.D{primitive.E{Key: "_id", Value: TestItemID}}
	var yt ItemList
	_ = ItemClient.Connection.FindOne(context.TODO(), filter).Decode(&yt)
	if yt.ItemName != "Changed name" {
		t.Errorf("Test: Didnt change - Expected item %s. Got %s \n", "Changed name", yt.ItemName)

	}
	defer ItemClient.DbDisconnect()

	fmt.Println("Passed ChangeItem tests")

}

func TestDeleteItem(t *testing.T) {

	type DeleteItemTests teststruct

	gt := assistants.RecieveId{ListId: "", ItemId: TestItemID}
	gt2 := assistants.RecieveId{ListId: "", ItemId: ""}
	gt3 := assistants.RecieveId{ListId: "", ItemId: "34534523"}

	genJson1, _ := json.Marshal(gt)
	genJson2, _ := json.Marshal(gt2)
	genJson3, _ := json.Marshal(gt3)

	testcasesStatusCode := []DeleteItemTests{
		{"Wrong method request, output 405", ExtraData{"GET", "/DeleteItem", string(genJson1)}, http.StatusMethodNotAllowed},
		{"Wrong jsonrequest, output 400", ExtraData{"DELETE", "/DeleteItem", ""}, http.StatusBadRequest},
		{"Missing ItemId output 400", ExtraData{"DELETE", "/DeleteItem", string(genJson2)}, http.StatusBadRequest},
		{"Non existing ItemId output 400", ExtraData{"DELETE", "/DeleteItem", string(genJson3)}, http.StatusBadRequest},
	}

	for _, v := range testcasesStatusCode {
		req, _ := http.NewRequest(v.Input.method, v.Input.url, CreateReader(v.Input.body))
		res := ExecuteRequest(req)
		CheckResCode(t, v.Output, res.Code, v.Name)
	}

	var ItemClient database.MongClient
	ItemClient.DbConnect(database.ConstGroceryItemsCollection)

	filter := bson.D{{Key: "_id", Value: TestItemID}}

	_, err := ItemClient.Connection.DeleteOne(context.TODO(), filter)
	if err != nil {
		t.Errorf("Error deleting")
	}

	var yt ItemList
	_ = ItemClient.Connection.FindOne(context.TODO(), filter).Decode(&yt)
	if yt.ID != "" {
		t.Errorf("Error deleting")
	}
	defer ItemClient.DbDisconnect()

	fmt.Println("Passed DeleteItem tests")
}

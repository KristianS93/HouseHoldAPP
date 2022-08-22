package assistants

type ItemData struct {
	Id       string `bson:"_id"`
	ItemName string `bson:"ItemName" json:"ItemName"`
	Quantity string `bson:"Quantity" json:"Quantity"`
	Unit     string `bson:"Unit" json:"Unit"`
}

type CreateItem struct {
	ID       string `bson:"_id, omitempty"`
	ListId   string `bson:"ListId, omitempty" json:"ListId"`
	ItemName string `bson:"ItemName, omitempty" json:"ItemName"`
	Quantity string `bson:"Quantity" json:"Quantity"`
	Unit     string `bson:"Unit" json:"Unit"`
}

type RecieveId struct {
	ListId string `json:"ListId"`
	ItemId string `json:"ItemId"`
}

type ItemHolder struct {
	Succes string     `json:"Succes"`
	Items  []ItemData `json:"Items"`
}

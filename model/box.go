package model

type Box struct {
	Id           string `json:"id" bson:"_id,omitempty"`
	ChainAddress string `json:"chainAddress" bson:"chainAddress"`
	CDN          string `json:"cdn"`
	Prizes       []interface{}
}

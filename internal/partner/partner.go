package partner

type Partner struct {
	ID           string `json:"id" bson:"id"`
	TradingName  string `json:"tradingName" bson:"tradingName"`
	OwnerName    string `json:"ownerName" bson:"ownerName"`
	Document     string `json:"document" bson:"document"`
	CoverageArea struct {
		Type        string          `json:"type" bson:"type"`
		Coordinates [][][][]float64 `json:"coordinates" bson:"coordinates"`
	} `json:"coverageArea" bson:"coverageArea"`
	Address struct {
		Type        string    `json:"type" bson:"type"`
		Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	} `json:"address" bson:"address"`
}

package partner

type Partner struct {
	Id           string `json:"id" bson:"_id"`
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

type PartnerPostRequest struct {
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

type Service interface {
	CreatePartner(partner PartnerPostRequest) (string, error)
	GetPartnerById(id string) (Partner, error)
	GetPartnerByLatLong(latLong []float64) (Partner, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreatePartner(req PartnerPostRequest) (string, error) {
	return s.repo.CreatePartner(req)
}

func (s *service) GetPartnerById(id string) (Partner, error) {
	return s.repo.GetById(id)
}

func (s *service) GetPartnerByLatLong(latLong []float64) (Partner, error) {
	return s.repo.GetNearPartner(latLong)
}

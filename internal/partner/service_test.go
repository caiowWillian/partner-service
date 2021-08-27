package partner_test

import (
	"testing"

	"github.com/caiowWillian/partner-service/internal/partner"
	"github.com/caiowWillian/partner-service/internal/partner/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockRepository *mocks.Repository
	service        partner.Service
)

func init() {
	mockRepository = new(mocks.Repository)
	service = partner.NewService(mockRepository)
}

func TestCreatePartner(t *testing.T) {
	mockRepository.On("CreatePartner", mock.Anything).Return("id", nil)

	_, err := service.CreatePartner(partner.PartnerPostRequest{})
	assert.Nil(t, err)
}

func TestGetPartnerById(t *testing.T) {
	p := partner.Partner{Id: "123"}
	mockRepository.On("GetById", mock.Anything).Return(p, nil)

	result, err := service.GetPartnerById("123")
	assert.Nil(t, err)
	assert.Equal(t, "123", result.Id)
}

func TestGetPartnerByLatLong(t *testing.T) {
	p := partner.Partner{Id: "123"}
	mockRepository.On("GetNearPartner", mock.Anything).Return(p, nil)
	result, err := service.GetPartnerByLatLong([]float64{1, 2})
	assert.Nil(t, err)
	assert.Equal(t, "123", result.Id)

}

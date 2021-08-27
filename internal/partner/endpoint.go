package partner

import (
	"context"

	"github.com/caiowWillian/partner-service/pkg/encodedError"
	"github.com/go-kit/kit/endpoint"
)

func makeCreatePartner(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id, err := s.CreatePartner(request.(PartnerPostRequest))

		if err != nil {
			return nil, err
		}

		return createPartnerPostResponse{Id: id, StatusCode: 201}, nil
	}
}

func makeGetPartnerById(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		p, err := s.GetPartnerById(request.(string))

		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				return nil, encodedError.NoContent
			}

			return nil, err
		}

		return p, nil
	}
}

func makeGetPartnerByLatLong(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		latlong := request.([]float64)
		p, err := s.GetPartnerByLatLong(latlong)

		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				return nil, encodedError.NoContent
			}
		}
		return p, nil
	}
}

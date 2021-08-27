package partner

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/caiowWillian/partner-service/pkg/encodedError"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type createPartnerPostResponse struct {
	Id         string `json:"id"`
	StatusCode int    `json:"-"`
}

func NewHTTPServer(ctx context.Context, s Service, router *mux.Router) http.Handler {
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodedError.EncodeError),
	}

	router.Use(commonMiddleware)

	router.Methods(http.MethodPost).Path("/partner").Handler(httptransport.NewServer(
		makeCreatePartner(s),
		decodeCreatePartnerReq,
		encodeCreatePartnerResp,
		opts...,
	))

	router.Methods(http.MethodGet).Path("/partner/{id}").Handler(httptransport.NewServer(
		makeGetPartnerById(s),
		decodeGetPartnerByIdReq,
		encodeGetPartnerResp,
		opts...,
	))

	router.Methods(http.MethodGet).Path("/partner").Handler(httptransport.NewServer(
		makeGetPartnerByLatLong(s),
		decodeGetPartnerByLatLongReq,
		encodeGetPartnerResp,
		opts...,
	))

	return router
}
func decodeCreatePartnerReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req PartnerPostRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, encodedError.BadRequest
	}
	return req, nil
}

func decodeGetPartnerByLatLongReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var lat float64
	var long float64
	var err error

	lat, err = strconv.ParseFloat(r.URL.Query()["lat"][0], 64)

	if err != nil {
		return nil, err
	}

	long, err = strconv.ParseFloat(r.URL.Query()["long"][0], 64)

	if err != nil {
		return nil, err
	}

	latLong := []float64{lat, long}
	return latLong, nil
}

func decodeGetPartnerByIdReq(ctx context.Context, r *http.Request) (interface{}, error) {
	path := strings.Split(r.URL.Path, "/")

	return path[len(path)-1], nil
}

func encodeGetPartnerResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeCreatePartnerResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(response.(createPartnerPostResponse).StatusCode)
	return json.NewEncoder(w).Encode(response)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

package route

import (
	"context"

	"github.com/caiowWillian/partner-service/internal/partner"
	"github.com/caiowWillian/partner-service/pkg/mongo"
	"github.com/gorilla/mux"
)

func MakeRoutes(ctx context.Context, router *mux.Router) {
	partner.NewHTTPServer(ctx, partner.NewService(partner.NewRepository(mongo.NewMongo())), router)
}

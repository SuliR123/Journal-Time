package server

import (
	"net/http"

	health "github.com/SuliR123/Journal-Time/internal/handlers/health"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateApp(connection *pgxpool.Pool) *http.ServeMux {

	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Jounral Time API", "1.0.0"))

	// Create all the routing groups:
	health.Route(api, connection)

	return mux
}

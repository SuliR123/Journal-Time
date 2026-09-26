package health

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Route(api huma.API, connection *pgxpool.Pool) {
	var service HealthService = *newService(connection)

	{
		group := huma.NewGroup(api, "/api/v1/health")
		huma.Register(group, huma.Operation{
			OperationID: "check-health",
			Method:      http.MethodGet,
			Path:        "/",
			Summary:     "Check health of server",
			Tags:        []string{"Health"},
		}, service.CheckHealth)
	}
}

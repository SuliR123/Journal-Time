package server

import (
	"github.com/SuliR123/Journal-Time/internal/handlers"
	health "github.com/SuliR123/Journal-Time/internal/handlers/Health"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateApp(connection *pgxpool.Pool) *gin.Engine {

	router := gin.Default()

	// Create all the routing groups:
	var routeGroups []handlers.RouteFN = []handlers.RouteFN{health.Route}
	for _, fn := range routeGroups {
		fn(router, connection)
	}

	return router
}

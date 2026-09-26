package health

import (
	"context"

	"github.com/SuliR123/Journal-Time/internal/handlers/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthService struct {
	healthDB *HealthDB
}

func newService(connection *pgxpool.Pool) *HealthService {
	return &HealthService{healthDB: newDB(connection)}
}

func (h *HealthService) CheckHealth(ctx context.Context, input *utils.BlankInput) (*HealthOutput, error) {
	resp := &HealthOutput{}
	resp.Body.Message = "AYYY WE GOOD"
	return resp, nil
}

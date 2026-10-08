package dashboard

import (

	"github.com/jackc/pgx/v5/pgxpool"
)

type DashboardService struct {
	db *pgxpool.Pool
}

func NewDashboardService(db *pgxpool.Pool) *DashboardService {
	return &DashboardService{
		db: db,
	}
}

func (d *DashboardService) GetAll() {
	
}
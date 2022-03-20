package main

import (
	"context"
	"database/sql"
)

const findServicesForUser = `
SELECT service_id, title, summary, s.org_id as org_id
FROM services s 
JOIN users u
ON s.org_id = u.org_id
WHERE u.user_id = ?
`

type Service struct {
	ServiceID int    `json:"serviceId"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	OrgID     int    `json:"orgId"`
}

type DataService interface {
	FindServicesForUser(ctx context.Context, userID int) ([]Service, error)
}

type SQLDataService struct {
	db *sql.DB
}

func (s *SQLDataService) FindServicesForUser(ctx context.Context, userID int) ([]Service, error) {
	rows, err := s.db.Query(findServicesForUser, userID)
	if err != nil {
		return nil, err
	}

	services := make([]Service, 0)
	for rows.Next() {
		var svc Service
		err = rows.Scan(&svc.ServiceID, &svc.Title, &svc.Summary, &svc.OrgID)
		if err != nil {
			return nil, err
		}
		services = append(services, svc)
	}

	return services, nil
}

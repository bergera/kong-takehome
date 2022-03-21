package main

import (
	"context"
	"database/sql"
)

const findServicesForUser = `
SELECT s.service_id, s.title, s.summary, s.org_id, COUNT(v.version_id) as version_count
FROM services s 
JOIN users u ON s.org_id = u.org_id
JOIN versions v ON s.service_id = v.service_id
WHERE u.user_id = $1
GROUP BY s.service_id
`

type Service struct {
	ServiceID    int    `json:"serviceId"`
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	OrgID        int    `json:"orgId"`
	VersionCount int    `json:"versionCount"`
}

type DataService interface {
	FindServicesForUser(ctx context.Context, userID string) ([]Service, error)
}

type SQLDataService struct {
	db *sql.DB
}

func (s *SQLDataService) FindServicesForUser(ctx context.Context, userID string) ([]Service, error) {
	rows, err := s.db.Query(findServicesForUser, userID)
	if err != nil {
		return nil, err
	}

	services := make([]Service, 0)
	for rows.Next() {
		var svc Service
		err = rows.Scan(&svc.ServiceID, &svc.Title, &svc.Summary, &svc.OrgID, &svc.VersionCount)
		if err != nil {
			return nil, err
		}
		services = append(services, svc)
	}

	return services, nil
}

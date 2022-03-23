package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

const findServicesSQL = `
SELECT s.service_id, s.title, s.summary, s.org_id, COUNT(v.version_id) as version_count
FROM services s 
JOIN users u ON s.org_id = u.org_id
JOIN versions v ON s.service_id = v.service_id
WHERE u.user_id = $1
GROUP BY s.service_id
LIMIT $2 OFFSET $3
`

const getServiceSQL = `
SELECT s.service_id, s.title, s.summary, s.org_id
FROM services s 
JOIN users u ON s.org_id = u.org_id
WHERE u.user_id = $1 AND s.service_id = $2
`

const findVersionsForServiceSQL = `
SELECT v.service_id, v.version_id, v.summary
FROM versions v
JOIN services s ON s.service_id = v.service_id
JOIN users u ON s.org_id = u.org_id
WHERE u.user_id = $1 AND v.service_id = $2
LIMIT $3 OFFSET $4
`

const findVersionSQL = `
SELECT v.service_id, v.version_id, v.summary
FROM versions v
JOIN services s ON s.service_id = v.service_id
JOIN users u ON s.org_id = u.org_id
WHERE u.user_id = $1 AND v.service_id = $2 AND v.version_id = $3
`

type Service struct {
	ServiceID    string `json:"serviceId"`
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	OrgID        string `json:"orgId"`
	VersionCount int    `json:"versionCount"`
}

type Version struct {
	ServiceID string `json:"serviceId"`
	VersionID string `json:"versionId"`
	Summary   string `json:"summary"`
}

type DataService interface {
	FindServices(ctx context.Context, limit, offset int) ([]Service, error)
	FindServiceByID(ctx context.Context, serviceID string) (*Service, error)
	FindVersionsForService(ctx context.Context, serviceID string, limit, offset int) ([]Version, error)
	FindVersionByID(ctx context.Context, serviceID string, versionID string) (*Version, error)
}

type SQLDataService struct {
	db *sql.DB
}

func (s *SQLDataService) FindServices(ctx context.Context, limit, offset int) ([]Service, error) {
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		fmt.Println("user ID not found in context")
		return nil, errors.New("user ID not found in context")
	}

	rows, err := s.db.QueryContext(ctx, findServicesSQL, userID, limit, offset)
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
	if err = rows.Err(); err != nil {
		return nil, err
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}

	return services, nil
}

func (s *SQLDataService) FindServiceByID(ctx context.Context, serviceID string) (*Service, error) {
	var svc Service

	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		fmt.Println("user ID not found in context")
		return nil, errors.New("user ID not found in context")
	}

	err := s.db.QueryRowContext(ctx, getServiceSQL, userID, serviceID).Scan(&svc.ServiceID, &svc.Title, &svc.Summary, &svc.OrgID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &svc, err
}

func (s *SQLDataService) FindVersionByID(ctx context.Context, serviceID string, versionID string) (*Version, error) {
	var v Version

	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		fmt.Println("user ID not found in context")
		return nil, errors.New("user ID not found in context")
	}

	err := s.db.QueryRowContext(ctx, findVersionSQL, userID, serviceID, versionID).Scan(&v.ServiceID, &v.VersionID, &v.Summary)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &v, err
}

func (s *SQLDataService) FindVersionsForService(ctx context.Context, serviceID string, limit, offset int) ([]Version, error) {
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		fmt.Println("user ID not found in context")
		return nil, errors.New("user ID not found in context")
	}

	rows, err := s.db.QueryContext(ctx, findVersionsForServiceSQL, userID, serviceID, limit, offset)
	if err != nil {
		return nil, err
	}

	versions := make([]Version, 0)
	for rows.Next() {
		var v Version
		err = rows.Scan(&v.ServiceID, &v.VersionID, &v.Summary)
		if err != nil {
			return nil, err
		}
		versions = append(versions, v)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}

	return versions, nil
}

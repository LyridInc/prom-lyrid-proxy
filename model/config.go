package model

import "time"

type TenantConfig struct {
	ID          string
	Name        string
	Description string

	CreationTime time.Time
}

type TenantKey struct {
	TenantID string
	Secret   string
}

type PrometheusConfig struct {
	ID  string
	URL string
}

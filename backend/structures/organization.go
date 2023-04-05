package structures

import "time"

type Organization struct {
	Id              int        `json:"id"`
	Name            string     `json:"name"`
	Address         string     `json:"address"`
	CreatedAt       *time.Time `json:"created_at"`
	DealAmount      int        `json:"deal_amount"`
	DaysTillRenewal int        `json:"days_till_renewal"`
}

type HoursOnProduct struct {
	Hourrange string `json:"hourrange"`
	Count     int    `json:"count"`
}

type OrganizationData struct {
	Organization
	HoursOnProductHistogramData []HoursOnProduct `json:"hours_on_product"`
}

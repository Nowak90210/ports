package infrastructure

import "github.com/Nowak90210/ports/internal/domain"

type PortDto struct {
	ID          string
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func (d *PortDto) ToDomain() *domain.Port {
	return domain.NewPort(d.ID, d.Name, d.City, d.Province, d.Country, d.Timezone, d.Code, d.Alias, d.Regions, d.Unlocs, d.Coordinates)
}

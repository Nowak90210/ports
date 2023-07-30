package domain

type Port struct {
	ID          string
	name        string
	coordinates []float64
	city        string
	province    string
	country     string
	alias       []string
	regions     []string
	timezone    string
	unlocs      []string
	code        string
}

func NewPort(ID, name, city, province, country, timezone, code string, alias, regions, unlocs []string, coordinates []float64) *Port {
	port := Port{
		ID:          ID,
		name:        name,
		coordinates: coordinates,
		city:        city,
		province:    province,
		country:     country,
		alias:       alias,
		regions:     regions,
		timezone:    timezone,
		unlocs:      unlocs,
		code:        code,
	}

	return &port
}

package domain

// Port represents a domain model
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

func (p *Port) Name() string {
	return p.name
}

func (p *Port) City() string {
	return p.city
}

func (p *Port) Province() string {
	return p.province
}

func (p *Port) Country() string {
	return p.country
}

func (p *Port) Alias() []string {
	return p.alias
}

func (p *Port) Regions() []string {
	return p.regions
}

func (p *Port) Coordinates() []float64 {
	return p.coordinates
}

func (p *Port) Timezone() string {
	return p.timezone
}

func (p *Port) Unlocs() []string {
	return p.unlocs
}

func (p *Port) Code() string {
	return p.code
}

package models

type Apartment struct {
	apartment_id   string
	apartment_name string
	address        string
	city           string
	state          string
	zip_code       string
}

type Unit struct {
	unit_id        string
	apartment_id   string
	unit_number    string
	unit_type      string
	square_footage string
	rent           float64
	status         string
	tenant_id      string
}

package utils

type Diag struct {
	UniApi       int    `json:"universitiesapi"`
	CountriesApi int    `json:"countriesapi"`
	Version      string `json:"version"`
	Uptime       string `json:"uptime"`
}

type Country struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Borders []string `json:"borders"`
}

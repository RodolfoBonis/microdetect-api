package entities

// Storage represents storage information.
// @Description Storage data
// @Example {"Used": "200GB", "Total": "500GB", "Percentage": "40%"}

type Storage struct {
	Used       string
	Total      string
	Percentage string
}

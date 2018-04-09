package locator

type (
	Location struct {
		LocationID   string `json:"location_id"`
		Country      string `json:"country"`
		Subdivision1 string `json:"subdivision_1"`
		Subdivision2 string `json:"subdivision_2"`
		City         string `json:"city"`
	}

	Network struct {
		LocationID string   `json:"location_id"`
		Mask       string   `json:"mask"`
		Location   *Location `json:"location"`
	}
)

func (l Location) ToMapInterface() map[string]interface{} {
	m := make(map[string]interface{})
	m["country"] = l.Country
	m["city"] = l.City
	m["subdivision_1"] = l.Subdivision1
	m["subdivision_2"] = l.Subdivision2
	m["location_id"] = l.LocationID
	return m
}

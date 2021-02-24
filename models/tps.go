package models
type (
	TPS struct {
		Id int `json:"id"`
		No_TPS string `json:"no_tps"`
		Lokasi string `json:"lokasi"`
		Kecamatan string `json:"kecamatan"`
		Nagari string `json:"nagari"`
		Jorong string `json:"jorong"`
		JPL int `json:"jpl"`
		Is_active int `json:"is_active"`
	}
)


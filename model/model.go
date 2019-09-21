package model

type Driver struct {
	ID   string  `json:"id"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type ReqDriver struct {
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
	Limit int     `json:"limit"`
}

type MassageStatus struct {
	Status      string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
}

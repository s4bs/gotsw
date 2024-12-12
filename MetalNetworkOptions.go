package gotsw

type MetalNetworkOptions struct {
	SpeedGbps    int32   `json:"speedGbps"`
	MonthlyPrice float64 `json:"monthlyPrice"`
	HourlyPrice  float64 `json:"hourlyPrice"`
	Default      bool    `json:"default"`
	IsBonded     bool    `json:"isBonded"`
}

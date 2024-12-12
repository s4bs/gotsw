package gotsw

type MetalMemoryOptions struct {
	Gb           int     `json:"gb"`
	MonthlyPrice float64 `json:"monthlyPrice"`
	HourlyPrice  float64 `json:"hourlyPrice"`
	Default      bool    `json:"default"`
}

package gotsw

type MetalDriveSlots struct {
	Id       string               `json:"id"`
	Default  string               `json:"default"`
	Required bool                 `json:"required"`
	Options  []MetalStorageDevice `json:"options"`
}

type MetalStorageDevice struct {
	Name         string                         `json:"name"`
	Default      bool                           `json:"default"`
	Type         int32                          `json:"type"`
	CapacityGb   int32                          `json:"capacityGb"`
	Details      MetalStorageDeviceDriveDetails `json:"details"`
	MonthlyPrice float64                        `json:"monthlyPrice"`
	HourlyPrice  float64                        `json:"hourlyPrice"`
	IsBossDrive  bool                           `json:"isBossDrive"`
	Id           string                         `json:"id"`
}

type MetalStorageDeviceDriveDetails struct {
	DeviceName string `json:"deviceName"`
	Serial     string `json:"serial"`
}

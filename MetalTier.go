package gotsw

type MetalTier struct {
	Id                 string `json:"id"`
	ExternalIdentifier string `json:"externalIdentifier"`
	Availability       map[string]struct {
		MaxQuantity int `json:"maxQuantity"`
	} `json:"availability"`
	Cpu                string                `json:"cpu"`
	CpuDescription     string                `json:"cpuDescription"`
	MemoryOptions      []MetalMemoryOptions  `json:"memoryOptions"`
	DriveSlots         []MetalDriveSlots     `json:"driveSlots"`
	NetworkOptions     []MetalNetworkOptions `json:"networkOptions"`
	MonthlyPrice       float64               `json:"monthlyPrice"`
	HourlyPrice        float64               `json:"hourlyPrice"`
	Hidden             bool                  `json:"hidden"`
	MemoryOptionSetId  int                   `json:"memoryOptionSetId"`
	DriveSlotSetId     int                   `json:"driveSlotSetId"`
	NetworkOptionSetId int                   `json:"networkOptionSetId"`
	TierType           int                   `json:"tierType"`
}

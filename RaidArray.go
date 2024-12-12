package gotsw

type RaidArray struct {
	Name    string   `json:"name"`
	Type    uint32   `json:"raidType"`
	Members []string `json:"members"`

	Partitions []Partition `json:"partitions"`

	SizeBytes  uint64 `json:"sizeBytes"`
	FileSystem uint32 `json:"fileSystem"`
	MountPoint string `json:"mountPoint"`
}

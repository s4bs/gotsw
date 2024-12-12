package gotsw

type Partition struct {
	Name       string `json:"name"`
	Device     string `json:"device"`
	SizeBytes  uint64 `json:"sizeBytes"`
	FileSystem uint32 `json:"fileSystem"`
	MountPoint string `json:"mountPoint"`
}

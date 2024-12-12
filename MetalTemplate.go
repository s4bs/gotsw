package gotsw

type MetalTemplate struct {
	Id          uint64             `json:"id"`
	ProjectId   uint64             `json:"projectId"`
	DisplayName string             `json:"displayName"`
	CreateModel MetalCreateRequest `json:"createModel"`
	CloudInit   string             `json:"cloudInit"`
	Deleted     string             `json:"deleted"`
	Created     string             `json:"created"`
}

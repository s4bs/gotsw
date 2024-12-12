package gotsw

import (
	"context"
	"fmt"
	"net/http"
)

const metalBasePath = "v2/Metal"

type MetalService interface {
	List(ctx context.Context, options *MetalListOptions) ([]Metal, error)
	Create(ctx context.Context, metalReq *MetalCreateRequest) (*Metal, error)
	Get(ctx context.Context, id uint64) (*Metal, error)
	Reinstall(ctx context.Context, id uint64, reinstallReq *MetalReinstallRequest) (*Metal, error)

	PowerOff(ctx context.Context, id uint64) (*Metal, error)
	PowerOn(ctx context.Context, id uint64) (*Metal, error)
	Reset(ctx context.Context, id uint64) (*Metal, error)

	ListTemplates(ctx context.Context) ([]MetalTemplate, error)
	ListTiers(ctx context.Context) ([]MetalTier, error)
}

type MetalListOptions struct {
	Limit         uint64 `url:"Limit,omitempty"`
	MetalTierType uint64 `url:"MetalTierType,omitempty"`
	ProjectId     uint64 `url:"ProjectId,omitempty"`
	RegionId      string `url:"Region,omitempty"`
	Skip          uint64 `url:"Skip,omitempty"`
	Status        string `url:"Status,omitempty"`
	Tag           string `url:"Tag,omitempty"`
	TierId        string `url:"Tier,omitempty"`
}

type MetalServiceHandler struct {
	client *Client
}

type metalsRoot struct {
	Metals []Metal `json:"result"`
}

type metalRoot struct {
	Metal *Metal `json:"result"`
}

type metalTemplatesRoot struct {
	MetalTemplates []MetalTemplate `json:"result"`
}

type metalTiersRoot struct {
	MetalTiers []MetalTier `json:"result"`
}

type PowerCommandRequest struct {
	Command string `json:"command"`
}

type MetalCreateRequest struct {
	Disks          map[string]string `json:"disks"`
	DisplayName    string            `json:"displayName"`
	ImageId        string            `json:"imageId"`
	IpxeUrl        string            `json:"ipxeUrl"`
	Partitions     []Partition       `json:"partitions"`
	Password       string            `json:"password"`
	Quantity       uint64            `json:"quantity"`
	RaidArrays     []RaidArray       `json:"raidArrays"`
	RegionId       string            `json:"regionId"`
	ReservePriving bool              `json:"reservePriving"`
	SshKeyIds      []uint64          `json:"sshKeyIds"`
	Tags           []string          `json:"tags"`
	TemplateId     uint64            `json:"templateId"`
	TierId         string            `json:"tierId"`
	UserData       string            `json:"userData"`
}

func (s *MetalServiceHandler) List(ctx context.Context, options *MetalListOptions) ([]Metal, error) {
	path := metalBasePath
	path, err := addOptions(path, options)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	root := new(metalsRoot)
	if err = s.client.Do(ctx, req, root); err != nil {
		return nil, err
	}

	return root.Metals, nil
}

func (s *MetalServiceHandler) PowerOff(ctx context.Context, id uint64) (*Metal, error) {
	path := fmt.Sprintf(metalBasePath+"/%d/PowerCommand", id)

	commandRequest := PowerCommandRequest{
		Command: "PowerOff",
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, commandRequest)
	if err != nil {
		return nil, err
	}

	root := new(metalRoot)
	if err = s.client.Do(ctx, req, root); err != nil {
		return nil, err
	}

	return root.Metal, err
}

func (s *MetalServiceHandler) PowerOn(ctx context.Context, id uint64) (*Metal, error) {
	path := fmt.Sprintf(metalBasePath+"/%d/PowerCommand", id)

	commandRequest := PowerCommandRequest{
		Command: "PowerOn",
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, commandRequest)
	if err != nil {
		return nil, err
	}

	root := new(metalRoot)
	if err = s.client.Do(ctx, req, root); err != nil {
		return nil, err
	}

	return root.Metal, err
}

func (s *MetalServiceHandler) Reset(ctx context.Context, id uint64) (*Metal, error) {
	path := fmt.Sprintf(metalBasePath+"/%d/PowerCommand", id)

	commandRequest := PowerCommandRequest{
		Command: "Reset",
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, commandRequest)
	if err != nil {
		return nil, err
	}

	root := new(metalRoot)
	if err = s.client.Do(ctx, req, root); err != nil {
		return nil, err
	}

	return root.Metal, err
}

func (s *MetalServiceHandler) Get(ctx context.Context, id uint64) (*Metal, error) {
	path := fmt.Sprintf(metalBasePath+"/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	root := new(metalRoot)
	if err = s.client.Do(ctx, req, root); err != nil {
		return nil, err
	}

	return root.Metal, nil
}

func (s *MetalServiceHandler) Create(ctx context.Context, createRequest *MetalCreateRequest) (*Metal, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, metalBasePath, createRequest)
	if err != nil {
		return nil, err
	}

	root := new(metalRoot)
	if err = s.client.Do(ctx, req, root); err != nil {
		return nil, err
	}

	return root.Metal, err
}

func (s *MetalServiceHandler) ListTemplates(ctx context.Context) ([]MetalTemplate, error) {
	path := metalBasePath + "/Templates"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	root := new(metalTemplatesRoot)
	if err = s.client.Do(ctx, req, root); err != nil {
		return nil, err
	}

	return root.MetalTemplates, nil
}

func (s *MetalServiceHandler) ListTiers(ctx context.Context) ([]MetalTier, error) {
	path := metalBasePath + "/Tiers"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	root := new(metalTiersRoot)
	if err = s.client.Do(ctx, req, root); err != nil {
		return nil, err
	}

	return root.MetalTiers, nil
}

type MetalReinstallRequest struct {
	Disks          map[string]string `json:"disks"`
	DisplayName    string            `json:"displayName"`
	ImageId        string            `json:"imageId"`
	IpxeUrl        string            `json:"ipxeUrl"`
	Partitions     []Partition       `json:"partitions"`
	Password       string            `json:"password"`
	Quantity       uint32            `json:"quantity"`
	RaidArrays     []RaidArray       `json:"raidArrays"`
	RegionId       string            `json:"regionId"`
	ReservePriving bool              `json:"reservePriving"`
	SshKeyIds      []uint64          `json:"sshKeyIds"`
	SshKeys        []struct {
		Id          int64  `json:"id"`
		Created     string `json:"created"`
		Deleted     string `json:"deleted"`
		ProjectId   int64  `json:"projectId"`
		DisplayName string `json:"displayName"`
		Key         string `json:"key"`
	} `json:"sshKeys"`
	Tags       []string  `json:"tags"`
	TemplateId uint64    `json:"templateId"`
	TierId     string    `json:"tierId"`
	UserData   string    `json:"userData"`
	ProjectId  uint64    `json:"projectId"` // new
	MemoryGb   uint32    `json:"memoryGb"`  // new
	TierObj    MetalTier `json:"tierObj"`   // new
	Service    Metal     `json:"service"`   // new
}

func (s *MetalServiceHandler) Reinstall(ctx context.Context, id uint64, reinstallRequest *MetalReinstallRequest) (*Metal, error) {
	path := fmt.Sprintf(metalBasePath+"/%d/Reinstall", id)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, reinstallRequest)
	if err != nil {
		return nil, err
	}

	root := new(metalRoot)
	if err = s.client.Do(ctx, req, root); err != nil {
		return nil, err
	}

	return root.Metal, nil
}

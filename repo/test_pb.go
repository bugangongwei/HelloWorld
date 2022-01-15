package repo

type GetHistoriesRequest struct {
	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

type GetHistoriesResponse struct{}

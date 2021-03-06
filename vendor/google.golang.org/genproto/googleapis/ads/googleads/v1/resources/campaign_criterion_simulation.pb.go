// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_criterion_simulation.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A campaign criterion simulation. Supported combinations of advertising
// channel type, criterion ids, simulation type and simulation modification
// method is detailed below respectively.
//
// SEARCH   30000,30001,30002  BID_MODIFIER  UNIFORM
// DISPLAY  30001              BID_MODIFIER  UNIFORM
type CampaignCriterionSimulation struct {
	// The resource name of the campaign criterion simulation.
	// Campaign criterion simulation resource names have the form:
	//
	//
	// `customers/{customer_id}/campaignCriterionSimulations/{campaign_id}~{criterion_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Campaign ID of the simulation.
	CampaignId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	// Criterion ID of the simulation.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,5,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v1.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// First day on which the simulation is based, in YYYY-MM-DD format.
	StartDate *wrappers.StringValue `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Last day on which the simulation is based, in YYYY-MM-DD format.
	EndDate *wrappers.StringValue `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are valid to be assigned to PointList:
	//	*CampaignCriterionSimulation_BidModifierPointList
	PointList            isCampaignCriterionSimulation_PointList `protobuf_oneof:"point_list"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CampaignCriterionSimulation) Reset()         { *m = CampaignCriterionSimulation{} }
func (m *CampaignCriterionSimulation) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterionSimulation) ProtoMessage()    {}
func (*CampaignCriterionSimulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eae0e5c50ed969c, []int{0}
}

func (m *CampaignCriterionSimulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterionSimulation.Unmarshal(m, b)
}
func (m *CampaignCriterionSimulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterionSimulation.Marshal(b, m, deterministic)
}
func (m *CampaignCriterionSimulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterionSimulation.Merge(m, src)
}
func (m *CampaignCriterionSimulation) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterionSimulation.Size(m)
}
func (m *CampaignCriterionSimulation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterionSimulation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterionSimulation proto.InternalMessageInfo

func (m *CampaignCriterionSimulation) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignCriterionSimulation) GetCampaignId() *wrappers.Int64Value {
	if m != nil {
		return m.CampaignId
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if m != nil {
		return m.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (m *CampaignCriterionSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if m != nil {
		return m.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (m *CampaignCriterionSimulation) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type isCampaignCriterionSimulation_PointList interface {
	isCampaignCriterionSimulation_PointList()
}

type CampaignCriterionSimulation_BidModifierPointList struct {
	BidModifierPointList *common.BidModifierSimulationPointList `protobuf:"bytes,8,opt,name=bid_modifier_point_list,json=bidModifierPointList,proto3,oneof"`
}

func (*CampaignCriterionSimulation_BidModifierPointList) isCampaignCriterionSimulation_PointList() {}

func (m *CampaignCriterionSimulation) GetPointList() isCampaignCriterionSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetBidModifierPointList() *common.BidModifierSimulationPointList {
	if x, ok := m.GetPointList().(*CampaignCriterionSimulation_BidModifierPointList); ok {
		return x.BidModifierPointList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterionSimulation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterionSimulation_BidModifierPointList)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignCriterionSimulation)(nil), "google.ads.googleads.v1.resources.CampaignCriterionSimulation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_criterion_simulation.proto", fileDescriptor_4eae0e5c50ed969c)
}

var fileDescriptor_4eae0e5c50ed969c = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0xe7, 0xf4, 0xbf, 0x9a, 0xed, 0xc2, 0x1b, 0xcc, 0xb4, 0x65, 0xa4, 0x1b, 0x85, 0x5c,
	0xc9, 0xa4, 0x1d, 0x1b, 0xb8, 0xa3, 0x2c, 0xe9, 0x4a, 0x97, 0xb1, 0x8c, 0xe0, 0x96, 0x5c, 0x8c,
	0x80, 0x51, 0x2c, 0xd5, 0x13, 0x44, 0x92, 0x91, 0xe4, 0x96, 0x3e, 0x44, 0x5f, 0x62, 0xbb, 0xdb,
	0xa3, 0xec, 0x51, 0xf6, 0x14, 0x23, 0xb2, 0xac, 0x94, 0x64, 0x69, 0x73, 0x77, 0x7c, 0xf4, 0x7d,
	0xbf, 0x63, 0x7d, 0xb2, 0x0c, 0xce, 0x32, 0x21, 0xb2, 0x31, 0x09, 0x11, 0x56, 0x61, 0x59, 0x4e,
	0xaa, 0xeb, 0x56, 0x28, 0x89, 0x12, 0x85, 0x4c, 0x89, 0x0a, 0x53, 0xc4, 0x72, 0x44, 0x33, 0x9e,
	0xa4, 0x92, 0x6a, 0x22, 0xa9, 0xe0, 0x89, 0xa2, 0xac, 0x18, 0x23, 0x4d, 0x05, 0x87, 0xb9, 0x14,
	0x5a, 0xf8, 0xfb, 0xa5, 0x17, 0x22, 0xac, 0xa0, 0xc3, 0xc0, 0xeb, 0x16, 0x74, 0x98, 0x9d, 0x70,
	0xd1, 0xa4, 0x54, 0x30, 0x26, 0x78, 0x38, 0xcb, 0xdc, 0xe9, 0x2c, 0x32, 0x10, 0x5e, 0x30, 0x75,
	0x4f, 0x9f, 0x30, 0x81, 0xe9, 0x15, 0x4d, 0xed, 0x03, 0xd1, 0x3f, 0x04, 0xb6, 0x8c, 0xa3, 0xa5,
	0x19, 0xfa, 0x36, 0x27, 0xd6, 0xf4, 0xca, 0x9a, 0xcc, 0xd3, 0xa8, 0xb8, 0x0a, 0x6f, 0x24, 0xca,
	0x73, 0x22, 0x95, 0x5d, 0xdf, 0xab, 0xa0, 0x39, 0x0d, 0x11, 0xe7, 0x42, 0x1b, 0x82, 0x5d, 0x7d,
	0xfd, 0x6b, 0x0d, 0xec, 0x9e, 0xda, 0xc8, 0x4e, 0xab, 0xc4, 0x2e, 0xdc, 0x20, 0xff, 0x0d, 0x78,
	0x5a, 0x85, 0x92, 0x70, 0xc4, 0x48, 0xe0, 0x35, 0xbc, 0xe6, 0x56, 0x5c, 0xaf, 0x9a, 0xdf, 0x10,
	0x23, 0xfe, 0x07, 0xb0, 0xed, 0x62, 0xa7, 0x38, 0xa8, 0x35, 0xbc, 0xe6, 0xf6, 0xe1, 0xae, 0x8d,
	0x16, 0x56, 0x2f, 0x06, 0xbb, 0x5c, 0xbf, 0x7b, 0x3b, 0x40, 0xe3, 0x82, 0xc4, 0xa0, 0xd2, 0x77,
	0xb1, 0x7f, 0x02, 0xea, 0xd3, 0xb3, 0xa2, 0x38, 0x58, 0x79, 0xdc, 0xbe, 0xed, 0x0c, 0x5d, 0xec,
	0x5f, 0x82, 0xd5, 0x49, 0x1c, 0xc1, 0x6a, 0xc3, 0x6b, 0x3e, 0x3b, 0xfc, 0x08, 0x17, 0x1d, 0xae,
	0x09, 0x11, 0x4e, 0xf7, 0x76, 0x79, 0x9b, 0x93, 0x33, 0x5e, 0xb0, 0x99, 0x56, 0x6c, 0x68, 0xfe,
	0x9d, 0x07, 0x9e, 0xff, 0xe7, 0xa4, 0x82, 0x35, 0x33, 0x65, 0xb8, 0xf4, 0x94, 0xde, 0x3d, 0x46,
	0xcf, 0x20, 0x66, 0x66, 0xce, 0x0b, 0x62, 0x9f, 0xcd, 0xf5, 0xfc, 0x63, 0x00, 0x94, 0x46, 0x52,
	0x27, 0x18, 0x69, 0x12, 0xac, 0x9b, 0x8c, 0xf6, 0xe6, 0x32, 0xba, 0xd0, 0x92, 0xf2, 0xac, 0x0c,
	0x69, 0xcb, 0xe8, 0x3f, 0x21, 0x4d, 0xfc, 0xf7, 0x60, 0x93, 0x70, 0x5c, 0x5a, 0x37, 0x96, 0xb0,
	0x6e, 0x10, 0x8e, 0x8d, 0xf1, 0x06, 0xbc, 0x1c, 0x51, 0x6c, 0x3f, 0x59, 0x22, 0x93, 0x5c, 0x50,
	0xae, 0x93, 0x31, 0x55, 0x3a, 0xd8, 0x34, 0x9c, 0x93, 0x85, 0x41, 0x94, 0x17, 0x05, 0x76, 0x28,
	0xee, 0x59, 0xf7, 0x74, 0xcf, 0xfd, 0x09, 0xe6, 0x2b, 0x55, 0xfa, 0xf3, 0x93, 0xf8, 0xc5, 0x68,
	0xaa, 0x70, 0xfd, 0x4e, 0x1d, 0x80, 0xe9, 0xac, 0xce, 0x5d, 0x0d, 0x1c, 0xa4, 0x82, 0xc1, 0x47,
	0xef, 0x6d, 0xa7, 0xf1, 0xc0, 0xc7, 0xdc, 0x9f, 0x6c, 0xb5, 0xef, 0x7d, 0xff, 0x62, 0x31, 0x99,
	0x18, 0x23, 0x9e, 0x41, 0x21, 0xb3, 0x30, 0x23, 0xdc, 0x04, 0x51, 0x5d, 0xbb, 0x9c, 0xaa, 0x07,
	0x7e, 0x32, 0xc7, 0xae, 0xfa, 0x59, 0x5b, 0x39, 0x6f, 0xb7, 0x7f, 0xd7, 0xf6, 0xcf, 0x4b, 0x64,
	0x1b, 0x2b, 0x58, 0x96, 0x93, 0x6a, 0xd0, 0x82, 0x71, 0xa5, 0xfc, 0x53, 0x69, 0x86, 0x6d, 0xac,
	0x86, 0x4e, 0x33, 0x1c, 0xb4, 0x86, 0x4e, 0xf3, 0xb7, 0x76, 0x50, 0x2e, 0x44, 0x51, 0x1b, 0xab,
	0x28, 0x72, 0xaa, 0x28, 0x1a, 0xb4, 0xa2, 0xc8, 0xe9, 0x46, 0xeb, 0xe6, 0x65, 0x8f, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xb8, 0x24, 0xd5, 0x5b, 0x10, 0x05, 0x00, 0x00,
}

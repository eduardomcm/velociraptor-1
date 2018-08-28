// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flows.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import proto2 "www.velocidex.com/golang/velociraptor/actions/proto"
import proto1 "www.velocidex.com/golang/velociraptor/crypto/proto"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FlowContext_State int32

const (
	FlowContext_UNSET      FlowContext_State = 0
	FlowContext_RUNNING    FlowContext_State = 1
	FlowContext_TERMINATED FlowContext_State = 2
	FlowContext_ERROR      FlowContext_State = 3
)

var FlowContext_State_name = map[int32]string{
	0: "UNSET",
	1: "RUNNING",
	2: "TERMINATED",
	3: "ERROR",
}
var FlowContext_State_value = map[string]int32{
	"UNSET":      0,
	"RUNNING":    1,
	"TERMINATED": 2,
	"ERROR":      3,
}

func (x FlowContext_State) String() string {
	return proto.EnumName(FlowContext_State_name, int32(x))
}
func (FlowContext_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_flows_bb384fe8cd32ec3a, []int{1, 0}
}

type StartFlowRequest struct {
	RunnerArgs           *FlowRunnerArgs `protobuf:"bytes,1,opt,name=runner_args,json=runnerArgs,proto3" json:"runner_args,omitempty"`
	Args                 *any.Any        `protobuf:"bytes,3,opt,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StartFlowRequest) Reset()         { *m = StartFlowRequest{} }
func (m *StartFlowRequest) String() string { return proto.CompactTextString(m) }
func (*StartFlowRequest) ProtoMessage()    {}
func (*StartFlowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_bb384fe8cd32ec3a, []int{0}
}
func (m *StartFlowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartFlowRequest.Unmarshal(m, b)
}
func (m *StartFlowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartFlowRequest.Marshal(b, m, deterministic)
}
func (dst *StartFlowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartFlowRequest.Merge(dst, src)
}
func (m *StartFlowRequest) XXX_Size() int {
	return xxx_messageInfo_StartFlowRequest.Size(m)
}
func (m *StartFlowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartFlowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartFlowRequest proto.InternalMessageInfo

func (m *StartFlowRequest) GetRunnerArgs() *FlowRunnerArgs {
	if m != nil {
		return m.RunnerArgs
	}
	return nil
}

func (m *StartFlowRequest) GetArgs() *any.Any {
	if m != nil {
		return m.Args
	}
	return nil
}

// The flow context.
// Next field: 19
type FlowContext struct {
	Backtrace            string                  `protobuf:"bytes,1,opt,name=backtrace,proto3" json:"backtrace,omitempty"`
	ClientResources      *proto1.ClientResources `protobuf:"bytes,2,opt,name=client_resources,json=clientResources,proto3" json:"client_resources,omitempty"`
	CreateTime           uint64                  `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CurrentState         string                  `protobuf:"bytes,5,opt,name=current_state,json=currentState,proto3" json:"current_state,omitempty"`
	KillTimestamp        uint64                  `protobuf:"varint,6,opt,name=kill_timestamp,json=killTimestamp,proto3" json:"kill_timestamp,omitempty"`
	NetworkBytesSent     uint64                  `protobuf:"varint,7,opt,name=network_bytes_sent,json=networkBytesSent,proto3" json:"network_bytes_sent,omitempty"`
	NextOutboundId       uint64                  `protobuf:"varint,8,opt,name=next_outbound_id,json=nextOutboundId,proto3" json:"next_outbound_id,omitempty"`
	NextProcessedRequest uint64                  `protobuf:"varint,9,opt,name=next_processed_request,json=nextProcessedRequest,proto3" json:"next_processed_request,omitempty"`
	OutstandingRequests  uint64                  `protobuf:"varint,11,opt,name=outstanding_requests,json=outstandingRequests,proto3" json:"outstanding_requests,omitempty"`
	TotalResults         uint64                  `protobuf:"varint,18,opt,name=total_results,json=totalResults,proto3" json:"total_results,omitempty"`
	UploadedFiles        []string                `protobuf:"bytes,19,rep,name=uploaded_files,json=uploadedFiles,proto3" json:"uploaded_files,omitempty"`
	Logs                 []*proto1.LogMessage    `protobuf:"bytes,20,rep,name=logs,proto3" json:"logs,omitempty"`
	SessionId            string                  `protobuf:"bytes,13,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	State                FlowContext_State       `protobuf:"varint,14,opt,name=state,proto3,enum=proto.FlowContext_State" json:"state,omitempty"`
	Status               string                  `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	UserNotified         bool                    `protobuf:"varint,16,opt,name=user_notified,json=userNotified,proto3" json:"user_notified,omitempty"`
	ActiveTime           uint64                  `protobuf:"varint,17,opt,name=active_time,json=activeTime,proto3" json:"active_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FlowContext) Reset()         { *m = FlowContext{} }
func (m *FlowContext) String() string { return proto.CompactTextString(m) }
func (*FlowContext) ProtoMessage()    {}
func (*FlowContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_bb384fe8cd32ec3a, []int{1}
}
func (m *FlowContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowContext.Unmarshal(m, b)
}
func (m *FlowContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowContext.Marshal(b, m, deterministic)
}
func (dst *FlowContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowContext.Merge(dst, src)
}
func (m *FlowContext) XXX_Size() int {
	return xxx_messageInfo_FlowContext.Size(m)
}
func (m *FlowContext) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowContext.DiscardUnknown(m)
}

var xxx_messageInfo_FlowContext proto.InternalMessageInfo

func (m *FlowContext) GetBacktrace() string {
	if m != nil {
		return m.Backtrace
	}
	return ""
}

func (m *FlowContext) GetClientResources() *proto1.ClientResources {
	if m != nil {
		return m.ClientResources
	}
	return nil
}

func (m *FlowContext) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *FlowContext) GetCurrentState() string {
	if m != nil {
		return m.CurrentState
	}
	return ""
}

func (m *FlowContext) GetKillTimestamp() uint64 {
	if m != nil {
		return m.KillTimestamp
	}
	return 0
}

func (m *FlowContext) GetNetworkBytesSent() uint64 {
	if m != nil {
		return m.NetworkBytesSent
	}
	return 0
}

func (m *FlowContext) GetNextOutboundId() uint64 {
	if m != nil {
		return m.NextOutboundId
	}
	return 0
}

func (m *FlowContext) GetNextProcessedRequest() uint64 {
	if m != nil {
		return m.NextProcessedRequest
	}
	return 0
}

func (m *FlowContext) GetOutstandingRequests() uint64 {
	if m != nil {
		return m.OutstandingRequests
	}
	return 0
}

func (m *FlowContext) GetTotalResults() uint64 {
	if m != nil {
		return m.TotalResults
	}
	return 0
}

func (m *FlowContext) GetUploadedFiles() []string {
	if m != nil {
		return m.UploadedFiles
	}
	return nil
}

func (m *FlowContext) GetLogs() []*proto1.LogMessage {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *FlowContext) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *FlowContext) GetState() FlowContext_State {
	if m != nil {
		return m.State
	}
	return FlowContext_UNSET
}

func (m *FlowContext) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FlowContext) GetUserNotified() bool {
	if m != nil {
		return m.UserNotified
	}
	return false
}

func (m *FlowContext) GetActiveTime() uint64 {
	if m != nil {
		return m.ActiveTime
	}
	return 0
}

// Next field: 23
type FlowRunnerArgs struct {
	Creator              string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	NotifyToUser         bool     `protobuf:"varint,2,opt,name=notify_to_user,json=notifyToUser,proto3" json:"notify_to_user,omitempty"`
	ClientId             string   `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CpuLimit             uint64   `protobuf:"varint,9,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	NetworkBytesLimit    uint64   `protobuf:"varint,13,opt,name=network_bytes_limit,json=networkBytesLimit,proto3" json:"network_bytes_limit,omitempty"`
	FlowName             string   `protobuf:"bytes,11,opt,name=flow_name,json=flowName,proto3" json:"flow_name,omitempty"`
	StartTime            uint64   `protobuf:"varint,15,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Args                 *any.Any `protobuf:"bytes,24,opt,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowRunnerArgs) Reset()         { *m = FlowRunnerArgs{} }
func (m *FlowRunnerArgs) String() string { return proto.CompactTextString(m) }
func (*FlowRunnerArgs) ProtoMessage()    {}
func (*FlowRunnerArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_bb384fe8cd32ec3a, []int{2}
}
func (m *FlowRunnerArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRunnerArgs.Unmarshal(m, b)
}
func (m *FlowRunnerArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRunnerArgs.Marshal(b, m, deterministic)
}
func (dst *FlowRunnerArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRunnerArgs.Merge(dst, src)
}
func (m *FlowRunnerArgs) XXX_Size() int {
	return xxx_messageInfo_FlowRunnerArgs.Size(m)
}
func (m *FlowRunnerArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRunnerArgs.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRunnerArgs proto.InternalMessageInfo

func (m *FlowRunnerArgs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *FlowRunnerArgs) GetNotifyToUser() bool {
	if m != nil {
		return m.NotifyToUser
	}
	return false
}

func (m *FlowRunnerArgs) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *FlowRunnerArgs) GetCpuLimit() uint64 {
	if m != nil {
		return m.CpuLimit
	}
	return 0
}

func (m *FlowRunnerArgs) GetNetworkBytesLimit() uint64 {
	if m != nil {
		return m.NetworkBytesLimit
	}
	return 0
}

func (m *FlowRunnerArgs) GetFlowName() string {
	if m != nil {
		return m.FlowName
	}
	return ""
}

func (m *FlowRunnerArgs) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *FlowRunnerArgs) GetArgs() *any.Any {
	if m != nil {
		return m.Args
	}
	return nil
}

// This is a short lived protobuf to hold various flow state
// information only valid between the start and end of the flow. It is
// used to keep state over multiple client round trips.
type VelociraptorFlowState struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Payload              *any.Any `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VelociraptorFlowState) Reset()         { *m = VelociraptorFlowState{} }
func (m *VelociraptorFlowState) String() string { return proto.CompactTextString(m) }
func (*VelociraptorFlowState) ProtoMessage()    {}
func (*VelociraptorFlowState) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_bb384fe8cd32ec3a, []int{3}
}
func (m *VelociraptorFlowState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VelociraptorFlowState.Unmarshal(m, b)
}
func (m *VelociraptorFlowState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VelociraptorFlowState.Marshal(b, m, deterministic)
}
func (dst *VelociraptorFlowState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VelociraptorFlowState.Merge(dst, src)
}
func (m *VelociraptorFlowState) XXX_Size() int {
	return xxx_messageInfo_VelociraptorFlowState.Size(m)
}
func (m *VelociraptorFlowState) XXX_DiscardUnknown() {
	xxx_messageInfo_VelociraptorFlowState.DiscardUnknown(m)
}

var xxx_messageInfo_VelociraptorFlowState proto.InternalMessageInfo

func (m *VelociraptorFlowState) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VelociraptorFlowState) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Used to describe a flow. The UI uses this information to work out
// how to render the flow and create its form.
type FlowDescriptor struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FriendlyName         string   `protobuf:"bytes,2,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	Category             string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Doc                  string   `protobuf:"bytes,4,opt,name=doc,proto3" json:"doc,omitempty"`
	ArgsType             string   `protobuf:"bytes,5,opt,name=args_type,json=argsType,proto3" json:"args_type,omitempty"`
	DefaultArgs          *any.Any `protobuf:"bytes,6,opt,name=default_args,json=defaultArgs,proto3" json:"default_args,omitempty"`
	Behaviours           string   `protobuf:"bytes,7,opt,name=behaviours,proto3" json:"behaviours,omitempty"`
	Internal             bool     `protobuf:"varint,8,opt,name=internal,proto3" json:"internal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowDescriptor) Reset()         { *m = FlowDescriptor{} }
func (m *FlowDescriptor) String() string { return proto.CompactTextString(m) }
func (*FlowDescriptor) ProtoMessage()    {}
func (*FlowDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_bb384fe8cd32ec3a, []int{4}
}
func (m *FlowDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowDescriptor.Unmarshal(m, b)
}
func (m *FlowDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowDescriptor.Marshal(b, m, deterministic)
}
func (dst *FlowDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowDescriptor.Merge(dst, src)
}
func (m *FlowDescriptor) XXX_Size() int {
	return xxx_messageInfo_FlowDescriptor.Size(m)
}
func (m *FlowDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_FlowDescriptor proto.InternalMessageInfo

func (m *FlowDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FlowDescriptor) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *FlowDescriptor) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *FlowDescriptor) GetDoc() string {
	if m != nil {
		return m.Doc
	}
	return ""
}

func (m *FlowDescriptor) GetArgsType() string {
	if m != nil {
		return m.ArgsType
	}
	return ""
}

func (m *FlowDescriptor) GetDefaultArgs() *any.Any {
	if m != nil {
		return m.DefaultArgs
	}
	return nil
}

func (m *FlowDescriptor) GetBehaviours() string {
	if m != nil {
		return m.Behaviours
	}
	return ""
}

func (m *FlowDescriptor) GetInternal() bool {
	if m != nil {
		return m.Internal
	}
	return false
}

// Interrogate flow discovers information about the client.
type VInterrogateArgs struct {
	// If set a light weight version of the flow is run.
	Lightweight bool `protobuf:"varint,1,opt,name=lightweight,proto3" json:"lightweight,omitempty"`
	// Additional VQL queries to run.
	Queries              []*proto2.VQLRequest `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VInterrogateArgs) Reset()         { *m = VInterrogateArgs{} }
func (m *VInterrogateArgs) String() string { return proto.CompactTextString(m) }
func (*VInterrogateArgs) ProtoMessage()    {}
func (*VInterrogateArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_bb384fe8cd32ec3a, []int{5}
}
func (m *VInterrogateArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VInterrogateArgs.Unmarshal(m, b)
}
func (m *VInterrogateArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VInterrogateArgs.Marshal(b, m, deterministic)
}
func (dst *VInterrogateArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VInterrogateArgs.Merge(dst, src)
}
func (m *VInterrogateArgs) XXX_Size() int {
	return xxx_messageInfo_VInterrogateArgs.Size(m)
}
func (m *VInterrogateArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_VInterrogateArgs.DiscardUnknown(m)
}

var xxx_messageInfo_VInterrogateArgs proto.InternalMessageInfo

func (m *VInterrogateArgs) GetLightweight() bool {
	if m != nil {
		return m.Lightweight
	}
	return false
}

func (m *VInterrogateArgs) GetQueries() []*proto2.VQLRequest {
	if m != nil {
		return m.Queries
	}
	return nil
}

type AFF4FlowObject struct {
	Urn                  string          `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	RunnerArgs           *FlowRunnerArgs `protobuf:"bytes,2,opt,name=runner_args,json=runnerArgs,proto3" json:"runner_args,omitempty"`
	FlowContext          *FlowContext    `protobuf:"bytes,3,opt,name=flow_context,json=flowContext,proto3" json:"flow_context,omitempty"`
	FlowState            *any.Any        `protobuf:"bytes,4,opt,name=flow_state,json=flowState,proto3" json:"flow_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AFF4FlowObject) Reset()         { *m = AFF4FlowObject{} }
func (m *AFF4FlowObject) String() string { return proto.CompactTextString(m) }
func (*AFF4FlowObject) ProtoMessage()    {}
func (*AFF4FlowObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_bb384fe8cd32ec3a, []int{6}
}
func (m *AFF4FlowObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AFF4FlowObject.Unmarshal(m, b)
}
func (m *AFF4FlowObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AFF4FlowObject.Marshal(b, m, deterministic)
}
func (dst *AFF4FlowObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AFF4FlowObject.Merge(dst, src)
}
func (m *AFF4FlowObject) XXX_Size() int {
	return xxx_messageInfo_AFF4FlowObject.Size(m)
}
func (m *AFF4FlowObject) XXX_DiscardUnknown() {
	xxx_messageInfo_AFF4FlowObject.DiscardUnknown(m)
}

var xxx_messageInfo_AFF4FlowObject proto.InternalMessageInfo

func (m *AFF4FlowObject) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *AFF4FlowObject) GetRunnerArgs() *FlowRunnerArgs {
	if m != nil {
		return m.RunnerArgs
	}
	return nil
}

func (m *AFF4FlowObject) GetFlowContext() *FlowContext {
	if m != nil {
		return m.FlowContext
	}
	return nil
}

func (m *AFF4FlowObject) GetFlowState() *any.Any {
	if m != nil {
		return m.FlowState
	}
	return nil
}

type VQLNameTags struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VQLNameTags) Reset()         { *m = VQLNameTags{} }
func (m *VQLNameTags) String() string { return proto.CompactTextString(m) }
func (*VQLNameTags) ProtoMessage()    {}
func (*VQLNameTags) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_bb384fe8cd32ec3a, []int{7}
}
func (m *VQLNameTags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLNameTags.Unmarshal(m, b)
}
func (m *VQLNameTags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLNameTags.Marshal(b, m, deterministic)
}
func (dst *VQLNameTags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLNameTags.Merge(dst, src)
}
func (m *VQLNameTags) XXX_Size() int {
	return xxx_messageInfo_VQLNameTags.Size(m)
}
func (m *VQLNameTags) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLNameTags.DiscardUnknown(m)
}

var xxx_messageInfo_VQLNameTags proto.InternalMessageInfo

func (m *VQLNameTags) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VQLNameTags) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type VQLNameTagsState struct {
	Tags                 []*VQLNameTags `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *VQLNameTagsState) Reset()         { *m = VQLNameTagsState{} }
func (m *VQLNameTagsState) String() string { return proto.CompactTextString(m) }
func (*VQLNameTagsState) ProtoMessage()    {}
func (*VQLNameTagsState) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_bb384fe8cd32ec3a, []int{8}
}
func (m *VQLNameTagsState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLNameTagsState.Unmarshal(m, b)
}
func (m *VQLNameTagsState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLNameTagsState.Marshal(b, m, deterministic)
}
func (dst *VQLNameTagsState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLNameTagsState.Merge(dst, src)
}
func (m *VQLNameTagsState) XXX_Size() int {
	return xxx_messageInfo_VQLNameTagsState.Size(m)
}
func (m *VQLNameTagsState) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLNameTagsState.DiscardUnknown(m)
}

var xxx_messageInfo_VQLNameTagsState proto.InternalMessageInfo

func (m *VQLNameTagsState) GetTags() []*VQLNameTags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*StartFlowRequest)(nil), "proto.StartFlowRequest")
	proto.RegisterType((*FlowContext)(nil), "proto.FlowContext")
	proto.RegisterType((*FlowRunnerArgs)(nil), "proto.FlowRunnerArgs")
	proto.RegisterType((*VelociraptorFlowState)(nil), "proto.VelociraptorFlowState")
	proto.RegisterType((*FlowDescriptor)(nil), "proto.FlowDescriptor")
	proto.RegisterType((*VInterrogateArgs)(nil), "proto.VInterrogateArgs")
	proto.RegisterType((*AFF4FlowObject)(nil), "proto.AFF4FlowObject")
	proto.RegisterType((*VQLNameTags)(nil), "proto.VQLNameTags")
	proto.RegisterType((*VQLNameTagsState)(nil), "proto.VQLNameTagsState")
	proto.RegisterEnum("proto.FlowContext_State", FlowContext_State_name, FlowContext_State_value)
}

func init() { proto.RegisterFile("flows.proto", fileDescriptor_flows_bb384fe8cd32ec3a) }

var fileDescriptor_flows_bb384fe8cd32ec3a = []byte{
	// 1809 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xdd, 0x72, 0x1b, 0xb7,
	0xf5, 0xff, 0x53, 0xa2, 0x6d, 0x12, 0x94, 0x68, 0x1a, 0x92, 0xfd, 0xa7, 0x95, 0xcc, 0x04, 0x61,
	0x3b, 0xb5, 0x9a, 0x71, 0x57, 0xb6, 0xac, 0x24, 0x1e, 0xa5, 0xb1, 0x2b, 0xea, 0xc3, 0xd5, 0x8c,
	0x2c, 0xd5, 0x10, 0xa5, 0x76, 0xc6, 0xe9, 0x70, 0xc0, 0x5d, 0x90, 0x42, 0xbc, 0x04, 0x68, 0x00,
	0x2b, 0x99, 0xd3, 0xf6, 0x3a, 0xcf, 0xd0, 0x8b, 0x3e, 0x40, 0x5f, 0xa1, 0xd7, 0x9d, 0xf6, 0x3d,
	0xda, 0xbb, 0x3e, 0x43, 0x2f, 0x3a, 0xe7, 0x00, 0x4b, 0xd3, 0x69, 0xe3, 0x49, 0x7c, 0x61, 0xad,
	0x70, 0x7e, 0xe7, 0x03, 0xe7, 0xe3, 0x87, 0x23, 0xd2, 0x18, 0xe6, 0xe6, 0xca, 0x25, 0x13, 0x6b,
	0xbc, 0xa1, 0xd7, 0xf0, 0xc7, 0xda, 0xdd, 0x91, 0x31, 0xa3, 0x5c, 0x6e, 0xe0, 0x6f, 0x83, 0x62,
	0xb8, 0x21, 0xf4, 0x34, 0x20, 0xd6, 0xb6, 0xaf, 0xae, 0xae, 0x92, 0x4b, 0x99, 0x9b, 0x54, 0x65,
	0xf2, 0x4d, 0x92, 0x9a, 0xf1, 0xc6, 0xc8, 0xe4, 0x42, 0x8f, 0x36, 0xc2, 0xa1, 0x15, 0x13, 0x6f,
	0x6c, 0xd0, 0xdc, 0x70, 0x72, 0x2c, 0xb4, 0x57, 0x69, 0xd4, 0x7d, 0xfa, 0x43, 0x74, 0x21, 0xac,
	0xfe, 0x58, 0x7a, 0x91, 0x09, 0x2f, 0xa2, 0x81, 0x2f, 0xbf, 0x9f, 0x01, 0x91, 0x7a, 0x65, 0xb4,
	0x8b, 0x86, 0x2e, 0x5f, 0xe7, 0x3f, 0x4c, 0x3d, 0xb5, 0xd3, 0x89, 0x37, 0x51, 0xfb, 0x6b, 0x33,
	0x88, 0xc9, 0xe9, 0x7c, 0x53, 0x21, 0xad, 0x53, 0x2f, 0xac, 0x3f, 0xc8, 0xcd, 0x15, 0x97, 0xaf,
	0x0b, 0xe9, 0x3c, 0xfd, 0x8c, 0x34, 0x6c, 0xa1, 0xb5, 0xb4, 0x7d, 0x61, 0x47, 0xae, 0x5d, 0x61,
	0x95, 0xf5, 0xc6, 0xe6, 0xed, 0xa0, 0x91, 0x20, 0x10, 0xa5, 0x3b, 0x76, 0xe4, 0x38, 0xb1, 0xb3,
	0x6f, 0xfa, 0x05, 0xa9, 0xa2, 0xc2, 0x22, 0x2a, 0xac, 0x26, 0x21, 0xe3, 0x49, 0x99, 0xf1, 0x64,
	0x47, 0x4f, 0xbb, 0xb7, 0xfe, 0xf1, 0xef, 0x7f, 0xfe, 0xb5, 0xd2, 0xa0, 0x75, 0xb0, 0xc3, 0xd0,
	0x04, 0x2a, 0x75, 0xfe, 0x58, 0x23, 0x0d, 0x38, 0xdb, 0x35, 0xda, 0xcb, 0x37, 0x9e, 0x7e, 0x48,
	0xea, 0x03, 0x91, 0xbe, 0xf2, 0x56, 0xa4, 0x12, 0x43, 0xa8, 0xf3, 0xb7, 0x07, 0x74, 0x87, 0xb4,
	0xd2, 0x5c, 0x49, 0xed, 0xfb, 0x56, 0x3a, 0x53, 0xd8, 0x54, 0xba, 0xf6, 0x02, 0xba, 0xbd, 0x13,
	0xe3, 0xdc, 0x45, 0x31, 0x2f, 0xa5, 0xfc, 0x66, 0xfa, 0xee, 0x01, 0xdd, 0x22, 0x8d, 0xd4, 0x4a,
	0xe1, 0x65, 0xdf, 0xab, 0xb1, 0xc4, 0xa0, 0xab, 0xdd, 0x15, 0x0c, 0x6f, 0x99, 0x34, 0xf8, 0xde,
	0xc1, 0x9e, 0xf0, 0x12, 0x44, 0x9c, 0x04, 0x5c, 0x4f, 0x8d, 0x25, 0xfd, 0x11, 0x59, 0x4e, 0x0b,
	0x6b, 0xc1, 0xb3, 0xf3, 0xc2, 0xcb, 0xf6, 0x35, 0x0c, 0x6d, 0x29, 0x1e, 0x9e, 0xc2, 0x19, 0xdd,
	0x26, 0xcd, 0x57, 0x2a, 0xcf, 0xd1, 0xb0, 0xf3, 0x62, 0x3c, 0x69, 0x5f, 0xff, 0x6e, 0xeb, 0xcb,
	0x00, 0xed, 0x95, 0x48, 0x7a, 0x9f, 0x50, 0x2d, 0xfd, 0x95, 0xb1, 0xaf, 0xfa, 0x83, 0xa9, 0x97,
	0xae, 0xef, 0xa4, 0xf6, 0xed, 0x1b, 0xa0, 0xcf, 0x5b, 0x51, 0xd2, 0x05, 0xc1, 0xa9, 0xd4, 0x9e,
	0xae, 0x93, 0x96, 0x96, 0x6f, 0x7c, 0xdf, 0x14, 0x7e, 0x60, 0x0a, 0x9d, 0xf5, 0x55, 0xd6, 0xae,
	0x21, 0xb6, 0x09, 0xe7, 0x27, 0xf1, 0xf8, 0x30, 0xa3, 0x5b, 0xe4, 0x0e, 0x22, 0x27, 0xd6, 0xa4,
	0xd2, 0x39, 0x99, 0xf5, 0x6d, 0x28, 0x77, 0xbb, 0x8e, 0xf8, 0x55, 0x90, 0xfe, 0xaa, 0x14, 0x96,
	0xad, 0xf0, 0x90, 0xac, 0x9a, 0xc2, 0x3b, 0x2f, 0x74, 0xa6, 0xf4, 0xa8, 0x54, 0x71, 0xed, 0x06,
	0xea, 0xac, 0xcc, 0xc9, 0xa2, 0x86, 0xa3, 0x17, 0x64, 0xd9, 0x1b, 0x2f, 0x72, 0xa8, 0x4c, 0x91,
	0x7b, 0xd7, 0xa6, 0x78, 0xf7, 0x5d, 0xbc, 0xfb, 0x97, 0xf4, 0x8b, 0xde, 0x85, 0x64, 0x08, 0x60,
	0xba, 0x18, 0x0f, 0xa4, 0x65, 0x66, 0xc8, 0x22, 0x94, 0x29, 0xcd, 0xfc, 0x85, 0x64, 0x30, 0x27,
	0xf7, 0x5c, 0x3c, 0x65, 0xa9, 0xc9, 0x73, 0x89, 0xed, 0xcf, 0x97, 0x50, 0x91, 0x07, 0x34, 0xfd,
	0x8a, 0x34, 0x8b, 0x49, 0x6e, 0x44, 0x26, 0xb3, 0xfe, 0x50, 0xe5, 0xd2, 0xb5, 0x57, 0xd8, 0xe2,
	0x7a, 0xbd, 0xfb, 0x29, 0xba, 0xda, 0xa0, 0x3f, 0xdb, 0x61, 0xb9, 0x72, 0x1e, 0x1c, 0xa0, 0x98,
	0x95, 0x68, 0x26, 0x1c, 0x9b, 0x08, 0x8b, 0x92, 0xd2, 0x5f, 0xc2, 0x97, 0x4b, 0xf1, 0x01, 0x80,
	0xe9, 0x53, 0x52, 0xcd, 0xcd, 0xc8, 0xb5, 0x57, 0xd9, 0xe2, 0x7a, 0x63, 0xf3, 0x56, 0x6c, 0xab,
	0x23, 0x33, 0x7a, 0x2e, 0x9d, 0x13, 0x23, 0xd9, 0x6d, 0xa3, 0x1b, 0x4a, 0x97, 0x8f, 0xa2, 0x13,
	0x80, 0x27, 0x6b, 0x95, 0x05, 0x8e, 0x8a, 0xf4, 0x01, 0x21, 0x4e, 0x3a, 0xa7, 0x8c, 0x86, 0xaa,
	0x2c, 0x43, 0x9f, 0x94, 0xed, 0x4f, 0xea, 0xa7, 0x41, 0x72, 0xb8, 0xc7, 0xeb, 0x11, 0x74, 0x98,
	0xd1, 0x84, 0x5c, 0x0b, 0x4d, 0xd5, 0x64, 0x95, 0xf5, 0xe6, 0x66, 0x7b, 0x6e, 0xe4, 0xe2, 0x58,
	0x24, 0xd8, 0x60, 0x3c, 0xc0, 0xe8, 0x09, 0xb9, 0x0e, 0x1f, 0x85, 0x6b, 0xdf, 0x44, 0xeb, 0x9f,
	0xa3, 0xf5, 0x87, 0x74, 0x03, 0xd1, 0xda, 0x3b, 0x88, 0x4a, 0x68, 0x26, 0xad, 0x35, 0x96, 0x05,
	0x28, 0x83, 0x7e, 0x62, 0x83, 0x29, 0x5e, 0x3c, 0xcc, 0x44, 0xc2, 0xa3, 0x19, 0xe8, 0xee, 0xc2,
	0x49, 0xdb, 0xd7, 0xc6, 0xab, 0xa1, 0x92, 0x59, 0xbb, 0xc5, 0x2a, 0xeb, 0x35, 0xbe, 0x04, 0x87,
	0xc7, 0xf1, 0x0c, 0x06, 0x07, 0xd8, 0xe8, 0x32, 0x0e, 0xce, 0xad, 0xf7, 0x0c, 0x4e, 0xc0, 0x41,
	0x73, 0x77, 0xb6, 0xc9, 0xb5, 0x30, 0x1c, 0x75, 0x72, 0xed, 0xec, 0xf8, 0x74, 0xbf, 0xd7, 0xfa,
	0x3f, 0xda, 0x20, 0x37, 0xf8, 0xd9, 0xf1, 0xf1, 0xe1, 0xf1, 0xb3, 0x56, 0x85, 0x36, 0x09, 0xe9,
	0xed, 0xf3, 0xe7, 0x87, 0xc7, 0x3b, 0xbd, 0xfd, 0xbd, 0xd6, 0x02, 0xe0, 0xf6, 0x39, 0x3f, 0xe1,
	0xad, 0xc5, 0xce, 0xdf, 0xaf, 0x93, 0xe6, 0xbb, 0xbc, 0x43, 0x15, 0xb9, 0x81, 0x53, 0x69, 0x6c,
	0x20, 0x87, 0xee, 0x09, 0x06, 0x70, 0x48, 0xf7, 0xa0, 0xbf, 0xa4, 0xf6, 0xca, 0xc3, 0x15, 0x85,
	0x67, 0x61, 0x78, 0x33, 0xe6, 0x2f, 0x94, 0x0b, 0x95, 0x66, 0x3d, 0xf8, 0x4c, 0x85, 0x66, 0x03,
	0xc9, 0x04, 0x83, 0xcb, 0x69, 0x31, 0x96, 0xcc, 0x58, 0x76, 0x51, 0x68, 0xcf, 0x54, 0x06, 0x45,
	0x2c, 0xed, 0xd3, 0x29, 0x69, 0x62, 0x3e, 0xa6, 0x7d, 0x6f, 0xfa, 0x00, 0x46, 0xa6, 0xa9, 0x75,
	0x4f, 0xd1, 0xe3, 0x73, 0xba, 0x71, 0x7a, 0x61, 0x8a, 0x3c, 0x63, 0x82, 0x85, 0xb4, 0xa5, 0x02,
	0x7a, 0x15, 0x6c, 0x63, 0xaa, 0xbd, 0xc1, 0x54, 0x2b, 0xad, 0xbc, 0x02, 0x7b, 0x49, 0x67, 0x15,
	0x53, 0x39, 0x65, 0xc2, 0xb3, 0x5d, 0x33, 0x9e, 0xe4, 0x12, 0xf0, 0x9b, 0x55, 0x6f, 0x0b, 0xc9,
	0x97, 0x82, 0xab, 0x9e, 0x39, 0x73, 0xd2, 0xd2, 0x5f, 0x93, 0x7a, 0xa4, 0x39, 0x95, 0x05, 0xa6,
	0xe9, 0x6e, 0xa3, 0xd7, 0x2d, 0x52, 0x0f, 0x04, 0x77, 0xc6, 0x8f, 0xe9, 0x8f, 0x7b, 0xb3, 0x52,
	0x32, 0x35, 0x77, 0x51, 0x66, 0x26, 0xd2, 0x0a, 0x2f, 0x1d, 0x33, 0x1a, 0xfb, 0xb2, 0x16, 0x40,
	0x87, 0x19, 0x1d, 0x90, 0x7a, 0x3a, 0x29, 0xfa, 0xb9, 0x1a, 0xab, 0x48, 0x00, 0xdd, 0x7d, 0x34,
	0xfc, 0x94, 0x6e, 0xc1, 0xd4, 0x8c, 0x95, 0x67, 0x46, 0xcf, 0xf5, 0x08, 0x4b, 0x27, 0x05, 0x73,
	0x32, 0x35, 0x3a, 0x73, 0x90, 0xb1, 0x2c, 0xb4, 0x50, 0x99, 0xd2, 0xb5, 0x4a, 0x65, 0xb3, 0xfa,
	0xf9, 0xe6, 0x83, 0x07, 0xbc, 0x96, 0x4e, 0x8a, 0x23, 0xd0, 0xa7, 0x9e, 0xac, 0xbc, 0xcb, 0x64,
	0xc1, 0xdb, 0x32, 0x7a, 0xdb, 0x43, 0x6f, 0x4f, 0xe8, 0xc6, 0xb7, 0xbc, 0x05, 0x66, 0xf0, 0x56,
	0x0c, 0x87, 0x2a, 0xfd, 0x0e, 0x47, 0xb5, 0x87, 0x0f, 0xc2, 0x3f, 0x7e, 0x6b, 0x9e, 0x10, 0x83,
	0xd7, 0x97, 0xa4, 0x8e, 0xcf, 0x2c, 0x54, 0x14, 0x69, 0xaa, 0xde, 0x7d, 0x82, 0xbe, 0x1e, 0xd3,
	0x47, 0x90, 0xa7, 0x50, 0xe9, 0x61, 0xbc, 0x9a, 0x70, 0x8e, 0x29, 0x28, 0xc1, 0x18, 0x5a, 0x46,
	0x8f, 0x66, 0x74, 0x00, 0x65, 0xb3, 0x45, 0x4c, 0x1b, 0x1c, 0x1c, 0x8b, 0xb1, 0xa4, 0x82, 0x10,
	0x07, 0xaf, 0x65, 0xe8, 0xfc, 0x9b, 0x78, 0x93, 0x2e, 0x5a, 0xff, 0xf9, 0x3b, 0x9d, 0x4f, 0x3f,
	0xd9, 0x33, 0xd0, 0x0d, 0x2c, 0xb2, 0xed, 0x5c, 0x4d, 0x0a, 0xed, 0x55, 0x1e, 0x7e, 0x07, 0x64,
	0xc2, 0xc0, 0x43, 0x1d, 0xad, 0xe2, 0x03, 0xf3, 0x55, 0x7c, 0x44, 0xdb, 0xef, 0x79, 0x44, 0x37,
	0xd1, 0xe5, 0x7d, 0x7a, 0x0f, 0x2e, 0x14, 0x39, 0x79, 0xce, 0xc5, 0x95, 0x70, 0x0c, 0x0d, 0xca,
	0x8c, 0x5d, 0x29, 0x7f, 0x11, 0x38, 0x09, 0x5f, 0xd9, 0x97, 0xe4, 0xf6, 0xf9, 0xdc, 0x5e, 0x00,
	0x43, 0x15, 0xa6, 0x92, 0x92, 0xaa, 0x9f, 0x4e, 0xca, 0x97, 0x16, 0xbf, 0x69, 0x42, 0x6e, 0x4c,
	0xc4, 0x14, 0x38, 0x31, 0xbe, 0xad, 0xff, 0x33, 0x1a, 0x5e, 0x82, 0x3a, 0x7f, 0xaa, 0x86, 0x31,
	0xdd, 0x93, 0x2e, 0xb5, 0x0a, 0xec, 0xc3, 0x4a, 0x80, 0x85, 0x08, 0x33, 0x7a, 0x0f, 0xe3, 0xfe,
	0x98, 0x7e, 0x84, 0x8f, 0x3f, 0x56, 0x42, 0x00, 0xcd, 0x8f, 0x94, 0xf3, 0xd2, 0xca, 0x0c, 0xf8,
	0xff, 0x19, 0xe7, 0x09, 0x47, 0x25, 0xfa, 0x0b, 0xb2, 0x3c, 0xb4, 0x4a, 0xea, 0x2c, 0x9f, 0x86,
	0x72, 0x2e, 0xa0, 0x95, 0x0f, 0xd0, 0xca, 0x6d, 0xba, 0x72, 0x10, 0x85, 0xe1, 0xd6, 0x80, 0x48,
	0xf8, 0x52, 0xa9, 0x81, 0xf5, 0xda, 0x24, 0xb5, 0x54, 0x78, 0x39, 0x32, 0x76, 0x8a, 0x0f, 0x7c,
	0xbd, 0x7b, 0x07, 0x95, 0x5b, 0x14, 0x03, 0x65, 0xa5, 0x30, 0xe1, 0x33, 0x1c, 0xdd, 0x22, 0x8b,
	0x99, 0x49, 0xdb, 0x55, 0x84, 0x77, 0x10, 0xfe, 0x21, 0x5d, 0x43, 0x78, 0x66, 0xd2, 0x02, 0x1a,
	0x25, 0x0c, 0xb8, 0xf3, 0x56, 0xe9, 0x51, 0xc2, 0x01, 0x4e, 0x9f, 0x90, 0x3a, 0x24, 0xb8, 0x8f,
	0x49, 0x0c, 0x93, 0xfa, 0x31, 0xea, 0x7e, 0x40, 0xef, 0xa2, 0xae, 0xb0, 0x23, 0xd4, 0x75, 0x0c,
	0x20, 0x31, 0xda, 0x1a, 0xe8, 0xf4, 0x20, 0xd7, 0xbf, 0x21, 0x4b, 0x99, 0x1c, 0x8a, 0x22, 0xf7,
	0x61, 0xe9, 0xba, 0xfe, 0x9e, 0xf2, 0x7f, 0x84, 0x86, 0xef, 0xd2, 0xff, 0xdf, 0x0b, 0x1a, 0xe1,
	0xfe, 0x33, 0x07, 0x09, 0x6f, 0x44, 0x53, 0xc8, 0x94, 0x8f, 0x09, 0x19, 0xc8, 0x0b, 0x71, 0xa9,
	0x4c, 0x61, 0x1d, 0x2e, 0x12, 0xf5, 0xd9, 0xd3, 0xd5, 0xc2, 0xd0, 0xde, 0x8a, 0x13, 0x3e, 0x87,
	0xa5, 0xc7, 0xa4, 0xa6, 0xb4, 0x07, 0x72, 0xcc, 0x71, 0xa9, 0xa8, 0xcd, 0x1a, 0xef, 0x93, 0xc3,
	0x78, 0x8e, 0xae, 0x1d, 0x13, 0x56, 0x62, 0xbb, 0x5f, 0x2a, 0xa7, 0x06, 0xb9, 0x2c, 0xdf, 0xf1,
	0x67, 0x67, 0x87, 0x09, 0x9f, 0xd9, 0xe8, 0xfc, 0x65, 0x81, 0xb4, 0xce, 0x51, 0xd3, 0x9a, 0x91,
	0xf0, 0x12, 0xc3, 0x1b, 0x92, 0x46, 0xae, 0x46, 0x17, 0xfe, 0x4a, 0xc2, 0xff, 0xd8, 0x28, 0xb5,
	0x19, 0x3b, 0x7c, 0xb6, 0x1b, 0xde, 0x7e, 0x66, 0x74, 0x3e, 0x65, 0x10, 0x6b, 0xb9, 0xee, 0x31,
	0xa5, 0x87, 0xc6, 0x8e, 0x43, 0x25, 0x86, 0xd6, 0x8c, 0xe7, 0xdf, 0xb3, 0xc8, 0xa5, 0xf3, 0x86,
	0xe9, 0x9f, 0x2b, 0xe4, 0xc6, 0xeb, 0x42, 0x5a, 0x85, 0x9b, 0xe2, 0xfc, 0x93, 0x7e, 0xfe, 0xe2,
	0x28, 0x2e, 0x2f, 0xdd, 0xdf, 0xa3, 0xdf, 0x4b, 0xba, 0xbd, 0x93, 0x65, 0x0a, 0x4c, 0x8b, 0x9c,
	0xa5, 0x85, 0xf3, 0x66, 0xcc, 0xce, 0x5f, 0x1c, 0xb1, 0xa8, 0x1f, 0xb9, 0x81, 0x65, 0x05, 0x34,
	0x00, 0x53, 0xb3, 0x1b, 0x29, 0xa3, 0x93, 0xcd, 0xc7, 0x2f, 0x7f, 0xd7, 0x81, 0xe6, 0xeb, 0x6c,
	0xb3, 0xce, 0x2f, 0x8d, 0xf3, 0x50, 0xe9, 0xce, 0x7d, 0xd6, 0x39, 0x7f, 0x71, 0x04, 0x47, 0x4e,
	0xe2, 0x85, 0x0e, 0x5e, 0x67, 0x31, 0x6e, 0xb8, 0xc8, 0xfa, 0x4f, 0x3b, 0x7f, 0xf8, 0x2d, 0x2f,
	0xe3, 0xdb, 0xa6, 0xff, 0xfa, 0x86, 0x35, 0xc9, 0xd2, 0xfc, 0xa4, 0x76, 0xfe, 0x56, 0x21, 0xcd,
	0x9d, 0x83, 0x83, 0x2d, 0xa8, 0xd8, 0xc9, 0xe0, 0x6b, 0x99, 0x7a, 0xda, 0x22, 0x8b, 0x85, 0xd5,
	0x71, 0x64, 0xe1, 0xf3, 0xdb, 0x9b, 0xfb, 0xc2, 0xf7, 0xdd, 0xdc, 0x3f, 0x25, 0x4b, 0x48, 0x9a,
	0x69, 0xd8, 0x32, 0xe2, 0x06, 0x4f, 0xff, 0x7b, 0xff, 0xe0, 0xf8, 0xa7, 0x55, 0xb9, 0xa3, 0x3f,
	0x22, 0x04, 0xd5, 0xc2, 0xd2, 0x52, 0x7d, 0x0f, 0x47, 0x20, 0x27, 0x23, 0xd3, 0x74, 0x1e, 0x91,
	0xc6, 0xf9, 0x8b, 0x23, 0xc8, 0x50, 0x4f, 0x8c, 0x1c, 0x10, 0xcf, 0x5b, 0x86, 0x88, 0x83, 0xdf,
	0x22, 0x8b, 0x5e, 0x8c, 0xc2, 0xb8, 0x73, 0xf8, 0xec, 0x6c, 0x93, 0xd6, 0x9c, 0x52, 0xa0, 0xac,
	0x9f, 0x90, 0xaa, 0x17, 0xf8, 0xf7, 0xc9, 0xe2, 0x5c, 0xb0, 0x73, 0x30, 0x8e, 0xf2, 0xc1, 0x75,
	0x14, 0x3c, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xa7, 0xde, 0x7f, 0x16, 0x0e, 0x00,
	0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/config.proto

/*
Package probes is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/probes/config.proto

It has these top-level messages:
	ProbeDef
*/
package probes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import cloudprober_metrics "github.com/google/cloudprober/metrics"
import cloudprober_probes_http "github.com/google/cloudprober/probes/http"
import cloudprober_probes_dns "github.com/google/cloudprober/probes/dns"
import cloudprober_probes_external "github.com/google/cloudprober/probes/external"
import cloudprober_probes_ping "github.com/google/cloudprober/probes/ping"
import cloudprober_probes_udp "github.com/google/cloudprober/probes/udp"
import cloudprober_probes_udplistener "github.com/google/cloudprober/probes/udplistener"
import cloudprober_targets "github.com/google/cloudprober/targets"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProbeDef_Type int32

const (
	ProbeDef_PING         ProbeDef_Type = 0
	ProbeDef_HTTP         ProbeDef_Type = 1
	ProbeDef_DNS          ProbeDef_Type = 2
	ProbeDef_EXTERNAL     ProbeDef_Type = 3
	ProbeDef_UDP          ProbeDef_Type = 4
	ProbeDef_UDP_LISTENER ProbeDef_Type = 5
	// One of the extension probe types. See "extensions" below for more
	// details.
	ProbeDef_EXTENSION ProbeDef_Type = 98
	// USER_DEFINED probe type is for a one off probe that you want to compile
	// into cloudprober, but you don't expect it to be reused. If you expect
	// it to be reused, you should consider adding it using the extensions
	// mechanism.
	ProbeDef_USER_DEFINED ProbeDef_Type = 99
)

var ProbeDef_Type_name = map[int32]string{
	0:  "PING",
	1:  "HTTP",
	2:  "DNS",
	3:  "EXTERNAL",
	4:  "UDP",
	5:  "UDP_LISTENER",
	98: "EXTENSION",
	99: "USER_DEFINED",
}
var ProbeDef_Type_value = map[string]int32{
	"PING":         0,
	"HTTP":         1,
	"DNS":          2,
	"EXTERNAL":     3,
	"UDP":          4,
	"UDP_LISTENER": 5,
	"EXTENSION":    98,
	"USER_DEFINED": 99,
}

func (x ProbeDef_Type) Enum() *ProbeDef_Type {
	p := new(ProbeDef_Type)
	*p = x
	return p
}
func (x ProbeDef_Type) String() string {
	return proto.EnumName(ProbeDef_Type_name, int32(x))
}
func (x *ProbeDef_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProbeDef_Type_value, data, "ProbeDef_Type")
	if err != nil {
		return err
	}
	*x = ProbeDef_Type(value)
	return nil
}
func (ProbeDef_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ProbeDef struct {
	Name *string        `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Type *ProbeDef_Type `protobuf:"varint,2,req,name=type,enum=cloudprober.probes.ProbeDef_Type" json:"type,omitempty"`
	// Which machines this probe should run on. If defined, cloudprober will run
	// this probe only if machine's hostname matches this value.
	RunOn *string `protobuf:"bytes,3,opt,name=run_on,json=runOn" json:"run_on,omitempty"`
	// Interval between two probes
	IntervalMsec *int32 `protobuf:"varint,4,opt,name=interval_msec,json=intervalMsec,def=2000" json:"interval_msec,omitempty"`
	// Timeout for each probe
	TimeoutMsec *int32 `protobuf:"varint,5,opt,name=timeout_msec,json=timeoutMsec,def=1000" json:"timeout_msec,omitempty"`
	// Targets for the probe
	Targets *cloudprober_targets.TargetsDef `protobuf:"bytes,6,req,name=targets" json:"targets,omitempty"`
	// Latency distribution. If specified, latency is stored as a distribution.
	LatencyDistribution *cloudprober_metrics.Dist `protobuf:"bytes,7,opt,name=latency_distribution,json=latencyDistribution" json:"latency_distribution,omitempty"`
	// Latency unit. Any string that's parseable by time.ParseDuration.
	// Valid values: "ns", "us" (or "µs"), "ms", "s", "m", "h".
	LatencyUnit *string `protobuf:"bytes,8,opt,name=latency_unit,json=latencyUnit,def=us" json:"latency_unit,omitempty"`
	// Types that are valid to be assigned to Probe:
	//	*ProbeDef_PingProbe
	//	*ProbeDef_HttpProbe
	//	*ProbeDef_DnsProbe
	//	*ProbeDef_ExternalProbe
	//	*ProbeDef_UdpProbe
	//	*ProbeDef_UdpListenerProbe
	//	*ProbeDef_UserDefinedProbe
	Probe                        isProbeDef_Probe `protobuf_oneof:"probe"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *ProbeDef) Reset()                    { *m = ProbeDef{} }
func (m *ProbeDef) String() string            { return proto.CompactTextString(m) }
func (*ProbeDef) ProtoMessage()               {}
func (*ProbeDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

var extRange_ProbeDef = []proto.ExtensionRange{
	{200, 536870911},
}

func (*ProbeDef) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ProbeDef
}

const Default_ProbeDef_IntervalMsec int32 = 2000
const Default_ProbeDef_TimeoutMsec int32 = 1000
const Default_ProbeDef_LatencyUnit string = "us"

type isProbeDef_Probe interface {
	isProbeDef_Probe()
}

type ProbeDef_PingProbe struct {
	PingProbe *cloudprober_probes_ping.ProbeConf `protobuf:"bytes,20,opt,name=ping_probe,json=pingProbe,oneof"`
}
type ProbeDef_HttpProbe struct {
	HttpProbe *cloudprober_probes_http.ProbeConf `protobuf:"bytes,21,opt,name=http_probe,json=httpProbe,oneof"`
}
type ProbeDef_DnsProbe struct {
	DnsProbe *cloudprober_probes_dns.ProbeConf `protobuf:"bytes,22,opt,name=dns_probe,json=dnsProbe,oneof"`
}
type ProbeDef_ExternalProbe struct {
	ExternalProbe *cloudprober_probes_external.ProbeConf `protobuf:"bytes,23,opt,name=external_probe,json=externalProbe,oneof"`
}
type ProbeDef_UdpProbe struct {
	UdpProbe *cloudprober_probes_udp.ProbeConf `protobuf:"bytes,24,opt,name=udp_probe,json=udpProbe,oneof"`
}
type ProbeDef_UdpListenerProbe struct {
	UdpListenerProbe *cloudprober_probes_udplistener.ProbeConf `protobuf:"bytes,25,opt,name=udp_listener_probe,json=udpListenerProbe,oneof"`
}
type ProbeDef_UserDefinedProbe struct {
	UserDefinedProbe string `protobuf:"bytes,99,opt,name=user_defined_probe,json=userDefinedProbe,oneof"`
}

func (*ProbeDef_PingProbe) isProbeDef_Probe()        {}
func (*ProbeDef_HttpProbe) isProbeDef_Probe()        {}
func (*ProbeDef_DnsProbe) isProbeDef_Probe()         {}
func (*ProbeDef_ExternalProbe) isProbeDef_Probe()    {}
func (*ProbeDef_UdpProbe) isProbeDef_Probe()         {}
func (*ProbeDef_UdpListenerProbe) isProbeDef_Probe() {}
func (*ProbeDef_UserDefinedProbe) isProbeDef_Probe() {}

func (m *ProbeDef) GetProbe() isProbeDef_Probe {
	if m != nil {
		return m.Probe
	}
	return nil
}

func (m *ProbeDef) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProbeDef) GetType() ProbeDef_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ProbeDef_PING
}

func (m *ProbeDef) GetRunOn() string {
	if m != nil && m.RunOn != nil {
		return *m.RunOn
	}
	return ""
}

func (m *ProbeDef) GetIntervalMsec() int32 {
	if m != nil && m.IntervalMsec != nil {
		return *m.IntervalMsec
	}
	return Default_ProbeDef_IntervalMsec
}

func (m *ProbeDef) GetTimeoutMsec() int32 {
	if m != nil && m.TimeoutMsec != nil {
		return *m.TimeoutMsec
	}
	return Default_ProbeDef_TimeoutMsec
}

func (m *ProbeDef) GetTargets() *cloudprober_targets.TargetsDef {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *ProbeDef) GetLatencyDistribution() *cloudprober_metrics.Dist {
	if m != nil {
		return m.LatencyDistribution
	}
	return nil
}

func (m *ProbeDef) GetLatencyUnit() string {
	if m != nil && m.LatencyUnit != nil {
		return *m.LatencyUnit
	}
	return Default_ProbeDef_LatencyUnit
}

func (m *ProbeDef) GetPingProbe() *cloudprober_probes_ping.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_PingProbe); ok {
		return x.PingProbe
	}
	return nil
}

func (m *ProbeDef) GetHttpProbe() *cloudprober_probes_http.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_HttpProbe); ok {
		return x.HttpProbe
	}
	return nil
}

func (m *ProbeDef) GetDnsProbe() *cloudprober_probes_dns.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_DnsProbe); ok {
		return x.DnsProbe
	}
	return nil
}

func (m *ProbeDef) GetExternalProbe() *cloudprober_probes_external.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_ExternalProbe); ok {
		return x.ExternalProbe
	}
	return nil
}

func (m *ProbeDef) GetUdpProbe() *cloudprober_probes_udp.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_UdpProbe); ok {
		return x.UdpProbe
	}
	return nil
}

func (m *ProbeDef) GetUdpListenerProbe() *cloudprober_probes_udplistener.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_UdpListenerProbe); ok {
		return x.UdpListenerProbe
	}
	return nil
}

func (m *ProbeDef) GetUserDefinedProbe() string {
	if x, ok := m.GetProbe().(*ProbeDef_UserDefinedProbe); ok {
		return x.UserDefinedProbe
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ProbeDef) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ProbeDef_OneofMarshaler, _ProbeDef_OneofUnmarshaler, _ProbeDef_OneofSizer, []interface{}{
		(*ProbeDef_PingProbe)(nil),
		(*ProbeDef_HttpProbe)(nil),
		(*ProbeDef_DnsProbe)(nil),
		(*ProbeDef_ExternalProbe)(nil),
		(*ProbeDef_UdpProbe)(nil),
		(*ProbeDef_UdpListenerProbe)(nil),
		(*ProbeDef_UserDefinedProbe)(nil),
	}
}

func _ProbeDef_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ProbeDef)
	// probe
	switch x := m.Probe.(type) {
	case *ProbeDef_PingProbe:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PingProbe); err != nil {
			return err
		}
	case *ProbeDef_HttpProbe:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpProbe); err != nil {
			return err
		}
	case *ProbeDef_DnsProbe:
		b.EncodeVarint(22<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DnsProbe); err != nil {
			return err
		}
	case *ProbeDef_ExternalProbe:
		b.EncodeVarint(23<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExternalProbe); err != nil {
			return err
		}
	case *ProbeDef_UdpProbe:
		b.EncodeVarint(24<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UdpProbe); err != nil {
			return err
		}
	case *ProbeDef_UdpListenerProbe:
		b.EncodeVarint(25<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UdpListenerProbe); err != nil {
			return err
		}
	case *ProbeDef_UserDefinedProbe:
		b.EncodeVarint(99<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.UserDefinedProbe)
	case nil:
	default:
		return fmt.Errorf("ProbeDef.Probe has unexpected type %T", x)
	}
	return nil
}

func _ProbeDef_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ProbeDef)
	switch tag {
	case 20: // probe.ping_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_ping.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_PingProbe{msg}
		return true, err
	case 21: // probe.http_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_http.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_HttpProbe{msg}
		return true, err
	case 22: // probe.dns_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_dns.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_DnsProbe{msg}
		return true, err
	case 23: // probe.external_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_external.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_ExternalProbe{msg}
		return true, err
	case 24: // probe.udp_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_udp.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_UdpProbe{msg}
		return true, err
	case 25: // probe.udp_listener_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_udplistener.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_UdpListenerProbe{msg}
		return true, err
	case 99: // probe.user_defined_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Probe = &ProbeDef_UserDefinedProbe{x}
		return true, err
	default:
		return false, nil
	}
}

func _ProbeDef_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ProbeDef)
	// probe
	switch x := m.Probe.(type) {
	case *ProbeDef_PingProbe:
		s := proto.Size(x.PingProbe)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_HttpProbe:
		s := proto.Size(x.HttpProbe)
		n += proto.SizeVarint(21<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_DnsProbe:
		s := proto.Size(x.DnsProbe)
		n += proto.SizeVarint(22<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_ExternalProbe:
		s := proto.Size(x.ExternalProbe)
		n += proto.SizeVarint(23<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_UdpProbe:
		s := proto.Size(x.UdpProbe)
		n += proto.SizeVarint(24<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_UdpListenerProbe:
		s := proto.Size(x.UdpListenerProbe)
		n += proto.SizeVarint(25<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_UserDefinedProbe:
		n += proto.SizeVarint(99<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.UserDefinedProbe)))
		n += len(x.UserDefinedProbe)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ProbeDef)(nil), "cloudprober.probes.ProbeDef")
	proto.RegisterEnum("cloudprober.probes.ProbeDef_Type", ProbeDef_Type_name, ProbeDef_Type_value)
}

func init() { proto.RegisterFile("github.com/google/cloudprober/probes/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0x87, 0xd7, 0x2c, 0x5d, 0x53, 0xb7, 0x9b, 0x22, 0xb3, 0x41, 0xd6, 0x1b, 0xca, 0x24, 0xa0,
	0xe3, 0x22, 0xed, 0x8a, 0x86, 0xc4, 0x24, 0x24, 0x60, 0x09, 0xac, 0x52, 0xc9, 0xaa, 0xb4, 0x93,
	0xe0, 0x2a, 0x6a, 0x63, 0xb7, 0xb3, 0xd4, 0x3a, 0x91, 0xff, 0x20, 0x76, 0xd7, 0xc7, 0xe3, 0x25,
	0x78, 0x17, 0x64, 0x27, 0x41, 0x0b, 0x6c, 0x53, 0xc5, 0x55, 0xa2, 0xe3, 0xef, 0xf7, 0xf9, 0x9c,
	0xc4, 0x06, 0x27, 0x0b, 0x22, 0xae, 0xe5, 0xcc, 0x8d, 0x93, 0x55, 0x77, 0x91, 0x24, 0x8b, 0x25,
	0xee, 0xc6, 0xcb, 0x44, 0xa2, 0x94, 0x25, 0x33, 0xcc, 0xba, 0xfa, 0xc1, 0xbb, 0x71, 0x42, 0xe7,
	0x64, 0xe1, 0xa6, 0x2c, 0x11, 0x09, 0x84, 0xb7, 0x00, 0x37, 0x03, 0x5a, 0xbd, 0x87, 0x35, 0x2b,
	0x2c, 0x18, 0x89, 0x79, 0x17, 0x11, 0x2e, 0x32, 0x4b, 0xeb, 0xcd, 0x46, 0x1b, 0x5f, 0x0b, 0x91,
	0x96, 0x76, 0x6f, 0x9d, 0x6e, 0x94, 0x43, 0xb4, 0xdc, 0x74, 0xeb, 0x6c, 0xa3, 0x18, 0xfe, 0x21,
	0x30, 0xa3, 0xd3, 0x65, 0x39, 0xbb, 0x59, 0xab, 0x29, 0xa1, 0x8b, 0xff, 0x69, 0x55, 0xa2, 0xbf,
	0x26, 0x7c, 0xb7, 0x69, 0x6c, 0x49, 0xb8, 0xc0, 0x14, 0xb3, 0x72, 0xfc, 0xf5, 0xc3, 0x71, 0x31,
	0x65, 0x0b, 0x2c, 0x78, 0xf1, 0xcc, 0x42, 0x47, 0xbf, 0x6a, 0xc0, 0x1a, 0x29, 0xc0, 0xc3, 0x73,
	0x08, 0x81, 0x49, 0xa7, 0x2b, 0xec, 0x54, 0xda, 0x46, 0xa7, 0x1e, 0xea, 0x77, 0x78, 0x0a, 0x4c,
	0x71, 0x93, 0x62, 0xc7, 0x68, 0x1b, 0x9d, 0xbd, 0xfe, 0x33, 0xf7, 0xdf, 0x33, 0xe0, 0x16, 0x79,
	0x77, 0x72, 0x93, 0xe2, 0x50, 0xe3, 0xf0, 0x00, 0xec, 0x30, 0x49, 0xa3, 0x84, 0x3a, 0xdb, 0xed,
	0x4a, 0xa7, 0x1e, 0x56, 0x99, 0xa4, 0x97, 0x14, 0x1e, 0x83, 0x5d, 0x42, 0x05, 0x66, 0xdf, 0xa7,
	0xcb, 0x68, 0xc5, 0x71, 0xec, 0x98, 0xed, 0x4a, 0xa7, 0x7a, 0x66, 0xf6, 0x7b, 0xbd, 0x5e, 0xd8,
	0x2c, 0x96, 0xbe, 0x70, 0x1c, 0xc3, 0x97, 0xa0, 0x29, 0xc8, 0x0a, 0x27, 0x52, 0x64, 0x64, 0x35,
	0x23, 0x4f, 0x14, 0xd9, 0xc8, 0x57, 0x34, 0xf8, 0x16, 0xd4, 0xf2, 0x99, 0x9c, 0x9d, 0xb6, 0xd1,
	0x69, 0xf4, 0x9f, 0x96, 0x9a, 0x2c, 0xe6, 0x9d, 0x64, 0x4f, 0x0f, 0xcf, 0xc3, 0x82, 0x87, 0x43,
	0xb0, 0xbf, 0x9c, 0x0a, 0x4c, 0xe3, 0x9b, 0x48, 0x9d, 0x50, 0x46, 0x66, 0x52, 0x90, 0x84, 0x3a,
	0xb5, 0x76, 0xa5, 0xd3, 0xe8, 0x1f, 0x96, 0x3c, 0xf9, 0x51, 0x76, 0x3d, 0xc2, 0x45, 0xf8, 0x28,
	0x8f, 0x79, 0xb7, 0x52, 0xf0, 0x39, 0x68, 0x16, 0x36, 0x49, 0x89, 0x70, 0x2c, 0x35, 0xf9, 0x99,
	0x21, 0x79, 0xd8, 0xc8, 0xeb, 0x57, 0x94, 0x08, 0x78, 0x0e, 0x80, 0x3a, 0x32, 0x91, 0xf6, 0x3a,
	0xfb, 0x7a, 0xab, 0xa3, 0xbb, 0xbe, 0xab, 0xa2, 0xb2, 0x8f, 0x7b, 0x9e, 0xd0, 0xf9, 0xc5, 0x56,
	0x58, 0x57, 0x15, 0x5d, 0x50, 0x12, 0x75, 0x45, 0x72, 0xc9, 0xc1, 0xfd, 0x12, 0x45, 0x95, 0x25,
	0xaa, 0x92, 0x49, 0xde, 0x83, 0x3a, 0xa2, 0x3c, 0x77, 0x3c, 0xd6, 0x8e, 0x3b, 0x7f, 0x30, 0xa2,
	0xbc, 0xa4, 0xb0, 0x10, 0xe5, 0x99, 0xe1, 0x12, 0xec, 0x15, 0x57, 0x27, 0xd7, 0x3c, 0xd1, 0x9a,
	0x17, 0x77, 0x69, 0x0a, 0xb2, 0xe4, 0xda, 0x2d, 0xaa, 0x7f, 0x5a, 0x92, 0xa8, 0x18, 0xcb, 0xb9,
	0xbf, 0x25, 0x89, 0xca, 0x53, 0x59, 0x12, 0xe5, 0x43, 0x7d, 0x03, 0x50, 0x19, 0x8a, 0x3b, 0x92,
	0xab, 0x0e, 0xb5, 0xea, 0xf8, 0x1e, 0x55, 0x01, 0x97, 0x94, 0xb6, 0x44, 0xe9, 0x30, 0x5f, 0xc8,
	0xd4, 0x2e, 0x80, 0x92, 0x63, 0x16, 0x21, 0x3c, 0x27, 0x14, 0xa3, 0x5c, 0x1d, 0xab, 0xdf, 0xac,
	0x79, 0x8e, 0x99, 0x97, 0x2d, 0x69, 0xfe, 0x68, 0x05, 0x4c, 0x75, 0x25, 0xa0, 0x05, 0xcc, 0xd1,
	0x20, 0xf8, 0x6c, 0x6f, 0xa9, 0xb7, 0x8b, 0xc9, 0x64, 0x64, 0x57, 0x60, 0x0d, 0x6c, 0x7b, 0xc1,
	0xd8, 0x36, 0x60, 0x13, 0x58, 0xfe, 0xd7, 0x89, 0x1f, 0x06, 0x1f, 0x86, 0xf6, 0xb6, 0x2a, 0x5f,
	0x79, 0x23, 0xdb, 0x84, 0x36, 0x68, 0x5e, 0x79, 0xa3, 0x68, 0x38, 0x18, 0x4f, 0xfc, 0xc0, 0x0f,
	0xed, 0x2a, 0xdc, 0x05, 0x75, 0x05, 0x06, 0xe3, 0xc1, 0x65, 0x60, 0xcf, 0x34, 0x30, 0xf6, 0xc3,
	0xc8, 0xf3, 0x3f, 0x0d, 0x02, 0xdf, 0xb3, 0xe3, 0x57, 0x75, 0xeb, 0x67, 0xc5, 0x5e, 0xaf, 0xd7,
	0x6b, 0xe3, 0x63, 0x0d, 0x54, 0x75, 0x73, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x08, 0x02,
	0x12, 0xe7, 0x05, 0x00, 0x00,
}

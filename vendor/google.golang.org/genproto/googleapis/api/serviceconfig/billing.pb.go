// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/billing.proto
// DO NOT EDIT!

package google_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/metric"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Billing related configuration of the service.
//
// The following example shows how to configure metrics for billing:
//
//     metrics:
//     - name: library.googleapis.com/read_calls
//       metric_kind: DELTA
//       value_type: INT64
//     - name: library.googleapis.com/write_calls
//       metric_kind: DELTA
//       value_type: INT64
//     billing:
//       metrics:
//       - library.googleapis.com/read_calls
//       - library.googleapis.com/write_calls
//
// The next example shows how to enable billing status check and customize the
// check behavior. It makes sure billing status check is included in the `Check`
// method of [Service Control API](https://cloud.google.com/service-control/).
// In the example, "google.storage.Get" method can be served when the billing
// status is either `current` or `delinquent`, while "google.storage.Write"
// method can only be served when the billing status is `current`:
//
//     billing:
//       rules:
//       - selector: google.storage.Get
//         allowed_statuses:
//         - current
//         - delinquent
//       - selector: google.storage.Write
//         allowed_statuses: current
//
// Mostly services should only allow `current` status when serving requests.
// In addition, services can choose to allow both `current` and `delinquent`
// statuses when serving read-only requests to resources. If there's no
// matching selector for operation, no billing status check will be performed.
//
type Billing struct {
	// Names of the metrics to report to billing. Each name must
	// be defined in [Service.metrics][google.api.Service.metrics] section.
	Metrics []string `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	// A list of billing status rules for configuring billing status check.
	Rules []*BillingStatusRule `protobuf:"bytes,5,rep,name=rules" json:"rules,omitempty"`
}

func (m *Billing) Reset()                    { *m = Billing{} }
func (m *Billing) String() string            { return proto.CompactTextString(m) }
func (*Billing) ProtoMessage()               {}
func (*Billing) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Billing) GetRules() []*BillingStatusRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Defines the billing status requirements for operations.
//
// When used with
// [Service Control API](https://cloud.google.com/service-control/), the
// following statuses are supported:
//
// - **current**: the associated billing account is up to date and capable of
//                paying for resource usages.
// - **delinquent**: the associated billing account has a correctable problem,
//                   such as late payment.
//
// Mostly services should only allow `current` status when serving requests.
// In addition, services can choose to allow both `current` and `delinquent`
// statuses when serving read-only requests to resources. If the list of
// allowed_statuses is empty, it means no billing requirement.
//
type BillingStatusRule struct {
	// Selects the operation names to which this rule applies.
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// Allowed billing statuses. The billing status check passes if the actual
	// billing status matches any of the provided values here.
	AllowedStatuses []string `protobuf:"bytes,2,rep,name=allowed_statuses,json=allowedStatuses" json:"allowed_statuses,omitempty"`
}

func (m *BillingStatusRule) Reset()                    { *m = BillingStatusRule{} }
func (m *BillingStatusRule) String() string            { return proto.CompactTextString(m) }
func (*BillingStatusRule) ProtoMessage()               {}
func (*BillingStatusRule) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func init() {
	proto.RegisterType((*Billing)(nil), "google.api.Billing")
	proto.RegisterType((*BillingStatusRule)(nil), "google.api.BillingStatusRule")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/billing.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x4f, 0xbd, 0x4b, 0x43, 0x31,
	0x10, 0xe7, 0x29, 0xb5, 0x7a, 0x8a, 0x1f, 0x99, 0x1e, 0x0f, 0x84, 0xd2, 0x49, 0x97, 0x17, 0xb0,
	0xb3, 0xcb, 0x03, 0x07, 0xb7, 0xf2, 0xba, 0x88, 0x8b, 0xa4, 0xf1, 0x0c, 0x81, 0x34, 0x57, 0x92,
	0x54, 0xff, 0x7d, 0xcf, 0x24, 0x7e, 0x80, 0x8b, 0xb8, 0x24, 0xdc, 0xdd, 0xef, 0x13, 0xee, 0x0c,
	0x91, 0x71, 0xd8, 0x1b, 0x72, 0xca, 0x9b, 0x9e, 0x82, 0x91, 0x06, 0xfd, 0x36, 0x50, 0x22, 0x59,
	0x4e, 0x6a, 0x6b, 0xa3, 0xe4, 0x47, 0x46, 0x0c, 0xaf, 0x56, 0xa3, 0x26, 0xff, 0x62, 0x8d, 0x5c,
	0x5b, 0xe7, 0x2c, 0x33, 0x32, 0x54, 0x40, 0x95, 0x61, 0x5c, 0x77, 0xff, 0x5f, 0x49, 0xe5, 0x3d,
	0x25, 0x95, 0x2c, 0xf9, 0x58, 0x64, 0xbb, 0xdb, 0xbf, 0x4b, 0x6d, 0x30, 0x05, 0xab, 0xeb, 0x57,
	0xe8, 0xf3, 0x07, 0x98, 0x0e, 0x25, 0xa6, 0x68, 0x61, 0x5a, 0x4e, 0xb1, 0x6d, 0x66, 0xfb, 0x57,
	0x47, 0xe3, 0xe7, 0x28, 0x16, 0x30, 0x09, 0x3b, 0x87, 0xb1, 0x9d, 0xf0, 0xfe, 0xf8, 0xe6, 0xb2,
	0xff, 0xae, 0xd2, 0x57, 0xf6, 0x8a, 0x53, 0xed, 0xe2, 0xc8, 0xa8, 0xb1, 0x60, 0xe7, 0x8f, 0x70,
	0xf1, 0xeb, 0x26, 0x3a, 0x38, 0x8c, 0xe8, 0x50, 0x27, 0x0a, 0x6c, 0xd2, 0xb0, 0xc9, 0xd7, 0x2c,
	0xae, 0xe1, 0x5c, 0x39, 0x47, 0x6f, 0xf8, 0xfc, 0x14, 0x33, 0x83, 0x0d, 0xf7, 0x72, 0x90, 0xb3,
	0xba, 0x5f, 0xd5, 0xf5, 0x30, 0x83, 0x53, 0x4d, 0x9b, 0x1f, 0x31, 0x86, 0x93, 0xea, 0xb5, 0xfc,
	0x68, 0xb5, 0x6c, 0xd6, 0x07, 0xb9, 0xde, 0xe2, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xe1, 0x19,
	0xb1, 0xbd, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/authorization.proto

// Istio Authorization Policy enables access control on workloads in the mesh.
//
// For example, the following authorization policy applies to workloads matched with
// label selector "app: httpbin, version: v1".
//
// It allows requests from:
// - service account "cluster.local/ns/default/sa/sleep" or
// - namespace "test"
// to access the workload with:
// - "GET" method at paths of prefix "/info" or,
// - "POST" method at path "/data".
// when the request has a valid JWT token issued by "https://accounts.google.com".
//
// Any other requests will be rejected.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: httpbin
//  namespace: foo
// spec:
//  selector:
//    matchLabels:
//      app: httpbin
//      version: v1
//  rules:
//  - from:
//    - source:
//        principals: ["cluster.local/ns/default/sa/sleep"]
//    - source:
//        namespaces: ["test"]
//    to:
//    - operation:
//        methods: ["GET"]
//        paths: ["/info*"]
//    - operation:
//        methods: ["POST"]
//        paths: ["/data"]
//    when:
//    - key: request.auth.claims[iss]
//      values: ["https://accounts.google.com"]
// ```
//
// Access control is enabled on a workload if there is any authorization policies selecting
// the workload. When access control is enabled, the default behavior is deny (deny-by-default)
// which means requests to the workload will be rejected if the request is not allowed by any of
// the authorization policies selecting the workload.
//
// Currently AuthorizationPolicy only supports "ALLOW" action. This means that
// if multiple authorization policies apply to the same workload, the effect is additive.
//
// Authorization Policy scope (target) is determined by "metadata/namespace" and
// an optional "selector".
// - "metadata/namespace" tells which namespace the policy applies. If set to root
// namespace, the policy applies to all namespaces in a mesh.
// - workload "selector" can be used to further restrict where a policy applies.
//
// For example,
//
// The following authorization policy applies to workloads containing label
// "app: httpbin" in namespace bar.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: policy
//  namespace: bar
// spec:
//  selector:
//    matchLabels:
//      app: httpbin
// ```
//
// The following authorization policy applies to all workloads in namespace foo.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: policy
//  namespace: foo
// spec:
// ```
//
// The following authorization policy applies to workloads containing label
// "version: v1" in all namespaces in the mesh. (Assuming the root namespace is
// configured to "istio-config").
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: policy
//  namespace: istio-config
// spec:
//  selector:
//    matchLabels:
//      version: v1
// ```

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/api/type/v1beta1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using AuthorizationPolicy within kubernetes types, where deepcopy-gen is used.
func (in *AuthorizationPolicy) DeepCopyInto(out *AuthorizationPolicy) {
	proto.Merge(out, in)
}
package introspection

import "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"

//should be called ApiIntrospector instead.
// the act is introspection. Actor is introspector. like stringer
type APIIntrospection interface {
	GetServices() []string
	GetServiceInfo(serviceID string) core.MethodResult
	GetCheckSum() string
	GetIntrospectionServices() []core.ApiInterface
	GetOperations(serviceID string) core.MethodResult
	GetOperationInfo(serviceID string, operationID string) core.MethodResult
}
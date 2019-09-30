package common

import (
	"github.com/google/uuid"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
)

/**
Generate random UUID based on RFC 4122
*/
func NewOpId() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Error(err)
		log.Info("Error generating UUID. Trying again ..")
		return NewOpId()
	}
	return uuid.String()
}

func NewDefaultApplicationContext() *core.ApplicationContext {
	appContext := core.NewApplicationContext(nil)
	InsertOperationId(appContext)
	return appContext
}

func InsertOperationId(appContext *core.ApplicationContext) {
	opId := NewOpId()
	appContext.SetProperty(lib.OPID, &opId)
}

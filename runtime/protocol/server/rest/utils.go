/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package rest

import (
	"net/http"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
	httpStatusCodes "github.com/vmware/vsphere-automation-sdk-go/runtime/lib/rest"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
)

var runtimePropertiesToVapiErrorMap = map[string]string{
	"com.vmware.vapi.rest.unsupported_media_type":           "com.vmware.vapi.std.errors.invalid_request",
	"vapi.protocol.server.rest.param.internal_server_error": "com.vmware.vapi.std.errors.internal_server_error",
	"vapi.protocol.server.rest.response.not_structure":      "com.vmware.vapi.std.errors.internal_server_error",
}

func returnError(err []error, rw http.ResponseWriter) {
	if l10nErr, ok := err[0].(*l10n.Error); ok {
		if val, ok := runtimePropertiesToVapiErrorMap[l10nErr.ID()]; ok {
			returnBadRequest(val, rw, err)
			return
		}
	}
	returnBadRequest("com.vmware.vapi.std.errors.invalid_argument", rw, err)
}

func returnBadRequest(errorStr string, rw http.ResponseWriter, dataErr []error) {
	var errorValue *data.ErrorValue
	var status int

	errorValue = bindings.CreateErrorValueFromMessages(bindings.ERROR_MAP[errorStr], dataErr)
	status = httpStatusCodes.VAPI_TO_HTTP_ERROR_MAP[errorStr]
	responseBody, err := getResponseBody(errorValue)
	if err != nil {
		log.Errorf("Error while setting error response body: %s", err)
	}
	rw.WriteHeader(status)
	rw.Write([]byte(responseBody))
}

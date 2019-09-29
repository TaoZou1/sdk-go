/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.cis.task.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package task

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "time"
)


// The ``Status`` enumeration class defines the status values that can be reported for an operation.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type Status string

const (
    // The operation is in pending state.
     Status_PENDING Status = "PENDING"
    // The operation is in progress.
     Status_RUNNING Status = "RUNNING"
    // The operation is blocked.
     Status_BLOCKED Status = "BLOCKED"
    // The operation completed successfully.
     Status_SUCCEEDED Status = "SUCCEEDED"
    // The operation failed.
     Status_FAILED Status = "FAILED"
)

func (s Status) Status() bool {
    switch s {
        case Status_PENDING:
            return true
        case Status_RUNNING:
            return true
        case Status_BLOCKED:
            return true
        case Status_SUCCEEDED:
            return true
        case Status_FAILED:
            return true
        default:
            return false
    }
}





// The ``Progress`` class contains information describe the progress of an operation.
type Progress struct {
    Total int64
    Completed int64
    Message std.LocalizableMessage
}






// The ``CommonInfo`` class contains information common to all tasks.
type CommonInfo struct {
    Description std.LocalizableMessage
    Service string
    Operation string
    Parent *string
    Target *std.DynamicID
    Status Status
    Cancelable bool
    Error *data.ErrorValue
    StartTime *time.Time
    EndTime *time.Time
    User *string
}






// The ``Info`` class contains information about a task.
type Info struct {
    Progress *Progress
    Result data.DataValue
    Description std.LocalizableMessage
    Service string
    Operation string
    Parent *string
    Target *std.DynamicID
    Status Status
    Cancelable bool
    Error *data.ErrorValue
    StartTime *time.Time
    EndTime *time.Time
    User *string
}










func ProgressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["total"] = bindings.NewIntegerType()
    fieldNameMap["total"] = "Total"
    fields["completed"] = bindings.NewIntegerType()
    fieldNameMap["completed"] = "Completed"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.task.progress",fields, reflect.TypeOf(Progress{}), fieldNameMap, validators)
}

func CommonInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service"] = "Service"
    fields["operation"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["operation"] = "Operation"
    fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""))
    fieldNameMap["parent"] = "Parent"
    fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
    fieldNameMap["target"] = "Target"
    fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(Status(Status_PENDING)))
    fieldNameMap["status"] = "Status"
    fields["cancelable"] = bindings.NewBooleanType()
    fieldNameMap["cancelable"] = "Cancelable"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.cis.task.common_info",fields, reflect.TypeOf(CommonInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(ProgressBindingType))
    fieldNameMap["progress"] = "Progress"
    fields["result"] = bindings.NewOptionalType(bindings.NewOpaqueType())
    fieldNameMap["result"] = "Result"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service"] = "Service"
    fields["operation"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["operation"] = "Operation"
    fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""))
    fieldNameMap["parent"] = "Parent"
    fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
    fieldNameMap["target"] = "Target"
    fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(Status(Status_PENDING)))
    fieldNameMap["status"] = "Status"
    fields["cancelable"] = bindings.NewBooleanType()
    fieldNameMap["cancelable"] = "Cancelable"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", true),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", true),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.cis.task.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


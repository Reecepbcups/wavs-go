package types

import (
	wavs "github.com/reecepbcups/wavs-go/wavs/worker/layer-trigger-world"
	"go.bytecodealliance.org/cm"
)

// TriggerResult is the return type for wavs trigger world Run function
// an alias of wavs Exports Run() result type
type TriggerResult = cm.Result[wavs.OptionListU8Shape, cm.Option[cm.List[uint8]], string]

// Ok (alias) returns an Ok response of type Some
func Ok(resp []byte) TriggerResult {
	return cm.OK[TriggerResult](cm.Some(cm.NewList(&resp[0], len(resp))))
}

// OkNone (alias) returns an Ok response of type None
func OkNone() TriggerResult {
	return cm.OK[TriggerResult](cm.None[cm.List[uint8]]())
}

// Err (alias) returns an Err response of type string
func Err(err string) TriggerResult {
	return cm.Err[TriggerResult](err)
}

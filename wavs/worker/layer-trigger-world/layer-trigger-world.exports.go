// Code generated by wit-bindgen-go. DO NOT EDIT.

package layertriggerworld

import (
	"go.bytecodealliance.org/cm"
)

// Exports represents the caller-defined exports from "wavs:worker/layer-trigger-world@0.3.0-rc1".
var Exports struct {
	// Run represents the caller-defined, exported function "run".
	//
	//	run: func(trigger-action: trigger-action) -> result<option<list<u8>>, string>
	Run func(triggerAction TriggerAction) (result cm.Result[OptionListU8Shape, cm.Option[cm.List[uint8]], string])
}

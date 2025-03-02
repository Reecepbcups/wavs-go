package types

import (
	wavs "github.com/reecepbcups/wavs-go/wavs/worker/layer-trigger-world"
	"go.bytecodealliance.org/cm"
)

// TriggerResult is the return type for wavs trigger world Run function
// an alias of wavs Exports Run() result type
type TriggerResult = cm.Result[wavs.OptionListU8Shape, cm.Option[cm.List[uint8]], string]

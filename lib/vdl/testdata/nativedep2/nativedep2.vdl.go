// This file was auto-generated by the veyron vdl tool.
// Source: nativedep2.vdl

package nativedep2

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"time"
	_ "v.io/core/veyron/lib/vdl/testdata/nativetest"
)

type MyTime time.Time

func (MyTime) __VDLReflect(struct {
	Name string "v.io/core/veyron/lib/vdl/testdata/nativedep2.MyTime"
}) {
}

func init() {
	vdl.Register((*MyTime)(nil))
}

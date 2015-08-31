// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package clock

import (
	// VDL system imports
	"v.io/v23/vdl"
)

// ClockData is the persistent state of syncbase clock used to estimate current
// NTP time and catch any unexpected changes to system clock.
type ClockData struct {
	// UTC time in unix nano seconds obtained from system clock at boot.
	SystemTimeAtBoot int64
	// Skew between the system clock and NTP time.
	Skew int64
	// The elapsed time since boot as last seen during a run of clockservice.
	// This is used to determine if the device rebooted since the last run of
	// clockservice.
	ElapsedTimeSinceBoot int64
}

func (ClockData) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/clock.ClockData"`
}) {
}

func init() {
	vdl.Register((*ClockData)(nil))
}

// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: profile.vdl

// Package profile defines types for the implementation of Vanadium profiles.
package profile

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/services/build"
)

// Library describes a shared library that applications may use.
type Library struct {
	// Name is the name of the library.
	Name string
	// MajorVersion is the major version of the library.
	MajorVersion string
	// MinorVersion is the minor version of the library.
	MinorVersion string
}

func (Library) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/profile.Library"`
}) {
}

func (m *Library) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_profile_v_io_x_ref_services_profile_Library == nil || __VDLTypeprofile0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	var2 := (m.Name == "")
	if !var2 {
		keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("Name")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget4.FromString(string(m.Name), vdl.StringType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
				return err
			}
		}
	}
	var5 := (m.MajorVersion == "")
	if !var5 {
		keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("MajorVersion")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget7.FromString(string(m.MajorVersion), vdl.StringType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
				return err
			}
		}
	}
	var8 := (m.MinorVersion == "")
	if !var8 {
		keyTarget9, fieldTarget10, err := fieldsTarget1.StartField("MinorVersion")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget10.FromString(string(m.MinorVersion), vdl.StringType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget9, fieldTarget10); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Library) MakeVDLTarget() vdl.Target {
	return nil
}

func (m *Library) IsZero() bool {

	var1 := (*m == Library{})
	return var1
}

// Specification is how we represent a profile internally. It should
// provide enough information to allow matching of binaries to devices.
type Specification struct {
	// Label is a human-friendly concise label for the profile,
	// e.g. "linux-media".
	Label string
	// Description is a human-friendly description of the profile.
	Description string
	// Arch is the target hardware architecture of the profile.
	Arch build.Architecture
	// Os is the target operating system of the profile.
	Os build.OperatingSystem
	// Format is the file format supported by the profile.
	Format build.Format
	// Libraries is a set of libraries the profile requires.
	Libraries map[Library]struct{}
}

func (Specification) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/profile.Specification"`
}) {
}

func (m *Specification) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_profile_v_io_x_ref_services_profile_Specification == nil || __VDLTypeprofile1 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	var2 := (m.Label == "")
	if !var2 {
		keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("Label")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget4.FromString(string(m.Label), vdl.StringType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
				return err
			}
		}
	}
	var5 := (m.Description == "")
	if !var5 {
		keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Description")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget7.FromString(string(m.Description), vdl.StringType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
				return err
			}
		}
	}
	var8 := m.Arch.IsZero()
	if !var8 {
		keyTarget9, fieldTarget10, err := fieldsTarget1.StartField("Arch")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {

			if err := m.Arch.FillVDLTarget(fieldTarget10, __VDLType_profile_v_io_v23_services_build_Architecture); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget9, fieldTarget10); err != nil {
				return err
			}
		}
	}
	var11 := m.Os.IsZero()
	if !var11 {
		keyTarget12, fieldTarget13, err := fieldsTarget1.StartField("Os")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {

			if err := m.Os.FillVDLTarget(fieldTarget13, __VDLType_profile_v_io_v23_services_build_OperatingSystem); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget12, fieldTarget13); err != nil {
				return err
			}
		}
	}
	var14 := m.Format.IsZero()
	if !var14 {
		keyTarget15, fieldTarget16, err := fieldsTarget1.StartField("Format")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {

			if err := m.Format.FillVDLTarget(fieldTarget16, __VDLType_profile_v_io_v23_services_build_Format); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget15, fieldTarget16); err != nil {
				return err
			}
		}
	}
	var var17 bool
	if len(m.Libraries) == 0 {
		var17 = true
	}
	if !var17 {
		keyTarget18, fieldTarget19, err := fieldsTarget1.StartField("Libraries")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {

			setTarget20, err := fieldTarget19.StartSet(__VDLTypeprofile2, len(m.Libraries))
			if err != nil {
				return err
			}
			for key22 := range m.Libraries {
				keyTarget21, err := setTarget20.StartKey()
				if err != nil {
					return err
				}

				if err := key22.FillVDLTarget(keyTarget21, __VDLType_profile_v_io_x_ref_services_profile_Library); err != nil {
					return err
				}
				if err := setTarget20.FinishKey(keyTarget21); err != nil {
					return err
				}
			}
			if err := fieldTarget19.FinishSet(setTarget20); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget18, fieldTarget19); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Specification) MakeVDLTarget() vdl.Target {
	return nil
}

func (m *Specification) IsZero() bool {

	var1 := true
	var2 := (m.Label == "")
	var1 = var1 && var2
	var3 := (m.Description == "")
	var1 = var1 && var3
	var4 := m.Arch.IsZero()
	var1 = var1 && var4
	var5 := m.Os.IsZero()
	var1 = var1 && var5
	var6 := m.Format.IsZero()
	var1 = var1 && var6
	var var7 bool
	if len(m.Libraries) == 0 {
		var7 = true
	}
	var1 = var1 && var7
	return var1
}

func init() {
	vdl.Register((*Library)(nil))
	vdl.Register((*Specification)(nil))
}

var __VDLTypeprofile0 *vdl.Type = vdl.TypeOf((*Library)(nil))
var __VDLTypeprofile1 *vdl.Type = vdl.TypeOf((*Specification)(nil))
var __VDLTypeprofile2 *vdl.Type = vdl.TypeOf(map[Library]struct{}(nil))
var __VDLType_profile_v_io_v23_services_build_Architecture *vdl.Type = vdl.TypeOf(build.ArchitectureAmd64)
var __VDLType_profile_v_io_v23_services_build_Format *vdl.Type = vdl.TypeOf(build.FormatElf)
var __VDLType_profile_v_io_v23_services_build_OperatingSystem *vdl.Type = vdl.TypeOf(build.OperatingSystemDarwin)
var __VDLType_profile_v_io_x_ref_services_profile_Library *vdl.Type = vdl.TypeOf(Library{})
var __VDLType_profile_v_io_x_ref_services_profile_Specification *vdl.Type = vdl.TypeOf(Specification{})

func __VDLEnsureNativeBuilt_profile() {
}

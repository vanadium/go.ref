// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"v.io/v23/naming"
	"v.io/x/lib/cmdline"
	"v.io/x/ref/lib/v23cmd"
	"v.io/x/ref/test"

	cmd_device "v.io/x/ref/services/device/device"
)

// TestUpdateCommand verifies the device update command.
func TestUpdateCommand(t *testing.T) {
	ctx, shutdown := test.InitForTest()
	defer shutdown()
	tapes := newTapeMap()
	server, endpoint, err := startServer(t, ctx, tapes)
	if err != nil {
		return
	}
	defer stopServer(t, server)

	cmd := cmd_device.CmdRoot
	appName := naming.JoinAddressName(endpoint.String(), "app")
	rootTape := tapes.forSuffix("")
	globName := naming.JoinAddressName(endpoint.String(), "glob")
	// TODO(caprita): Move joinLines to a common place.
	joinLines := func(args ...string) string {
		return strings.Join(args, "\n")
	}
	for _, c := range []struct {
		globResponses   []string
		statusResponses map[string][]interface{}
		expectedStimuli map[string][]interface{}
		expectedStdout  string
		expectedStderr  string
		expectedError   string
	}{
		{ // Everything succeeds.
			[]string{"app/2", "app/1", "app/3"},
			map[string][]interface{}{
				"app/1": []interface{}{instanceRunning, nil, nil, nil},
				"app/2": []interface{}{instanceNotRunning, nil},
				"app/3": []interface{}{installationActive, nil},
			},
			map[string][]interface{}{
				"app/1": []interface{}{"Status", KillStimulus{"Kill", 10 * time.Second}, "Update", "Run"},
				"app/2": []interface{}{"Status", "Update"},
				"app/3": []interface{}{"Status", "Update"},
			},
			joinLines(
				fmt.Sprintf("Successfully updated version for installation \"%s/3\".", appName),
				fmt.Sprintf("Successfully updated instance \"%s/1\".", appName),
				fmt.Sprintf("Successfully updated instance \"%s/2\".", appName)),
			"",
			"",
		},
		{ // Assorted failure modes.
			[]string{"app/1", "app/2", "app/3", "app/4", "app/5"},
			map[string][]interface{}{
				// Starts as running, fails Kill, but then
				// recovers. This ultimately counts as a success.
				"app/1": []interface{}{instanceRunning, fmt.Errorf("Simulate Kill failing"), instanceNotRunning, nil, nil},
				// Starts as running, fails Kill, and stays running.
				"app/2": []interface{}{instanceRunning, fmt.Errorf("Simulate Kill failing"), instanceRunning},
				// Starts as running, Kill and Update succeed, but Run fails.
				"app/3": []interface{}{instanceRunning, nil, nil, fmt.Errorf("Simulate Run failing")},
				// Starts as running, Kill succeeds, Update fails, but Run succeeds.
				"app/4": []interface{}{instanceRunning, nil, fmt.Errorf("Simulate Update failing"), nil},
				// Starts as running, Kill succeeds, Update fails, and Run fails.
				"app/5": []interface{}{instanceRunning, nil, fmt.Errorf("Simulate Update failing"), fmt.Errorf("Simulate Run failing")},
			},
			map[string][]interface{}{
				"app/1": []interface{}{"Status", KillStimulus{"Kill", 10 * time.Second}, "Status", "Update", "Run"},
				"app/2": []interface{}{"Status", KillStimulus{"Kill", 10 * time.Second}, "Status"},
				"app/3": []interface{}{"Status", KillStimulus{"Kill", 10 * time.Second}, "Update", "Run"},
				"app/4": []interface{}{"Status", KillStimulus{"Kill", 10 * time.Second}, "Update", "Run"},
				"app/5": []interface{}{"Status", KillStimulus{"Kill", 10 * time.Second}, "Update", "Run"},
			},
			joinLines(
				fmt.Sprintf("Successfully updated instance \"%s/1\".", appName),
				fmt.Sprintf("Successfully updated instance \"%s/3\".", appName),
			),
			joinLines(
				fmt.Sprintf("WARNING for \"%s/1\": recovered from Kill error (device.test:<rpc.Client>\"%s/1\".Kill: Error: Simulate Kill failing). Proceeding with update.", appName, appName),
				fmt.Sprintf("ERROR for \"%s/2\": Kill failed: device.test:<rpc.Client>\"%s/2\".Kill: Error: Simulate Kill failing.", appName, appName),
				fmt.Sprintf("ERROR for \"%s/3\": Run failed: device.test:<rpc.Client>\"%s/3\".Run: Error: Simulate Run failing.", appName, appName),
				fmt.Sprintf("ERROR for \"%s/4\": Update failed: device.test:<rpc.Client>\"%s/4\".Update: Error: Simulate Update failing.", appName, appName),
				fmt.Sprintf("ERROR for \"%s/5\": Run failed: device.test:<rpc.Client>\"%s/5\".Run: Error: Simulate Run failing.", appName, appName),
				fmt.Sprintf("ERROR for \"%s/5\": Update failed: device.test:<rpc.Client>\"%s/5\".Update: Error: Simulate Update failing.", appName, appName),
			),
			"encountered a total of 4 error(s)",
		},
	} {
		var stdout, stderr bytes.Buffer
		env := &cmdline.Env{Stdout: &stdout, Stderr: &stderr}
		tapes.rewind()
		rootTape.SetResponses(GlobResponse{c.globResponses})
		for n, r := range c.statusResponses {
			tapes.forSuffix(n).SetResponses(r...)
		}
		args := []string{"update", globName}
		if err := v23cmd.ParseAndRunForTest(cmd, ctx, env, args); err != nil {
			if want, got := c.expectedError, err.Error(); want != got {
				t.Errorf("Unexpected error: want %v, got %v", want, got)
			}
		} else {
			if c.expectedError != "" {
				t.Errorf("Expected to get error %v, but didn't get any error.", c.expectedError)
			}
		}

		if expected, got := c.expectedStdout, strings.TrimSpace(stdout.String()); got != expected {
			t.Errorf("Unexpected stdout output from update. Got %q, expected %q", got, expected)
		}
		if expected, got := c.expectedStderr, strings.TrimSpace(stderr.String()); got != expected {
			t.Errorf("Unexpected stderr output from update. Got %q, expected %q", got, expected)
		}
		for n, m := range c.expectedStimuli {
			if want, got := m, tapes.forSuffix(n).Play(); !reflect.DeepEqual(want, got) {
				t.Errorf("Unexpected stimuli for %v. Want: %v, got %v.", n, want, got)
			}
		}
		cmd_device.ResetGlobFlags()
	}
}
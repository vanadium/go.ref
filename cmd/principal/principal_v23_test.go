// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"v.io/x/ref"
	"v.io/x/ref/lib/v23test"
	"v.io/x/ref/test/expect"
)

func withCreds(dir string, cmd *v23test.Cmd) *v23test.Cmd {
	cmd.Vars[ref.EnvCredentials] = dir
	return cmd
}

// redirect redirects the stdout of the given command to the file at the given
// path.
func redirect(t *testing.T, cmd *v23test.Cmd, path string) {
	if err := ioutil.WriteFile(path, []byte(cmd.CombinedOutput()), 0600); err != nil {
		t.Fatalf("WriteFile(%q) failed: %v\n", path, err)
	}
}

// removePublicKeys replaces public keys (16 hex bytes, :-separated) with
// XX:....  This substitution enables comparison with golden output even when
// keys are freshly minted by the "principal create" command.
func removePublicKeys(input string) string {
	return regexp.MustCompile("([0-9a-f]{2}:){15}[0-9a-f]{2}").ReplaceAllString(input, "XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX")
}

func removeCaveats(input string) string {
	input = regexp.MustCompile(`0xa64c2d0119fba3348071feeb2f308000\(time\.Time=.*\)`).ReplaceAllString(input, "ExpiryCaveat")
	input = regexp.MustCompile(`0x54a676398137187ecdb26d2d69ba0003\(\[]string=.*\)`).ReplaceAllString(input, "MethodCaveat")
	input = regexp.MustCompile(`0x00000000000000000000000000000000\(bool=true\)`).ReplaceAllString(input, "Unconstrained")
	return input
}

func TestV23BlessSelf(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		outputDir         = sh.MakeTempDir()
		aliceDir          = filepath.Join(outputDir, "alice")
		aliceBlessingFile = filepath.Join(outputDir, "aliceself")
	)

	bin := sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
	sh.Cmd(bin, "create", aliceDir, "alice").Run()

	redirect(t, withCreds(aliceDir, sh.Cmd(bin, "blessself", "alicereborn")), aliceBlessingFile)
	got := removePublicKeys(withCreds(aliceDir, sh.Cmd(bin, "dumpblessings", aliceBlessingFile)).CombinedOutput())
	want := `Blessings          : alicereborn
PublicKey          : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Certificate chains : 1
Chain #0 (1 certificates). Root certificate public key: XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
  Certificate #0: alicereborn with 0 caveats
`
	if want != got {
		t.Fatalf("unexpected output, wanted \n%s, got\n%s", want, got)
	}
}

func TestV23Store(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		outputDir   = sh.MakeTempDir()
		bin         = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		aliceDir    = filepath.Join(outputDir, "alice")
		aliceFriend = filepath.Join(outputDir, "alice.bless")
		bobDir      = filepath.Join(outputDir, "bob")
		bobForPeer  = filepath.Join(outputDir, "bob.get.forpeer")
	)

	// Create two principals: alice and bob.
	sh.Cmd(bin, "create", aliceDir, "alice").Run()
	sh.Cmd(bin, "create", bobDir, "bob").Run()

	// Bless Bob with Alice's principal.
	redirect(t, withCreds(aliceDir, sh.Cmd(bin, "bless", "--for=1m", bobDir, "friend")), aliceFriend)

	// Run store forpeer on bob.
	sh.Cmd(bin, "--v23.credentials="+bobDir, "set", "forpeer", aliceFriend, "alice").Run()
	redirect(t, withCreds(aliceDir, sh.Cmd(bin, "--v23.credentials="+bobDir, "get", "forpeer", "alice:server")), bobForPeer)

	got := removeCaveats(removePublicKeys(sh.Cmd(bin, "dumpblessings", bobForPeer).CombinedOutput()))
	want := `Blessings          : bob,alice:friend
PublicKey          : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Certificate chains : 2
Chain #0 (1 certificates). Root certificate public key: XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
  Certificate #0: bob with 0 caveats
Chain #1 (2 certificates). Root certificate public key: XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
  Certificate #0: alice with 0 caveats
  Certificate #1: friend with 1 caveat
    (0) ExpiryCaveat
`
	if want != got {
		t.Errorf("unexpected output, got\n%s, wanted\n%s", got, want)
	}

	// Test the names flag.
	got = withCreds(aliceDir, sh.Cmd(bin, "--v23.credentials="+bobDir, "get", "forpeer", "--names", "alice:server")).CombinedOutput()
	want = `bob
alice:friend
`
	if got != want {
		t.Errorf("unexpected output, got %s, want %s", got, want)
	}

	// Test the rootkey flag. In particular alice:friend's rootkey should be equal to alice's publickey.
	got = withCreds(aliceDir, sh.Cmd(bin, "--v23.credentials="+bobDir, "get", "forpeer", "--rootkey", "alice:friend", "alice:server")).CombinedOutput()
	want = withCreds(aliceDir, sh.Cmd(bin, "get", "publickey", "--pretty")).CombinedOutput()
	if got != want {
		t.Errorf("unexpected output, got %s, want %s", got, want)
	}

	// Test the caveats flag.
	got = withCreds(aliceDir, sh.Cmd(bin, "--v23.credentials="+bobDir, "get", "forpeer", "--caveats", "alice:friend", "alice:server")).CombinedOutput()
	want = "Expires at"
	if !strings.HasPrefix(got, want) {
		t.Errorf("unexpected output, got %s, want %s", got, want)
	}
}

func TestV23Dump(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		outputDir       = sh.MakeTempDir()
		bin             = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		aliceDir        = filepath.Join(outputDir, "alice")
		aliceExpiredDir = filepath.Join(outputDir, "alice-expired")
	)

	sh.Cmd(bin, "create", aliceDir, "alice").Run()

	got := removePublicKeys(withCreds(aliceDir, sh.Cmd(bin, "dump")).CombinedOutput())
	want := `Public key : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Default Blessings : alice
---------------- BlessingStore ----------------
Default Blessings                alice
Peer pattern                     Blessings
...                              alice
---------------- BlessingRoots ----------------
Public key                                        Pattern
XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX   [alice]
`
	if want != got {
		t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
	}

	got = withCreds(aliceDir, sh.Cmd(bin, "dump", "-s")).CombinedOutput()
	want = "alice\n"
	if want != got {
		t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
	}

	sh.Cmd(bin, "--v23.credentials="+aliceDir, "fork", "--for", "-1h", aliceExpiredDir, "expired").Run()
	got = removePublicKeys(withCreds(aliceExpiredDir, sh.Cmd(bin, "dump")).CombinedOutput())
	want = `Public key : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Default Blessings : alice:expired [EXPIRED]
---------------- BlessingStore ----------------
Default Blessings                alice:expired
Peer pattern                     Blessings
...                              alice:expired
---------------- BlessingRoots ----------------
Public key                                        Pattern
XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX   [alice]
`
	if want != got {
		t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
	}

	got = withCreds(aliceExpiredDir, sh.Cmd(bin, "dump", "-s")).CombinedOutput()
	want = "alice:expired [EXPIRED]\n"
	if want != got {
		t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
	}
}

func TestV23GetRecognizedRoots(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		outputDir = sh.MakeTempDir()
		bin       = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		aliceDir  = filepath.Join(outputDir, "alice")
	)

	sh.Cmd(bin, "create", aliceDir, "alice").Run()

	got := removePublicKeys(withCreds(aliceDir, sh.Cmd(bin, "get", "recognizedroots")).CombinedOutput())
	want := `Public key                                        Pattern
XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX   [alice]
`
	if want != got {
		t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
	}
}

func TestV23GetPeermap(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		outputDir = sh.MakeTempDir()
		bin       = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		aliceDir  = filepath.Join(outputDir, "alice")
	)

	sh.Cmd(bin, "create", aliceDir, "alice").Run()

	got := withCreds(aliceDir, sh.Cmd(bin, "get", "peermap")).CombinedOutput()
	want := `Default Blessings                alice
Peer pattern                     Blessings
...                              alice
`
	if want != got {
		t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
	}
}

// Given an invocation of "principal recvblessings", this function returns the
// arguments to provide to "principal bless" provided by the "recvblessings"
// invocation.
//
// For example,
// principal recvblessings
// would typically print something like:
//    principal bless --remote-key=<some_public_key> --remote-token=<some_token> extensionfoo
// as an example of command line to use to send the blessings over.
//
// In that case, this method would return:
// { "--remote-key=<some_public_key>", "--remote-token=<some_token>", "extensionfoo"}
func blessArgsFromRecvBlessings(s *expect.Session) []string {
	cmd := s.ExpectSetEventuallyRE("(^principal bless .*$)")[0][0]
	return strings.Split(cmd, " ")[2:]
}

func TestV23RecvBlessings(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		outputDir    = sh.MakeTempDir()
		bin          = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		aliceDir     = filepath.Join(outputDir, "alice")
		bobDir       = filepath.Join(outputDir, "bob")
		carolDir     = filepath.Join(outputDir, "carol")
		bobBlessFile = filepath.Join(outputDir, "bobBlessInfo")
	)

	// Generate principals
	sh.Cmd(bin, "create", aliceDir, "alice").Run()
	sh.Cmd(bin, "create", bobDir, "bob").Run()
	sh.Cmd(bin, "create", carolDir, "carol").Run()

	// Run recvblessings on carol, and have alice send blessings over
	// (blessings received must be set as default and shareable with all peers).
	var args []string
	{
		cmd := sh.Cmd(bin, "--v23.credentials="+carolDir, "--v23.tcp.address=127.0.0.1:0", "recvblessings")
		session := expect.NewSession(t, cmd.StdoutPipe(), time.Minute)
		cmd.Start()
		args = append([]string{"bless", "--require-caveats=false"}, blessArgsFromRecvBlessings(session)...)
		// Use the "friend:carol" extension
		args = append(args, "friend:carol")
	}
	withCreds(aliceDir, sh.Cmd(bin, args...)).Run()

	// Run recvblessings on carol, and have alice send blessings over
	// (blessings received must be set as shareable with peers matching 'alice:...'.)
	{
		cmd := sh.Cmd(bin, "--v23.credentials="+carolDir, "--v23.tcp.address=127.0.0.1:0", "recvblessings", "--for-peer=alice", "--set-default=false")
		session := expect.NewSession(t, cmd.StdoutPipe(), time.Minute)
		cmd.Start()
		// recvblessings suggests a random extension, find the extension and replace it with friend:carol:foralice.
		args = append([]string{"bless", "--require-caveats=false"}, blessArgsFromRecvBlessings(session)...)
		args = append(args, "friend:carol:foralice")
	}
	withCreds(aliceDir, sh.Cmd(bin, args...)).Run()

	// Run recvblessings on carol with the --remote-arg-file flag, and have bob send blessings over with the --remote-arg-file flag.
	{
		cmd := sh.Cmd(bin, "--v23.credentials="+carolDir, "--v23.tcp.address=127.0.0.1:0", "recvblessings", "--for-peer=bob", "--set-default=false", "--remote-arg-file="+bobBlessFile)
		session := expect.NewSession(t, cmd.StdoutPipe(), time.Minute)
		cmd.Start()
		// recvblessings suggests a random extension, use friend:carol:forbob instead.
		args = append([]string{"bless", "--require-caveats=false"}, blessArgsFromRecvBlessings(session)...)
		args = append(args, "friend:carol:forbob")
	}
	withCreds(bobDir, sh.Cmd(bin, args...)).Run()

	cmd := sh.Cmd(bin, "--v23.credentials="+carolDir, "--v23.tcp.address=127.0.0.1:0", "recvblessings", "--for-peer=alice:...", "--set-default=false", "--vmodule=*=2", "--logtostderr")
	session := expect.NewSession(t, cmd.StdoutPipe(), time.Minute)
	cmd.Start()
	args = append([]string{"bless", "--require-caveats=false"}, blessArgsFromRecvBlessings(session)...)
	args = append(args, "willfail")

	{
		// Mucking around with remote-key should fail.
		cpy := strings.Split(regexp.MustCompile("remote-key=").ReplaceAllString(strings.Join(args, " "), "remote-key=BAD"), " ")
		cmd := withCreds(aliceDir, sh.Cmd(bin, cpy...))
		cmd.ExitErrorIsOk = true
		_, stderr := cmd.Output()
		if cmd.Err == nil {
			t.Fatalf("%v should have failed, but did not", cpy)
		}

		if want, got := "key mismatch", stderr; !strings.Contains(got, want) {
			t.Fatalf("expected %q to be contained within\n%s\n, but was not", want, got)
		}
	}

	{
		// Mucking around with the token should fail.
		cpy := strings.Split(regexp.MustCompile("remote-token=").ReplaceAllString(strings.Join(args, " "), "remote-token=BAD"), " ")
		cmd := withCreds(aliceDir, sh.Cmd(bin, cpy...))
		cmd.ExitErrorIsOk = true
		_, stderr := cmd.Output()
		if cmd.Err == nil {
			t.Fatalf("%v should have failed, but did not", cpy)
		}

		if want, got := "blessings received from unexpected sender", stderr; !strings.Contains(got, want) {
			t.Fatalf("expected %q to be contained within\n%s\n, but was not", want, got)
		}
	}

	// Dump carol out, the only blessing that survives should be from the
	// first "bless" command. (alice:friend:carol).
	got := removePublicKeys(sh.Cmd(bin, "--v23.credentials="+carolDir, "dump").CombinedOutput())
	want := `Public key : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Default Blessings : alice:friend:carol
---------------- BlessingStore ----------------
Default Blessings                alice:friend:carol
Peer pattern                     Blessings
...                              alice:friend:carol
alice                            alice:friend:carol:foralice
bob                              bob:friend:carol:forbob
---------------- BlessingRoots ----------------
Public key                                        Pattern
XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX   [alice]
XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX   [bob]
XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX   [carol]
`
	if want != got {
		t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
	}
}

func TestV23RecvBlessingsInteractive(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		outputDir = sh.MakeTempDir()
		bin       = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		aliceDir  = filepath.Join(outputDir, "alice")
		bobDir    = filepath.Join(outputDir, "bob")
	)

	// Generate principals
	sh.Cmd(bin, "create", aliceDir, "alice").Run()
	sh.Cmd(bin, "create", bobDir, "bob").Run()

	// Run recvblessings on bob
	cmd := sh.Cmd(bin, "--v23.credentials="+bobDir, "--v23.tcp.address=127.0.0.1:0", "recvblessings")
	session := expect.NewSession(t, cmd.StdoutPipe(), time.Minute)
	cmd.Start()
	args := blessArgsFromRecvBlessings(session)

	// When running the exact command, must be prompted about caveats.
	{
		cmd := withCreds(aliceDir, sh.Cmd(bin, append([]string{"bless"}, args...)...))
		session := expect.NewSession(t, cmd.StdoutPipe(), time.Minute)
		cmd.Stdin = bytes.NewBufferString("yeah\n")
		cmd.ExitErrorIsOk = true
		cmd.Start()
		session.Expect("WARNING: No caveats provided")
		// Saying something other than "yes" or "YES"
		// should fail.
		if cmd.Wait(); cmd.Err == nil {
			t.Fatalf("Expected principal bless to fail because the wrong input was provided")
		}
	}
	// When agreeing to have no caveats, must specify an extension
	{
		cmd := withCreds(aliceDir, sh.Cmd(bin, append([]string{"bless"}, args...)...))
		session := expect.NewSession(t, cmd.StdoutPipe(), time.Minute)
		cmd.Stdin = bytes.NewBufferString("yes\n")
		cmd.ExitErrorIsOk = true
		cmd.Start()
		session.Expect("WARNING: No caveats provided")
		if cmd.Wait(); cmd.Err == nil {
			t.Fatalf("Expected principal bless to fail because the wrong input was provided")
		}
	}
	// When providing both, the bless command should succeed.
	{
		cmd := withCreds(aliceDir, sh.Cmd(bin, append([]string{"bless"}, args...)...))
		cmd.Stdin = bytes.NewBufferString("YES\nfriend:bobby\n")
		cmd.Run()
	}
	got := removePublicKeys(sh.Cmd(bin, "--v23.credentials="+bobDir, "dump").CombinedOutput())
	want := `Public key : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Default Blessings : alice:friend:bobby
---------------- BlessingStore ----------------
Default Blessings                alice:friend:bobby
Peer pattern                     Blessings
...                              alice:friend:bobby
---------------- BlessingRoots ----------------
Public key                                        Pattern
XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX   [alice]
XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX   [bob]
`
	if want != got {
		t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
	}
}

func TestV23Fork(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		outputDir             = sh.MakeTempDir()
		bin                   = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		aliceDir              = filepath.Join(outputDir, "alice")
		alicePhoneDir         = filepath.Join(outputDir, "alice-phone")
		alicePhoneCalendarDir = filepath.Join(outputDir, "alice-phone-calendar")
		tmpfile               = filepath.Join(outputDir, "tmpfile")
	)

	// Generate principals for alice.
	sh.Cmd(bin, "create", aliceDir, "alice").Run()

	// Run fork to setup up credentials for alice:phone that are
	// blessed by alice under the extension "phone".
	sh.Cmd(bin, "--v23.credentials="+aliceDir, "fork", "--for", "1h", alicePhoneDir, "phone").Run()

	// Dump alice-phone out, the only blessings it has must be from alice (alice:phone).
	{
		got := removePublicKeys(sh.Cmd(bin, "--v23.credentials="+alicePhoneDir, "dump").CombinedOutput())
		want := `Public key : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Default Blessings : alice:phone
---------------- BlessingStore ----------------
Default Blessings                alice:phone
Peer pattern                     Blessings
...                              alice:phone
---------------- BlessingRoots ----------------
Public key                                        Pattern
XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX   [alice]
`
		if want != got {
			t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
		}
	}
	// And it should have an expiry caveat
	{
		redirect(t, sh.Cmd(bin, "--v23.credentials", alicePhoneDir, "get", "default"), tmpfile)
		got := removeCaveats(removePublicKeys(sh.Cmd(bin, "dumpblessings", tmpfile).CombinedOutput()))
		want := `Blessings          : alice:phone
PublicKey          : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Certificate chains : 1
Chain #0 (2 certificates). Root certificate public key: XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
  Certificate #0: alice with 0 caveats
  Certificate #1: phone with 1 caveat
    (0) ExpiryCaveat
`
		if want != got {
			t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
		}
	}

	// Run fork to setup up credentials for alice:phone:calendar that are
	// blessed by alice:phone under the extension "calendar".
	sh.Cmd(bin, "--v23.credentials="+alicePhoneDir, "fork", "--for", "1h", alicePhoneCalendarDir, "calendar").Run()
	{
		got := removePublicKeys(sh.Cmd(bin, "--v23.credentials="+alicePhoneCalendarDir, "dump").CombinedOutput())
		want := `Public key : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Default Blessings : alice:phone:calendar
---------------- BlessingStore ----------------
Default Blessings                alice:phone:calendar
Peer pattern                     Blessings
...                              alice:phone:calendar
---------------- BlessingRoots ----------------
Public key                                        Pattern
XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX   [alice]
`
		if want != got {
			t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
		}
	}
	{
		redirect(t, sh.Cmd(bin, "--v23.credentials", alicePhoneCalendarDir, "get", "default"), tmpfile)
		got := removeCaveats(removePublicKeys(sh.Cmd(bin, "dumpblessings", tmpfile).CombinedOutput()))
		want := `Blessings          : alice:phone:calendar
PublicKey          : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Certificate chains : 1
Chain #0 (3 certificates). Root certificate public key: XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
  Certificate #0: alice with 0 caveats
  Certificate #1: phone with 1 caveat
    (0) ExpiryCaveat
  Certificate #2: calendar with 1 caveat
    (0) ExpiryCaveat
`
		if want != got {
			t.Fatalf("unexpected output, got\n%s, wanted\n%s", got, want)
		}
	}
}

func TestV23Create(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		outputDir = sh.MakeTempDir()
		bin       = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		aliceDir  = filepath.Join(outputDir, "alice")
	)

	// Creating a principal should succeed the first time.
	sh.Cmd(bin, "create", aliceDir, "alice").Run()

	// The second time should fail (the create command won't override an existing principal).
	cmd := sh.Cmd(bin, "create", aliceDir, "alice")
	cmd.ExitErrorIsOk = true
	if cmd.Run(); cmd.Err == nil {
		t.Fatalf("principal creation should have failed, but did not")
	}

	// If we specify -overwrite, it will.
	sh.Cmd(bin, "create", "--overwrite", aliceDir, "alice").Run()
}

func TestV23Caveats(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		outputDir         = sh.MakeTempDir()
		aliceDir          = filepath.Join(outputDir, "alice")
		aliceBlessingFile = filepath.Join(outputDir, "aliceself")
	)

	bin := sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
	sh.Cmd(bin, "create", aliceDir, "alice").Run()

	args := []string{
		"blessself",
		"--caveat=\"v.io/v23/security\".MethodCaveat={\"method\"}",
		"--caveat={{0x54,0xa6,0x76,0x39,0x81,0x37,0x18,0x7e,0xcd,0xb2,0x6d,0x2d,0x69,0xba,0x0,0x3},typeobject([]string)}={\"method\"}",
		"alicereborn",
	}
	redirect(t, withCreds(aliceDir, sh.Cmd(bin, args...)), aliceBlessingFile)
	got := removeCaveats(removePublicKeys(withCreds(aliceDir, sh.Cmd(bin, "dumpblessings", aliceBlessingFile)).CombinedOutput()))
	want := `Blessings          : alicereborn
PublicKey          : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Certificate chains : 1
Chain #0 (1 certificates). Root certificate public key: XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
  Certificate #0: alicereborn with 2 caveats
    (0) MethodCaveat
    (1) MethodCaveat
`
	if want != got {
		t.Fatalf("unexpected output, wanted \n%s, got\n%s", want, got)
	}
}

func TestV23ForkWithoutVDLPATH(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	delete(sh.Vars, "JIRI_ROOT")
	delete(sh.Vars, "VDLPATH")

	var (
		parent = sh.MakeTempDir()
		bin    = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
	)

	sh.Cmd(bin, "create", parent, "parent").Run()
	sh.Cmd(bin, "--v23.credentials="+parent, "fork", "--for=1s", sh.MakeTempDir(), "child").Run()
}

func TestV23ForkWithoutCaveats(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		parent = sh.MakeTempDir()
		child  = sh.MakeTempDir()
		bin    = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
	)

	sh.Cmd(bin, "create", parent, "parent").Run()

	cmd := sh.Cmd(bin, "--v23.credentials", parent, "fork", child, "child")
	cmd.ExitErrorIsOk = true
	if _, stderr := cmd.Output(); cmd.Err == nil {
		t.Errorf("fork should have failed without any caveats, but did not")
	} else if got, want := stderr, "ERROR: no caveats provided"; !strings.Contains(got, want) {
		t.Errorf("fork returned error: %q, expected error to contain %q", got, want)
	}

	cmd = sh.Cmd(bin, "--v23.credentials", parent, "fork", "--for=0", child, "child")
	cmd.ExitErrorIsOk = true
	if _, stderr := cmd.Output(); cmd.Err == nil {
		t.Errorf("fork should have failed without any caveats, but did not")
	} else if got, want := stderr, "ERROR: no caveats provided"; !strings.Contains(got, want) {
		t.Errorf("fork returned error: %q, expected error to contain %q", got, want)
	}

	sh.Cmd(bin, "--v23.credentials", parent, "fork", "--require-caveats=false", child, "child").Run()
}

func TestV23Bless(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		bin      = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		dir      = sh.MakeTempDir()
		aliceDir = filepath.Join(dir, "alice")
		bobDir   = filepath.Join(dir, "bob")
		tmpfile  = filepath.Join(dir, "tmpfile")
	)
	// Create two principals: alice and bob
	sh.Cmd(bin, "create", aliceDir, "alice").Run()
	sh.Cmd(bin, "create", bobDir, "bob").Run()

	// All blessings will be done by "alice"
	aliceCmd := func(name string, args ...string) *v23test.Cmd {
		return withCreds(aliceDir, sh.Cmd(name, args...))
	}

	{
		// "alice" should fail to bless "bob" without any caveats
		cmd := aliceCmd(bin, "bless", bobDir, "friend")
		cmd.ExitErrorIsOk = true
		if _, stderr := cmd.Output(); cmd.Err == nil {
			t.Errorf("bless should have failed when no caveats are specified")
		} else if got, want := stderr, "ERROR: no caveats provided"; !strings.Contains(got, want) {
			t.Errorf("got error %q, expected to match %q", got, want)
		}
	}
	{
		// But succeed if --require-caveats=false is specified
		redirect(t, aliceCmd(bin, "bless", "--require-caveats=false", bobDir, "friend"), tmpfile)
		got := removeCaveats(removePublicKeys(aliceCmd(bin, "dumpblessings", tmpfile).CombinedOutput()))
		want := `Blessings          : alice:friend
PublicKey          : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Certificate chains : 1
Chain #0 (2 certificates). Root certificate public key: XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
  Certificate #0: alice with 0 caveats
  Certificate #1: friend with 1 caveat
    (0) Unconstrained
`
		if got != want {
			t.Errorf("Got\n%vWant\n%v", got, want)
		}
	}
	{
		// And succeed if --for is specified
		redirect(t, aliceCmd(bin, "bless", "--for=1m", bobDir, "friend"), tmpfile)
		got := removeCaveats(removePublicKeys(aliceCmd(bin, "dumpblessings", tmpfile).CombinedOutput()))
		want := `Blessings          : alice:friend
PublicKey          : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Certificate chains : 1
Chain #0 (2 certificates). Root certificate public key: XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
  Certificate #0: alice with 0 caveats
  Certificate #1: friend with 1 caveat
    (0) ExpiryCaveat
`
		if got != want {
			t.Errorf("Got\n%vWant\n%v", got, want)
		}
	}
	{
		// If the Blessings are expired, dumpBlessings should print so.
		redirect(t, aliceCmd(bin, "bless", "--for=-1s", bobDir, "friend"), tmpfile)
		got := removeCaveats(removePublicKeys(aliceCmd(bin, "dumpblessings", tmpfile).CombinedOutput()))
		want := `Blessings          : alice:friend [EXPIRED]
PublicKey          : XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
Certificate chains : 1
Chain #0 (2 certificates). Root certificate public key: XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
  Certificate #0: alice with 0 caveats
  Certificate #1: friend with 1 caveat
    (0) ExpiryCaveat
`
		if got != want {
			t.Errorf("Got\n%vWant\n%v", got, want)
		}
	}
	{
		// But not if --for=0
		cmd := aliceCmd(bin, "bless", "--for=0", bobDir, "friend")
		cmd.ExitErrorIsOk = true
		if _, stderr := cmd.Output(); cmd.Err == nil {
			t.Errorf("bless should have failed when no caveats are specified")
		} else if got, want := stderr, "ERROR: no caveats provided"; !strings.Contains(got, want) {
			t.Errorf("got error %q, expected to match %q", got, want)
		}
	}
}

func TestV23AddBlessingsToRoots(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		bin          = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		aliceDir     = sh.MakeTempDir()
		bobDir       = sh.MakeTempDir()
		blessingFile = filepath.Join(sh.MakeTempDir(), "bobfile")

		// Extract the public key from the first line of output from
		// "principal dump", which is formatted as:
		// Public key : <the public key>
		publicKey = func(dir string) string {
			output := sh.Cmd(bin, "--v23.credentials="+dir, "dump").CombinedOutput()
			line := strings.SplitN(output, "\n", 2)[0]
			fields := strings.Split(line, " ")
			return fields[len(fields)-1]
		}
	)
	// Create two principals, "alice" and "bob"
	sh.Cmd(bin, "create", aliceDir, "alice").Run()
	sh.Cmd(bin, "create", bobDir, "bob").Run()
	// Have bob create a "bob/friend" blessing and have alice recognize that.
	redirect(t, sh.Cmd(bin, "--v23.credentials="+bobDir, "bless", "--require-caveats=false", aliceDir, "friend"), blessingFile)
	sh.Cmd(bin, "--v23.credentials="+aliceDir, "recognize", blessingFile).Run()

	want := fmt.Sprintf(`Public key                                        Pattern
%v   [alice]
%v   [bob]
`, publicKey(aliceDir), publicKey(bobDir))

	// Finally view alice's recognized roots, it should have lines corresponding to aliceLine and bobLine.
	got := sh.Cmd(bin, "--v23.credentials="+aliceDir, "get", "recognizedroots").CombinedOutput()
	if got != want {
		t.Fatalf("Got:\n%v\n\nWant:\n%v", got, want)
	}
}

func TestV23AddKeyToRoots(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		bin       = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		outputDir = sh.MakeTempDir()
		aliceDir  = filepath.Join(outputDir, "alice")
		bobDir    = filepath.Join(outputDir, "bob")
	)
	sh.Cmd(bin, "create", aliceDir, "alice").Run()
	sh.Cmd(bin, "create", bobDir, "bob").Run()
	// Get bob's public key and add it to roots for alice
	bobKey := strings.TrimSpace(sh.Cmd(bin, "--v23.credentials="+bobDir, "get", "publickey").CombinedOutput())
	bobPrettyKey := strings.TrimSpace(sh.Cmd(bin, "--v23.credentials="+bobDir, "get", "publickey", "--pretty").CombinedOutput())
	sh.Cmd(bin, "--v23.credentials="+aliceDir, "recognize", "bob", bobKey).Run()

	// Verify that it has been added
	output := sh.Cmd(bin, "--v23.credentials="+aliceDir, "dump").CombinedOutput()
	want := fmt.Sprintf("%v   [bob]", bobPrettyKey)
	for _, line := range strings.Split(output, "\n") {
		if line == want {
			return
		}
	}
	t.Errorf("Could not find line:\n%v\nin output:\n%v\n", want, output)
}

func TestV23DumpRoots(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()

	var (
		bin             = sh.JiriBuildGoPkg("v.io/x/ref/cmd/principal")
		outputDir       = sh.MakeTempDir()
		aliceDir        = filepath.Join(outputDir, "alice")
		bobDir          = filepath.Join(outputDir, "bob")
		aliceFriend     = filepath.Join(outputDir, "alicefriend")
		aliceFriendRoot = filepath.Join(outputDir, "alicefriendroot")
		aliceDefault    = filepath.Join(outputDir, "alicedefault")
	)
	sh.Cmd(bin, "create", aliceDir, "alice").Run()
	sh.Cmd(bin, "create", bobDir, "bob").Run()
	redirect(t, sh.Cmd(bin, "--v23.credentials="+aliceDir, "bless", "--require-caveats=false", bobDir, "friend"), aliceFriend)
	redirect(t, sh.Cmd(bin, "dumproots", aliceFriend), aliceFriendRoot)
	redirect(t, sh.Cmd(bin, "--v23.credentials="+aliceDir, "get", "default"), aliceDefault)

	want := sh.Cmd(bin, "dumpblessings", aliceDefault).CombinedOutput()
	got := sh.Cmd(bin, "dumpblessings", aliceFriendRoot).CombinedOutput()
	if got != want {
		t.Errorf("Got:\n%s\nWant:\n%s\n", got, want)
	}
}

func TestMain(m *testing.M) {
	os.Exit(v23test.Run(m.Run))
}

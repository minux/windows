// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package attachconsole is meant to be used as a side-effect only
// import. It attaches the calling process to its parent's console.
// It is most useful for an application built with
//	-ldflags "-H windowsgui"
// flags so that it won't have a console of itself. After importing
// this package, if the application is run from the command prompt
// window, the console output will appear in the command prompt window
// instead of discarded.
package attachconsole // import "github.com/minux/windows/attachconsole"

import "syscall"

func attach() error {
	const ATTACH_PARENT_PROCESS = ^uintptr(0)
	proc := syscall.MustLoadDLL("kernel32.dll").MustFindProc("AttachConsole")
	_, _, err := proc.Call(ATTACH_PARENT_PROCESS)
	return err
}

func init() {
	attach()
}

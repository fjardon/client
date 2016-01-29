// Copyright 2016 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

// +build fuse

package test

import (
	"testing"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"bazil.org/fuse/fs/fstestutil"
	"github.com/keybase/client/go/logger"
	"github.com/keybase/kbfs/libfuse"
	"github.com/keybase/kbfs/libkbfs"
	"golang.org/x/net/context"
)

type fuseEngine struct {
	fsEngine
}

func createEngine() Engine { return &fuseEngine{} }

// Name returns the name of the Engine.
func (*fuseEngine) Name() string {
	return "fuse"
}

func (e *fuseEngine) Init() {
	e.createUser = createUserFuse
}

func createUserFuse(t *testing.T, ith int, config *libkbfs.ConfigLocal, tlf string) User {
	log := logger.NewTestLogger(t)
	fuse.Debug = func(msg interface{}) {
		log.Debug("%s", msg)
	}

	filesys := libfuse.NewFS(config, nil, false)
	ctx := context.Background()
	ctx = context.WithValue(ctx, libfuse.CtxAppIDKey, filesys)
	ctx, cancelFn := context.WithCancel(ctx)
	fn := func(mnt *fstestutil.Mount) fs.FS {
		filesys.SetFuseConn(mnt.Server, mnt.Conn)
		return filesys
	}
	mnt, err := fstestutil.MountedFuncT(t, fn, &fs.Config{
		GetContext: func() context.Context {
			return ctx
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	filesys.LaunchNotificationProcessor(ctx)
	return &fsUser{
		mntDir: mnt.Dir,
		config: config,
		cancel: cancelFn,
		close:  mnt.Close,
		tlf:    tlf,
		notificationGroupWait: filesys.NotificationGroupWait,
	}
}

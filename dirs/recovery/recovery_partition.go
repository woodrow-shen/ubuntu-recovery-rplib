package recovery

import (
	"path/filepath"
)

// the various file paths of recovery partition
var (
	KernelSnap string
	GadgetSnap string
	OsSnap     string

	RecoveryMaterialDir          string
	WritableLocalIncludeSquashfs string
)

func init() {
	// init the global directories at startup
	SetRootDir("/")
}

func SetRootDir(rootdir string) {
	KernelSnap = filepath.Join(rootdir, "kernel.snap")
	GadgetSnap = filepath.Join(rootdir, "gadget.snap")
	OsSnap = filepath.Join("os.snap")

	RecoveryMaterialDir = filepath.Join(rootdir, "recovery")
	WritableLocalIncludeSquashfs = filepath.Join(RecoveryMaterialDir, "writable_local-include.squashfs")
}
package all

import (
	_ "github.com/rclone/rclone/cmd"

	// Required for FUSE mount
	_ "github.com/rclone/rclone/cmd/mount"

	// Required for password obfuscation
	_ "github.com/rclone/rclone/cmd/obscure"

	// Diagnostics
	_ "github.com/rclone/rclone/cmd/version"
	_ "github.com/rclone/rclone/cmd/listremotes"
	_ "github.com/rclone/rclone/cmd/ls"
	_ "github.com/rclone/rclone/cmd/lsd"
	_ "github.com/rclone/rclone/cmd/lsf"
	_ "github.com/rclone/rclone/cmd/lsjson"
	_ "github.com/rclone/rclone/cmd/about"

	// Useful recovery transfer commands
	_ "github.com/rclone/rclone/cmd/copy"
	_ "github.com/rclone/rclone/cmd/copyto"
	_ "github.com/rclone/rclone/cmd/sync"
	_ "github.com/rclone/rclone/cmd/check"
	_ "github.com/rclone/rclone/cmd/mkdir"
	_ "github.com/rclone/rclone/cmd/delete"
	_ "github.com/rclone/rclone/cmd/deletefile"
	_ "github.com/rclone/rclone/cmd/purge"

	// Useful for testing/config from recovery shell
	_ "github.com/rclone/rclone/cmd/config"
)

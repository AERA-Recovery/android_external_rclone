package all

import (
	// Core / local
	_ "github.com/rclone/rclone/backend/local"
	_ "github.com/rclone/rclone/backend/memory"

	// Standard NAS / network protocols
	_ "github.com/rclone/rclone/backend/sftp"
	_ "github.com/rclone/rclone/backend/smb"
	_ "github.com/rclone/rclone/backend/webdav"
	_ "github.com/rclone/rclone/backend/ftp"
	_ "github.com/rclone/rclone/backend/http"

	// Useful wrapper / composition backends
	_ "github.com/rclone/rclone/backend/alias"
	_ "github.com/rclone/rclone/backend/crypt"
	_ "github.com/rclone/rclone/backend/union"
	_ "github.com/rclone/rclone/backend/combine"
	_ "github.com/rclone/rclone/backend/chunker"
	_ "github.com/rclone/rclone/backend/compress"
	_ "github.com/rclone/rclone/backend/hasher"
)

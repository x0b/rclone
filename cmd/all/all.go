// Package all imports all the commands
package all

import (
	// Active commands
	_ "github.com/rclone/rclone/cmd"
	_ "github.com/rclone/rclone/cmd/about"
	_ "github.com/rclone/rclone/cmd/authorize"
	_ "github.com/rclone/rclone/cmd/backend"
	//DISABLED:_ "github.com/rclone/rclone/cmd/cachestats"
	_ "github.com/rclone/rclone/cmd/cat"
	_ "github.com/rclone/rclone/cmd/check"
	_ "github.com/rclone/rclone/cmd/cleanup"
	//DISABLED:_ "github.com/rclone/rclone/cmd/cmount"
	_ "github.com/rclone/rclone/cmd/config"
	_ "github.com/rclone/rclone/cmd/copy"
	_ "github.com/rclone/rclone/cmd/copyto"
	_ "github.com/rclone/rclone/cmd/copyurl"
	//DISABLED:_ "github.com/rclone/rclone/cmd/cryptcheck"
	//DISABLED:_ "github.com/rclone/rclone/cmd/cryptdecode"
	//DISABLED:_ "github.com/rclone/rclone/cmd/dbhashsum"
	//DISABLED:_ "github.com/rclone/rclone/cmd/dedupe"
	_ "github.com/rclone/rclone/cmd/delete"
	_ "github.com/rclone/rclone/cmd/deletefile"
	//DISABLED:_ "github.com/rclone/rclone/cmd/genautocomplete"
	//DISABLED:_ "github.com/rclone/rclone/cmd/gendocs"
	_ "github.com/rclone/rclone/cmd/hashsum"
	//DISABLED:_ "github.com/rclone/rclone/cmd/info"
	_ "github.com/rclone/rclone/cmd/link"
	_ "github.com/rclone/rclone/cmd/listremotes"
	//DISABLED:_ "github.com/rclone/rclone/cmd/ls"
	//DISABLED:_ "github.com/rclone/rclone/cmd/lsd"
	//DISABLED:_ "github.com/rclone/rclone/cmd/lsf"
	_ "github.com/rclone/rclone/cmd/lsjson"
	//DISABLED:_ "github.com/rclone/rclone/cmd/lsl"
	_ "github.com/rclone/rclone/cmd/md5sum"
	//DISABLED:_ "github.com/rclone/rclone/cmd/memtest"
	_ "github.com/rclone/rclone/cmd/mkdir"
	//DISABLED:_ "github.com/rclone/rclone/cmd/mount"
	//DISABLED:_ "github.com/rclone/rclone/cmd/mount2"
	_ "github.com/rclone/rclone/cmd/move"
	_ "github.com/rclone/rclone/cmd/moveto"
	_ "github.com/rclone/rclone/cmd/ncdu"
	_ "github.com/rclone/rclone/cmd/obscure"
	_ "github.com/rclone/rclone/cmd/purge"
	_ "github.com/rclone/rclone/cmd/rc"
	_ "github.com/rclone/rclone/cmd/rcat"
	_ "github.com/rclone/rclone/cmd/rcd"
	_ "github.com/rclone/rclone/cmd/reveal"
	_ "github.com/rclone/rclone/cmd/rmdir"
	//DISABLED:	_ "github.com/rclone/rclone/cmd/rmdirs"
	_ "github.com/rclone/rclone/cmd/serve"
	//DISABLED:	_ "github.com/rclone/rclone/cmd/settier"
	_ "github.com/rclone/rclone/cmd/sha1sum"
	_ "github.com/rclone/rclone/cmd/size"
	_ "github.com/rclone/rclone/cmd/sync"
	_ "github.com/rclone/rclone/cmd/touch"
	//DISABLED:	_ "github.com/rclone/rclone/cmd/tree"
	_ "github.com/rclone/rclone/cmd/version"
)

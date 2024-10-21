package main

import (
	_ "github.com/Jelingam/Warp/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}

package main

import (
	"github.com/code-serenade/web-server-go/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":17111")
}

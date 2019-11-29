package main

import (
	"github.com/lwc/protoc-gen-benum/benum"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		benum.Benum(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}

package main

import (
	"os"

	"knative.dev/serving/pkg/queue/sharedmain"

	"github.com/diarmuidie/qpoptions-example/pkg/foo"
)

func main() {
	fooQPO := foo.NewFooQPOption()
	defer fooQPO.Shutdown()

	if sharedmain.Main(fooQPO.Setup) != nil {
		fooQPO.Shutdown()
		os.Exit(1)
	}
}

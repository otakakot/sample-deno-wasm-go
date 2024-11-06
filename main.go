//go:generate sh -c "GOOS=js GOARCH=wasm go build -o main.wasm main.go && cat main.wasm | deno run https://denopkg.com/syumai/binpack/mod.ts > mainwasm.ts"
package main

import (
	"fmt"
	"log/slog"
	"syscall/js"
)

var close = make(chan struct{})

func init() {
	js.Global().Set("handle", js.FuncOf(handle))
}

func handle(this js.Value, args []js.Value) any {
	defer func() {
		slog.Info("handle done")
		close <- struct{}{}
	}()

	println("Hello, Go Wasm World! from handle")

	fmt.Printf("this: %v\n", this)
	fmt.Printf("args: %v\n", args)

	return nil
}

func main() {
	<-close
}

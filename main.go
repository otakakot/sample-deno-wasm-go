//go:generate sh -c "GOOS=js GOARCH=wasm go build -o main.wasm main.go && cat main.wasm | deno run https://denopkg.com/syumai/binpack/mod.ts > mainwasm.ts"
package main

import (
	"fmt"
	"log/slog"
	"syscall/js"
)

var done = make(chan struct{})

func init() {
	js.Global().Set("hello", js.FuncOf(hello))
}

func hello(this js.Value, args []js.Value) any {
	defer func() {
		slog.Info("hello done")
		done <- struct{}{}
	}()

	fmt.Printf("this: %v\n", this)
	fmt.Printf("args: %v\n", args)

	println("Hello, World from Wasm Go!")

	return nil
}

func main() {
	<-done
}

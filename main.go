//go:generate sh -c "GOOS=js GOARCH=wasm go build -o main.wasm main.go && cat main.wasm | deno run https://denopkg.com/syumai/binpack/mod.ts > mainwasm.ts"
package main

import (
	"fmt"
	"syscall/js"
)

var done = make(chan struct{})

func init() {
	js.Global().Set("hello", js.FuncOf(hello))
}

func hello(this js.Value, args []js.Value) any {
	defer func() {
		fmt.Printf("hello done\n")
		done <- struct{}{}
	}()

	fmt.Printf("this: %v\n", this)

	fmt.Printf("args: %v\n", args)

	fmt.Printf("Hello World from Wasm Go!\n")

	js.Global().Get("console").Call("info", "Hello World! from console.log")

	return nil
}

func main() {
	<-done
}

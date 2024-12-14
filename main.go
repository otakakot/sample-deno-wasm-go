//go:generate sh -c "GOOS=js GOARCH=wasm go build -o main.wasm main.go"
package main

import (
	"syscall/js"
)

var done = make(chan struct{})

func init() {
	js.Global().Set("golog", js.FuncOf(golog))
}

func golog(this js.Value, args []js.Value) any {
	defer func() {
		done <- struct{}{}
	}()

	println(args[0].String())

	return nil
}

func main() {
	<-done
}

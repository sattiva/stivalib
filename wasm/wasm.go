//go:build js && wasm

package wasm

import (
	"crypto/subtle"
	"syscall/js"
)

func RegisterWasmExports() {
	js.Global().Set("sativasConstantCompare", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 2 {
			return false
		}
		a := args[0].String()
		b := args[1].String()
		return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
	}))
}

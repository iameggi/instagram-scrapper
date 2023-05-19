
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	document := js.Global().Get("document")
	postsDiv := document.Call("getElementById", "posts")

	fetch := js.Global().Get("fetch")
	fetch.Call("ig/v1/?username=iameggiii").Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response := args[0]
		response.Call("json").Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			data := args[0]

			for _, post := range data.Index() {
				postElement := document.Call("createElement", "div")
				postElement.Set("innerHTML", fmt.Sprintf(`
				  <h2>%v</h2>
				  <img src="%v">
				`, post.Get("Caption"), post.Get("ImageURL")))

				postsDiv.Call("appendChild", postElement)
			}

			return nil
		}))

		return nil
	})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		error := args[0]
		fmt.Println("Error:", error)
		return nil
	}))

	<-c
}

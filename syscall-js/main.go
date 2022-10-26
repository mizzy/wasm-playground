package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	// グローバルオブジェクト(window)を取得します
	window := js.Global()

	// document オブジェクトを取得します
	document := window.Get("document")

	// bodyを取得します
	body := document.Get("body")

	// p のDOMを作成します
	counter := 0
	p := document.Call("createElement", "p")
	p.Set("id", "counter")
	p.Set("innerHTML", strconv.Itoa(counter))

	// ボタンのDOMを作成し、Clickイベントを設定します
	btn := document.Call("createElement", "button")
	btn.Set("textContent", "count up!")
	btn.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
		counter++
		fmt.Println(counter)                                                               // console.logに出力します
		document.Call("getElementById", "counter").Set("innerHTML", strconv.Itoa(counter)) // カウンターの表示を更新します

		return nil
	}))

	// pをbodyに追加します
	body.Call("appendChild", p)
	// ボタンをbodyに追加します
	body.Call("appendChild", btn)

	// プログラムが終了しないように待機します
	select {}
}

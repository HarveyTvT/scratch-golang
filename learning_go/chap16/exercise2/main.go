package main

import (
	"fmt"
	"unsafe"
)

type OrderInfo struct {
	OrderCode   rune     // 4 bytes
	Amount      int      // 8 bytes
	OrderNumber uint16   // 2 bytes
	Items       []string // 24 bytes
	IsReady     bool     // 1 bytes
}

type SmallOrderInfo struct {
	OrderCode   rune     // 4 bytes
	OrderNumber uint16   // 2 bytes
	IsReady     bool     // 1 bytes
	Amount      int      // 8 bytes
	Items       []string // 24 bytes
}

func main() {
	foo := OrderInfo{}
	fmt.Println(unsafe.Sizeof(foo), unsafe.Offsetof(foo.OrderCode), unsafe.Offsetof(foo.Amount), unsafe.Offsetof(foo.OrderNumber), unsafe.Offsetof(foo.Items), unsafe.Offsetof(foo.IsReady))

	bar := SmallOrderInfo{}
	fmt.Println(unsafe.Sizeof(bar),
		unsafe.Offsetof(bar.OrderCode),
		unsafe.Offsetof(bar.OrderNumber),
		unsafe.Offsetof(bar.IsReady),
		unsafe.Offsetof(bar.Amount),
		unsafe.Offsetof(bar.Items),
	)
}

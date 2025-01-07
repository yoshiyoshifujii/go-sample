package main

import "fmt"

// Color Enumタイプの定義
type Color string

const (
	Red   Color = "Red"
	Green Color = "Green"
	Blue  Color = "Blue"
)

func PrintColor(c Color) {
	switch c {
	case Red, Green, Blue:
		fmt.Println(c)
	default:
		fmt.Println("Unknown color")
	}
}

func main() {
	PrintColor(Red)
	PrintColor("Yellow")
}

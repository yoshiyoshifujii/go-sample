package main

import "fmt"

// Color Enumタイプの定義
type Color int

const (
	Red Color = iota
	Green
	Blue
)

func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	default:
		return "Unknown"
	}
}

func main() {
	var c Color = Green
	fmt.Println(c)
}

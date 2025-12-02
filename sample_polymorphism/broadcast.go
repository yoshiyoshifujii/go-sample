package main

// 共通の Broadcast 実装: Notify(string) を持つ型をまとめて処理する。
func broadcast[T interface {
	Notify(string) string
}](ns []T, message string) []string {
	out := make([]string, 0, len(ns))
	for _, n := range ns {
		out = append(out, n.Notify(message))
	}
	return out
}

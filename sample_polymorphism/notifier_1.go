package main

// Version 1: interface + value receivers (シンプルな多態).

type (
	Notifier1 interface {
		Notify(message string) string
	}
	EmailNotifier1 struct {
		Address string
	}
	SMSNotifier1 struct {
		Number string
	}
)

func (e EmailNotifier1) Notify(message string) string {
	return "Email to " + e.Address + ": " + message
}

func (s SMSNotifier1) Notify(message string) string {
	return "SMS to " + s.Number + ": " + message
}

func BroadcastV1(ns []Notifier1, message string) []string {
	out := make([]string, 0, len(ns))
	for _, n := range ns {
		out = append(out, n.Notify(message))
	}
	return out
}

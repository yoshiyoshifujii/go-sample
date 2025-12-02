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

package main

import "encoding/json"

// Version 2: channel interfaceをフィールドに埋め込み、ポインタ実装を隠蔽.

type (
	channel2 interface {
		Send(message string) string
	}
	Notifier2 struct {
		ch channel2
	}
	EmailNotifier2 struct {
		Address string
	}
	SMSNotifier2 struct {
		Number string
	}
)

func NewEmailNotifier2(address string) Notifier2 {
	return Notifier2{ch: &EmailNotifier2{Address: address}}
}

func NewSMSNotifier2(number string) Notifier2 {
	return Notifier2{ch: &SMSNotifier2{Number: number}}
}

func (n Notifier2) Notify(message string) string {
	if n.ch == nil {
		return "Notifier channel missing backend"
	}
	return n.ch.Send(message)
}

// MarshalJSON ensures the serialized form matches the concrete backend fields.
func (n Notifier2) MarshalJSON() ([]byte, error) {
	switch v := n.ch.(type) {
	case *EmailNotifier2:
		return json.Marshal(struct {
			Address string
		}{Address: v.Address})
	case *SMSNotifier2:
		return json.Marshal(struct {
			Number string
		}{Number: v.Number})
	default:
		return json.Marshal(struct{}{})
	}
}

func (e *EmailNotifier2) Send(message string) string {
	return "Email to " + e.Address + ": " + message
}

func (s *SMSNotifier2) Send(message string) string {
	return "SMS to " + s.Number + ": " + message
}

func BroadcastV2(ns []Notifier2, message string) []string {
	out := make([]string, 0, len(ns))
	for _, n := range ns {
		out = append(out, n.Notify(message))
	}
	return out
}

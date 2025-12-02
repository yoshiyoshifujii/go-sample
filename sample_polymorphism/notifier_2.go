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

const (
	notifier2TypeEmail = "email"
	notifier2TypeSMS   = "sms"
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
			Type    string `json:"type"`
			Address string `json:"address"`
		}{Type: notifier2TypeEmail, Address: v.Address})
	case *SMSNotifier2:
		return json.Marshal(struct {
			Type   string `json:"type"`
			Number string `json:"number"`
		}{Type: notifier2TypeSMS, Number: v.Number})
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

// UnmarshalJSON reconstructs the backend from serialized fields.
func (n *Notifier2) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type    string `json:"type"`
		Address string `json:"address,omitempty"`
		Number  string `json:"number,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw.Type {
	case notifier2TypeEmail:
		n.ch = &EmailNotifier2{Address: raw.Address}
	case notifier2TypeSMS:
		n.ch = &SMSNotifier2{Number: raw.Number}
	default:
		n.ch = nil
	}
	return nil
}

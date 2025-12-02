package main

import (
	"encoding/json"
	"fmt"
)

// Version 1: interface + value receivers (シンプルな多態).

type (
	notifier1Type string

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

const (
	notifier1TypeEmail notifier1Type = "email"
	notifier1TypeSMS   notifier1Type = "sms"
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

func (e EmailNotifier1) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type    notifier1Type `json:"type"`
		Address string        `json:"address"`
	}{
		Type:    notifier1TypeEmail,
		Address: e.Address,
	})
}

func (s SMSNotifier1) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type   notifier1Type `json:"type"`
		Number string        `json:"number"`
	}{
		Type:   notifier1TypeSMS,
		Number: s.Number,
	})
}

// UnmarshalNotifier1List reconstructs concrete Notifier1 values from JSON.
func UnmarshalNotifier1List(data []byte) ([]Notifier1, error) {
	var raw []struct {
		Type    notifier1Type `json:"type"`
		Address string        `json:"address,omitempty"`
		Number  string        `json:"number,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	out := make([]Notifier1, 0, len(raw))
	for _, item := range raw {
		switch item.Type {
		case notifier1TypeEmail:
			out = append(out, EmailNotifier1{Address: item.Address})
		case notifier1TypeSMS:
			out = append(out, SMSNotifier1{Number: item.Number})
		default:
			return nil, fmt.Errorf("unknown notifier payload: %+v", item)
		}
	}
	return out, nil
}

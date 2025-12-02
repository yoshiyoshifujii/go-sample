package main

import (
	"encoding/json"
	"fmt"
)

type notifierPayload struct {
	Type    string `json:"type"`
	Address string `json:"address,omitempty"`
	Number  string `json:"number,omitempty"`
}

func toPayloads[T any](ns []T, convert func(T) (notifierPayload, error)) ([]notifierPayload, error) {
	out := make([]notifierPayload, 0, len(ns))
	for _, n := range ns {
		payload, err := convert(n)
		if err != nil {
			return nil, err
		}
		out = append(out, payload)
	}
	return out, nil
}

func fromPayloads[T any](raw []notifierPayload, convert func(notifierPayload) (T, error)) ([]T, error) {
	out := make([]T, 0, len(raw))
	for _, p := range raw {
		n, err := convert(p)
		if err != nil {
			return nil, err
		}
		out = append(out, n)
	}
	return out, nil
}

// V1
func toPayloadsV1(ns []Notifier1) ([]notifierPayload, error) {
	return toPayloads(ns, func(n Notifier1) (notifierPayload, error) {
		switch v := n.(type) {
		case EmailNotifier1:
			return notifierPayload{Type: "email", Address: v.Address}, nil
		case SMSNotifier1:
			return notifierPayload{Type: "sms", Number: v.Number}, nil
		default:
			return notifierPayload{}, fmt.Errorf("unknown notifier1 type %T", n)
		}
	})
}

func fromPayloadsV1(raw []notifierPayload) ([]Notifier1, error) {
	return fromPayloads(raw, func(p notifierPayload) (Notifier1, error) {
		switch p.Type {
		case "email":
			return EmailNotifier1{Address: p.Address}, nil
		case "sms":
			return SMSNotifier1{Number: p.Number}, nil
		default:
			return nil, fmt.Errorf("unknown notifier1 payload type %q", p.Type)
		}
	})
}

// V2
func toPayloadsV2(ns []Notifier2) ([]notifierPayload, error) {
	return toPayloads(ns, func(n Notifier2) (notifierPayload, error) {
		switch v := n.ch.(type) {
		case *EmailNotifier2:
			return notifierPayload{Type: "email", Address: v.Address}, nil
		case *SMSNotifier2:
			return notifierPayload{Type: "sms", Number: v.Number}, nil
		default:
			return notifierPayload{}, fmt.Errorf("unknown notifier2 backend %T", n.ch)
		}
	})
}

func fromPayloadsV2(raw []notifierPayload) ([]Notifier2, error) {
	return fromPayloads(raw, func(p notifierPayload) (Notifier2, error) {
		switch p.Type {
		case "email":
			return NewEmailNotifier2(p.Address), nil
		case "sms":
			return NewSMSNotifier2(p.Number), nil
		default:
			return Notifier2{}, fmt.Errorf("unknown notifier2 payload type %q", p.Type)
		}
	})
}

// V3
func toPayloadsV3(ns []Notifier3) ([]notifierPayload, error) {
	return toPayloads(ns, func(n Notifier3) (notifierPayload, error) {
		switch n.notifierType {
		case notifierTypeEmail:
			return notifierPayload{Type: "email", Address: n.emailNotifier.Address}, nil
		case notifierTypeSMS:
			return notifierPayload{Type: "sms", Number: n.smsNotifier.Number}, nil
		default:
			return notifierPayload{}, fmt.Errorf("unknown notifier3 type %q", n.notifierType)
		}
	})
}

func fromPayloadsV3(raw []notifierPayload) ([]Notifier3, error) {
	return fromPayloads(raw, func(p notifierPayload) (Notifier3, error) {
		switch p.Type {
		case "email":
			return NewEmailNotifier3(p.Address), nil
		case "sms":
			return NewSMSNotifier3(p.Number), nil
		default:
			return Notifier3{}, fmt.Errorf("unknown notifier3 payload type %q", p.Type)
		}
	})
}

// V4
func toPayloadsV4(ns []Notifier4) ([]notifierPayload, error) {
	return toPayloads(ns, func(n Notifier4) (notifierPayload, error) {
		if n.payload.Type == "" {
			return notifierPayload{}, fmt.Errorf("unknown notifier4 payload type (empty)")
		}
		return n.payload, nil
	})
}

func fromPayloadsV4(raw []notifierPayload) ([]Notifier4, error) {
	return fromPayloads(raw, func(p notifierPayload) (Notifier4, error) {
		switch p.Type {
		case "email":
			return NewEmailNotifier4(p.Address), nil
		case "sms":
			return NewSMSNotifier4(p.Number), nil
		default:
			return Notifier4{}, fmt.Errorf("unknown notifier4 payload type %q", p.Type)
		}
	})
}

// Helpers for tests (serialize/deserialize lists)
func marshalPayloads(payloads []notifierPayload) ([]byte, error) {
	return json.Marshal(payloads)
}

func unmarshalPayloads(data []byte) ([]notifierPayload, error) {
	var payloads []notifierPayload
	if err := json.Unmarshal(data, &payloads); err != nil {
		return nil, err
	}
	return payloads, nil
}

package main

type (
	notifierType string
)

const (
	notifierTypeEmail notifierType = "email"
	notifierTypeSMS   notifierType = "sms"
)

type (
	Notifier3 struct {
		notifierType  notifierType
		emailNotifier *EmailNotifier3
		smsNotifier   *SMSNotifier3
	}

	EmailNotifier3 struct {
		Address string
	}

	SMSNotifier3 struct {
		Number string
	}
)

// NewEmailNotifier3 creates an email notifier instance.
func NewEmailNotifier3(address string) Notifier3 {
	return Notifier3{
		notifierType:  notifierTypeEmail,
		emailNotifier: &EmailNotifier3{Address: address},
	}
}

// NewSMSNotifier3 creates an SMS notifier instance.
func NewSMSNotifier3(number string) Notifier3 {
	return Notifier3{
		notifierType: notifierTypeSMS,
		smsNotifier:  &SMSNotifier3{Number: number},
	}
}

// Notify dispatches the notification based on the stored type.
func (n Notifier3) Notify(message string) string {
	switch n.notifierType {
	case notifierTypeEmail:
		if n.emailNotifier == nil {
			return "Email notifier missing backend"
		}
		return n.emailNotifier.Notify(message)
	case notifierTypeSMS:
		if n.smsNotifier == nil {
			return "SMS notifier missing backend"
		}
		return n.smsNotifier.Notify(message)
	default:
		return "Unknown notifier"
	}
}

func (e *EmailNotifier3) Notify(message string) string {
	return "Email to " + e.Address + ": " + message
}

func (s *SMSNotifier3) Notify(message string) string {
	return "SMS to " + s.Number + ": " + message
}

// BroadcastV3 sends using tag-based dispatch.
func BroadcastV3(ns []Notifier3, message string) []string {
	out := make([]string, 0, len(ns))
	for _, n := range ns {
		out = append(out, n.Notify(message))
	}
	return out
}

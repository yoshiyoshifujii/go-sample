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

// dispatchNotifier3 runs the provided callback based on the notifier type.
func dispatchNotifier3[T any](
	n Notifier3,
	email func(*EmailNotifier3) (T, error),
	sms func(*SMSNotifier3) (T, error),
) (T, error) {
	switch n.notifierType {
	case notifierTypeEmail:
		if n.emailNotifier == nil {
			panic("email notifier missing backend")
		}
		return email(n.emailNotifier)
	case notifierTypeSMS:
		if n.smsNotifier == nil {
			panic("sms notifier missing backend")
		}
		return sms(n.smsNotifier)
	default:
		panic("unknown notifier type: " + string(n.notifierType))
	}
}

// Notify dispatches the notification based on the stored type.
func (n Notifier3) Notify(message string) string {
	res, err := dispatchNotifier3(
		n,
		func(e *EmailNotifier3) (string, error) {
			return e.Notify(message), nil
		},
		func(s *SMSNotifier3) (string, error) {
			return s.Notify(message), nil
		},
	)
	if err != nil {
		panic(err)
	}
	return res
}

func (e *EmailNotifier3) Notify(message string) string {
	return "Email to " + e.Address + ": " + message
}

func (s *SMSNotifier3) Notify(message string) string {
	return "SMS to " + s.Number + ": " + message
}

package main

// Version 4: 関数型の戦略注入 (notify関数を閉じ込める構造体).

type Notifier4 struct {
	notify  func(message string) string
	payload notifierPayload
}

func NewEmailNotifier4(address string) Notifier4 {
	payload := notifierPayload{Type: "email", Address: address}
	return Notifier4{
		payload: payload,
		notify: func(message string) string {
			return "Email to " + address + ": " + message
		},
	}
}

func NewSMSNotifier4(number string) Notifier4 {
	payload := notifierPayload{Type: "sms", Number: number}
	return Notifier4{
		payload: payload,
		notify: func(message string) string {
			return "SMS to " + number + ": " + message
		},
	}
}

func (n Notifier4) Notify(message string) string {
	if n.notify == nil {
		panic("Notifier strategy missing backend")
	}
	return n.notify(message)
}

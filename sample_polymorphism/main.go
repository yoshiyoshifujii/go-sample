package main

import "fmt"

func main() {
	notifiersV1 := []Notifier1{
		EmailNotifier1{Address: "user@example.com"},
		SMSNotifier1{Number: "+81000000000"},
	}

	fmt.Println("== Version 1 (interface, value receivers) ==")
	for _, msg := range BroadcastV1(notifiersV1, "Deployment finished (v1)") {
		fmt.Println(msg)
	}

	notifiersV2 := []Notifier2{
		NewEmailNotifier2("user@example.com"),
		NewSMSNotifier2("+81000000000"),
	}

	fmt.Println("== Version 2 (tagged struct dispatch) ==")
	for _, msg := range BroadcastV2(notifiersV2, "Deployment finished (v2)") {
		fmt.Println(msg)
	}

	notifiersV3 := []Notifier3{
		NewEmailNotifier3("user@example.com"),
		NewSMSNotifier3("+81000000000"),
	}

	fmt.Println("== Version 3 (tagged struct with pointer backends) ==")
	for _, msg := range BroadcastV3(notifiersV3, "Deployment finished (v3)") {
		fmt.Println(msg)
	}
}

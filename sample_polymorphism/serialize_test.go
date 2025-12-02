package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerializeV1(t *testing.T) {
	notifiers := []Notifier1{
		EmailNotifier1{Address: "user@example.com"},
		SMSNotifier1{Number: "+81000000000"},
	}

	b, err := json.Marshal(notifiers)
	require.NoError(t, err)
	assert.JSONEq(t, `[{"type":"email","address":"user@example.com"},{"type":"sms","number":"+81000000000"}]`, string(b))

	got, err := UnmarshalNotifier1List(b)
	require.NoError(t, err)
	assert.Equal(t, notifiers, got)
}

func TestSerializeV2(t *testing.T) {
	notifiers := []Notifier2{
		NewEmailNotifier2("user@example.com"),
		NewSMSNotifier2("+81000000000"),
	}

	b, err := json.Marshal(notifiers)
	require.NoError(t, err)
	assert.JSONEq(t, `[{"type":"email","address":"user@example.com"},{"type":"sms","number":"+81000000000"}]`, string(b))

	var got []Notifier2
	require.NoError(t, json.Unmarshal(b, &got))
	assert.Equal(t, notifiers, got)
}

func TestSerializeV3(t *testing.T) {
	notifiers := []Notifier3{
		NewEmailNotifier3("user@example.com"),
		NewSMSNotifier3("+81000000000"),
	}

	b, err := json.Marshal(notifiers)
	require.NoError(t, err)
	assert.JSONEq(t, `[{"type":"email","address":"user@example.com"},{"type":"sms","number":"+81000000000"}]`, string(b))

	var got []Notifier3
	require.NoError(t, json.Unmarshal(b, &got))
	assert.Equal(t, notifiers, got)
}

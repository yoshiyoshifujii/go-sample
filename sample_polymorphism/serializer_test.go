package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerializeV1(t *testing.T) {
	notifiers := []Notifier1{
		EmailNotifier1{Address: "user@example.com"},
		SMSNotifier1{Number: "+81000000000"},
	}

	payloads, err := toPayloadsV1(notifiers)
	require.NoError(t, err)

	b, err := marshalPayloads(payloads)
	require.NoError(t, err)
	assert.JSONEq(t, `[{"type":"email","address":"user@example.com"},{"type":"sms","number":"+81000000000"}]`, string(b))

	decoded, err := unmarshalPayloads(b)
	require.NoError(t, err)
	got, err := fromPayloadsV1(decoded)
	require.NoError(t, err)
	assert.Equal(t, notifiers, got)
}

func TestSerializeV2(t *testing.T) {
	notifiers := []Notifier2{
		NewEmailNotifier2("user@example.com"),
		NewSMSNotifier2("+81000000000"),
	}

	payloads, err := toPayloadsV2(notifiers)
	require.NoError(t, err)

	b, err := marshalPayloads(payloads)
	require.NoError(t, err)
	assert.JSONEq(t, `[{"type":"email","address":"user@example.com"},{"type":"sms","number":"+81000000000"}]`, string(b))

	decoded, err := unmarshalPayloads(b)
	require.NoError(t, err)
	got, err := fromPayloadsV2(decoded)
	require.NoError(t, err)
	assert.Equal(t, notifiers, got)
}

func TestSerializeV3(t *testing.T) {
	notifiers := []Notifier3{
		NewEmailNotifier3("user@example.com"),
		NewSMSNotifier3("+81000000000"),
	}

	payloads, err := toPayloadsV3(notifiers)
	require.NoError(t, err)

	b, err := marshalPayloads(payloads)
	require.NoError(t, err)
	assert.JSONEq(t, `[{"type":"email","address":"user@example.com"},{"type":"sms","number":"+81000000000"}]`, string(b))

	decoded, err := unmarshalPayloads(b)
	require.NoError(t, err)
	got, err := fromPayloadsV3(decoded)
	require.NoError(t, err)
	assert.Equal(t, notifiers, got)
}

func TestSerializeV4(t *testing.T) {
	notifiers := []Notifier4{
		NewEmailNotifier4("user@example.com"),
		NewSMSNotifier4("+81000000000"),
	}

	payloads, err := toPayloadsV4(notifiers)
	require.NoError(t, err)

	b, err := marshalPayloads(payloads)
	require.NoError(t, err)
	assert.JSONEq(t, `[{"type":"email","address":"user@example.com"},{"type":"sms","number":"+81000000000"}]`, string(b))

	decoded, err := unmarshalPayloads(b)
	require.NoError(t, err)
	got, err := fromPayloadsV4(decoded)
	require.NoError(t, err)

	roundTripPayloads, err := toPayloadsV4(got)
	require.NoError(t, err)
	assert.Equal(t, payloads, roundTripPayloads)
}

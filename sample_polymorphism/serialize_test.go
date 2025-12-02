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
	assert.JSONEq(t, `[{"Address":"user@example.com"},{"Number":"+81000000000"}]`, string(b))
}

func TestSerializeV2(t *testing.T) {
	notifiers := []Notifier2{
		NewEmailNotifier2("user@example.com"),
		NewSMSNotifier2("+81000000000"),
	}

	b, err := json.Marshal(notifiers)
	require.NoError(t, err)
	assert.Equal(t, `[{},{}]`, string(b)) // chは非公開フィールドのためシリアライズ対象外
}

func TestSerializeV3(t *testing.T) {
	notifiers := []Notifier3{
		NewEmailNotifier3("user@example.com"),
		NewSMSNotifier3("+81000000000"),
	}

	b, err := json.Marshal(notifiers)
	require.NoError(t, err)
	assert.Equal(t, `[{},{}]`, string(b)) // タグやフィールドが非公開のため空オブジェクトになる
}

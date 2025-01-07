package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockService struct{}

func (ms *MockService) DoSomething() string {
	return "Mock Service is doing something"
}

// TestController_HandleRequest テスト関数
func TestController_HandleRequest(t *testing.T) {
	// 依存関係を注入
	service := MockService{}
	controller := NewController(&service)

	// 処理実行
	result := controller.HandleRequest()

	// 結果確認
	assert.Equal(t, "Mock Service is doing something", result)
}

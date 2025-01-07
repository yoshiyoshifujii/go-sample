package main

import "fmt"

// Service インタフェース定義
type Service interface {
	DoSomething() string
}

// RealService Serviceの実装
type RealService struct{}

func (rs RealService) DoSomething() string {
	return "Real Service is doing something"
}

// Controller 構造体(依存先)
type Controller struct {
	service Service
}

// NewController コンストラクタ関数(依存関係を注入)
func NewController(service Service) *Controller {
	return &Controller{service: service}
}

// HandleRequest メソッド
func (c *Controller) HandleRequest() string {
	return c.service.DoSomething()
}

func main() {

	// 依存関係を注入
	service := RealService{}
	controller := NewController(service)

	// 処理実行
	result := controller.HandleRequest()

	// 結果出力
	fmt.Println(result)

}

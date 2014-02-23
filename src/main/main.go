package main

import (
	"fmt"
	"test"
)

func main() {
	directions := []string{"S2N", "S2W", "E2W", "E2S", "N2S", "N2E", "W2E", "W2N", "S2E", "E2N", "N2W", "W2S"}
	//创建12条路
	for i := 0; i < len(directions); i++ {
		new(test.Road).Init(directions[i])
	}
	//打开控制器
	new(test.LampController).Init()

	//避免主goroutine迅速运行完
	var input string
	fmt.Scanln(&input)
}

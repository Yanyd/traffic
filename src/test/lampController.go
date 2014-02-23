package test

import (
	"time"
)

type LampController struct {
	currentLamp *Lamp
}

//var currentLamp *Lamp

func (this *LampController) Init() {
	this.currentLamp = LampMap["S2N"]
	this.currentLamp.light()

	//每隔10s就变换红绿灯
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for _ = range ticker.C {
			this.currentLamp = this.currentLamp.blackOut()
		}
	}()
}

package test

import (
	"fmt"
)

type Lamp struct {
	name     string
	opposite string //相对应的信号灯
	next     string //下一个要变亮的信号灯
	lighted  bool   //当前信号灯的状态
}

//12个信号灯
var LampMap = map[string]*Lamp{
	"S2N": &Lamp{"S2N", "N2S", "S2W", false},
	"S2W": &Lamp{"S2W", "N2E", "E2W", false},
	"E2W": &Lamp{"E2W", "W2E", "E2S", false},
	"E2S": &Lamp{"E2S", "W2N", "S2N", false},
	"N2S": &Lamp{"N2S", "", "", false},
	"N2E": &Lamp{"N2E", "", "", false},
	"W2E": &Lamp{"W2E", "", "", false},
	"W2N": &Lamp{"W2N", "", "", false},
	"S2E": &Lamp{"S2E", "", "", true}, //常绿
	"E2N": &Lamp{"E2N", "", "", true},
	"N2W": &Lamp{"N2W", "", "", true},
	"W2S": &Lamp{"W2S", "", "", true},
}

func (this *Lamp) light() {
	this.lighted = true
	if "" != this.opposite { //避免死循环
		LampMap[this.opposite].light()
	}
	fmt.Println(this.name + " lamp is green ,下面共应该有6个方向能看到汽车穿过!")
}

func (this *Lamp) blackOut() *Lamp {
	this.lighted = false
	if "" != this.opposite {
		LampMap[this.opposite].blackOut()
	}
	//下一个变绿的的灯
	nextLamp := new(Lamp)
	if "" != this.next {
		nextLamp = LampMap[this.next]
		fmt.Println("绿灯从" + this.name + " 切换为------> " + this.next)
		nextLamp.light()
	}
	return nextLamp
}

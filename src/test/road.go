package test

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

type Road struct {
	name      string    //每条道路都有一个名字 S2N..
	vechicles list.List //道路上的车辆集合
}

func (this *Road) Init(name string) {
	this.name = name
	//每隔一段时间道路上就添加一辆车
	go func() {
		rand.Seed(time.Now().UnixNano())
		for i := 1; i < 1000; i++ {
			sleepTime := (rand.Intn(10) + 1) * 1e6
			time.Sleep(time.Duration(sleepTime))
			this.vechicles.PushBack(name + "_" + fmt.Sprintf("%d", i))
		}
	}()
	//每隔一秒去查看此道路的信号灯是否为绿  使用定时器
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			if 0 < this.vechicles.Len() {
				lighted := LampMap[name].lighted
				if lighted {
					first := this.vechicles.Front()
					this.vechicles.Remove(first)
					fmt.Println(first.Value.(string) + " is traversing !")
				}
			}
		}
	}()
}

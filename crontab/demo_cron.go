package crontab

import (
	"fmt"
	"time"
)

type DemoCron struct {}

func (c DemoCron) Run()  {
	fmt.Println("定时任务开启，打印时间:",time.Now().Format("2006-01-02 15:04:05"))
}


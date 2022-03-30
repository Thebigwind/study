package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	//"github.com/alibaba/sentinel-golang/util"
)

func main() {
	_ = sentinel.InitDefault()
	_, _ = flow.LoadRules([]*flow.Rule{
		{
			Resource:        "控制吃煎鱼的速度",
			Threshold:       60,
			ControlBehavior: flow.Reject,
		},
	})

	e, b := sentinel.Entry("控制吃煎鱼的速度", sentinel.WithTrafficType(base.Inbound))
	if b != nil {
		// Blocked
	} else {
		// Passed
		e.Exit()
	}
	fmt.Println("嘻嘻嘻")
}

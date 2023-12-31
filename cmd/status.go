package cmd

import (
	"fmt"
	"todo-cli-go/service"
)

type StatusPuppet struct {
	statusMaster service.StatusMaster
}

func NewStatusPuppet(status service.StatusMaster) *StatusPuppet {
	return &StatusPuppet{statusMaster: status}
}

func (p StatusPuppet) Done() {
	doneStats, err := p.statusMaster.GetDone()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(doneStats.String())
}

func (p StatusPuppet) Status() {
	stats, err := p.statusMaster.GetOverall()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, status := range stats {
		fmt.Println(status.String())
	}
}

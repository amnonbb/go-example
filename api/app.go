package api

import (
	"fmt"
	"github.com/Bnei-Baruch/go-example/pkg/workflow"
)

type App struct {
	A string
}


func (a *App) InitApp() {
	json := []byte(`{"capture_id":"c123123123123","date":"2021-12-12"}`)
	workflow.WorkflowState(json)
	getState()
}

func getState() {

	var cs *workflow.State

	//cs := workflow.CaptureState{}
	fmt.Println("getState:", &cs)
	fmt.Println("getState:", cs)
}
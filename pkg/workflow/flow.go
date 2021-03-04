package workflow

import (
	"encoding/json"
	"fmt"
)

type State struct {
	CaptureID string `json:"capture_id"`
	Date      string `json:"date"`
}


func WorkflowState(m []byte) {
	cs := &State{}
	fmt.Println("State:", cs)
	err := json.Unmarshal(m, &cs)
	fmt.Println("State:", cs)
	if err != nil {
		fmt.Println(err.Error())
	}
}
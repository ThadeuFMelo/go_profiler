package models

import "time"

type ProcessMessage struct {
	Pid       uint32    `json:"pid"`
	Cpu       float64   `json:"cpu"`
	Mem       float32   `json:"mem"`
	Name      string    `json:"name"`
	TimeStamp time.Time `json:"time"`
	Ctime     int64     `json:"ctime"`
}

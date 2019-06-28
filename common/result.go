package common

import (
	"encoding/json"
	"time"
)

type ResultCode int16

const (
	Success ResultCode = iota
	ParamNull
)

type Result struct {
	Code ResultCode  `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func GetResult() Result {
	return Result{Success, nil, ""}
}

type TimeStamp time.Time

func (t TimeStamp) MarshalJSON() ([]byte, error) {
	myTime := time.Time(t)
	rs := myTime.Format("2006-1-2 15:04:05")
	if rs == "0001-1-1 00:00:00" {
		return json.Marshal(nil)
	} else {
		js, er := json.Marshal(rs)
		return js, er
	}
}

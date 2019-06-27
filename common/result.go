package common

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

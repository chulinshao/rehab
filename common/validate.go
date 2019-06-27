package common

import "errors"

func IsNull2(params ...interface{}) error {
	for _, param := range params {
		if param == nil || param == "" {
			return errors.New("参数不能为空")
		}
	}
	return nil
}

func IsNull(result *Result, param interface{}, paramName string) error {
	if param == nil || param == "" {
		result.Code = ParamNull
		result.Msg = paramName + "参数不能为空"
		return errors.New(result.Msg)
	}
	return nil
}

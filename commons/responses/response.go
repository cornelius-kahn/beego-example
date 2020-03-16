package responses

type ReturnStruct struct {
	Ret   int         `json:"ret"`
	Msg   string      `json:"msg"`
	Count int64       `json:"count,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func SuccessReturn() *ReturnStruct {
	result := new(ReturnStruct)
	result.Ret = 0
	result.Msg = "ok"
	return result
}

func SuccessDataReturn(data interface{}) *ReturnStruct {
	result := new(ReturnStruct)
	result.Ret = 0
	result.Msg = "ok"
	result.Data = data
	return result
}

func SuccessDataPageReturn(count int64, data interface{}) *ReturnStruct {
	result := new(ReturnStruct)
	result.Ret = 0
	result.Msg = "ok"
	result.Count = count
	result.Data = data
	return result
}

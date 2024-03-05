package models

type VO struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func (v VO) Success(data any) VO {
	v.Code = 200
	v.Data = data
	v.Message = "success"
	return v
}

func (v VO) Fail(msg string) VO {
	v.Code = 500
	v.Data = nil
	v.Message = msg
	return v
}

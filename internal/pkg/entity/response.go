package entity

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *Response) ReplaceData(d interface{}) *Response {
	r.Data = d
	return r
}

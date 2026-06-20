package types

import (
	"github.com/Y1le/gotolist/consts"
)

// Response 鍩虹搴忓垪鍖栧櫒
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Error   string      `json:"error"`
	TrackId string      `json:"track_id"`
}

// DataList 甯︽湁鎬绘暟鐨凞ata缁撴瀯
type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

// // TokenData 甯︽湁token鐨凞ata缁撴瀯
// type TokenData struct {
// 	User         interface{} `json:"user"`
// 	AccessToken  string      `json:"access_token"`
// 	RefreshToken string      `json:"refresh_token"`
// }

// RespList 甯︽湁鎬绘暟鐨勫垪琛ㄦ瀯寤哄櫒
func RespList(items interface{}, total int64) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}

// RespSuccess 鎴愬姛杩斿洖
func RespSuccess(code ...int) *Response {
	status := consts.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   "鎿嶄綔鎴愬姛",
		Msg:    consts.GetMsg(status),
	}

	return r
}

// RespSuccessWithData 甯ata鎴愬姛杩斿洖
func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := consts.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    consts.GetMsg(status),
	}

	return r
}

// RespError 閿欒杩斿洖
func RespError(err error, data string, code ...int) *Response {
	status := consts.ERROR
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Msg:    consts.GetMsg(status),
		Data:   data,
		Error:  err.Error(),
	}

	return r
}

package qhttpx

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type State = string

var defaultStateInfo = map[State]string{
	StateOk:               "ok",
	StateInvalidParameter: "invalid params",
	StateFailed:           "failed",
}

type CommonResponse struct {
	State     string      `json:"state"`               //状态值
	StateInfo string      `json:"state_info"`          //状态描述
	CustInfo  string      `json:"cust_info,omitempty"` //自定义描述
	Seq       int         `json:"seq"`
	Data      interface{} `json:"data"`
}
type CommonErrorMsg struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func state(s State) string {
	info := defaultStateInfo[s]
	if info == "" {
		return s
	}
	return info
}

func send(w http.ResponseWriter, code State, data interface{}, stateInfo, CustomInfo string, seq int) {
	info := stateInfo
	if info == "" {
		info = state(code)
	}
	ret := CommonResponse{
		State:     code,
		StateInfo: info,
		CustInfo:  CustomInfo,
		Seq:       seq,
		Data:      data,
	}
	httpx.WriteJson(w, http.StatusOK, &ret)
}

func sendOk(w http.ResponseWriter) {
	send(w, StateOk, "", "", "", 0)
}
func sendData(w http.ResponseWriter, data interface{}) {
	send(w, StateOk, data, "", "", 0)
}
func sendFailed(w http.ResponseWriter, data interface{}) {
	errStr := fmt.Sprintf("%s", data)
	send(w, StateFailed, "", "", errStr, 0)
}
func sendCode(w http.ResponseWriter, code State, stateInfo string, data interface{}) {
	send(w, code, data, stateInfo, "", 0)
}
func sendParamError(w http.ResponseWriter, data interface{}) {
	errStr := fmt.Sprintf("%s", data)
	send(w, StateInvalidParameter, "", "", errStr, 0)
}

func SendOk(w http.ResponseWriter) {
	sendOk(w)
}
func SendData(w http.ResponseWriter, data interface{}) {
	sendData(w, data)
}

func SendFailed(w http.ResponseWriter, data interface{}) {
	sendFailed(w, data)
}
func SendCode(w http.ResponseWriter, code State, stateInfo string, data interface{}) {
	sendCode(w, code, stateInfo, data)
}
func SendParamError(w http.ResponseWriter, data interface{}) {
	sendParamError(w, data)
}

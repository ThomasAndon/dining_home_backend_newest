// Code generated by goctl. DO NOT EDIT.
package types

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type BondRequest struct {
	Identity string `json:"identity"`
}
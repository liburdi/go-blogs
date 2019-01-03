package init

import (
	"encoding/json"
	"net/http"
)

type ApiRestful struct {
	Code int
	Message string
}

/**
 *restful api return
 */
func (api *ApiRestful) ApiRestful(w http.ResponseWriter) error {
	//方案一
	//var body json.RawMessage
	////todo
	//
	//b, _ := body.MarshalJSON()
	//w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	//w.Write(b)
	//方案二
	w.Header().Set("Access-Control-Allow-Origin", "*")
	resData , err := json. Marshal ( api )
	if err!=nil{
		return err
	}
	_,err=w.Write(resData)
	if err!=nil{
		return err
	}
	return nil
}

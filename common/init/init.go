package init

import (
	"encoding/json"
	"net/http"
)

type ApiRestful struct {
	Code int
	Message string
	Data interface{}
}
type Seo struct{
	PageTitle string
	Keywords string
	Description string
}

/**
 * common restful api response
 */
func (api *ApiRestful) ApiRestful(w http.ResponseWriter) error {
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

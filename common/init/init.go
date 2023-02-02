package init

import (
	"encoding/json"
	"github.com/liburdi/go-blogs/config"
	"net/http"
)

type ApiRestful struct {
	Code    int
	Message string
	Data    interface{}
}
type Seo struct {
	PageTitle   string
	Keywords    string
	Description string
}

/**
 * common restful api response
 */
func (api *ApiRestful) ApiRestful(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Connection", "close")
	resData, err := json.Marshal(api)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write(resData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func Success(w http.ResponseWriter) {
	api := ApiRestful{
		Code:    200,
		Message: "Success",
		Data:    map[string]interface{}{},
	}

	api.ApiRestful(w)
}

func Error(code int, w http.ResponseWriter) {
	message := config.GetErrorMessage(code)
	api := ApiRestful{
		Code:    code,
		Message: message,
		Data:    map[string]interface{}{},
	}

	api.ApiRestful(w)
}

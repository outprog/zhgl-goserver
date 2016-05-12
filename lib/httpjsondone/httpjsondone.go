package httpjsondone

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
)

func GenRes(data []map[string]string, stat string, info string, template []map[string]string) []byte {
	if data == nil {
		data = []map[string]string{}
	}
	res := []map[string]string{}
	res = append(res, map[string]string{
		"stat": stat,
		"info": info,
	})
	if template == nil {
		data = []map[string]string{}
	}
	var m = map[string]([]map[string]string){
		"data":     data,
		"res":      res,
		"template": template,
	}
	json, _ := json.Marshal(m)
	return []byte(json)
}

func GetBody(r *http.Request) map[string]string {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return map[string]string{}
	}
	js, err := simplejson.NewJson(body)
	if err != nil {
		return map[string]string{}
	}
	res := map[string]string{}
	for key, value := range js.GetIndex(0).MustMap() {
		switch value := value.(type) {
		case string:
			res[key] = value
		}
	}

	return res
}

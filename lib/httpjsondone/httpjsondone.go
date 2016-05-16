package httpjsondone

import (
	"io/ioutil"
	"net/http"

	"github.com/outprog/go-simplejson"
)

func GenRes(data interface{}, res interface{}, template interface{}) []byte {

  genres, _:= simplejson.NewJson([]byte(`{}`))

	if data == nil {
		data = []map[string]string{}
	}
  genres.Set("data", data)

  if res == nil {
    res = map[string]string{}
  }
  genres.Set("res", res)

	if template == nil {
		template = map[string]string{}
	}
  genres.Set("template", template)

  by, _ := genres.Encode()
	return by
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
	for key, value := range js.MustMap() {
		switch value := value.(type) {
		case string:
			res[key] = value
		}
	}

	return res
}

package stdres

import (
	"encoding/json"
)

func Get(data []map[string]string, stat string, info string) []byte {
	res := []map[string]string{}
	res = append(res, map[string]string{
		"stat": stat,
		"info": info,
	})
	var m = map[string]([]map[string]string){
		"data": data,
		"res":  res,
	}
	json, _ := json.Marshal(m)
	return []byte(json)
}

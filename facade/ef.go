package facade

import "encoding/json"

// enhanced
// 带参数的门面模式，使用包导出方法
func Query(payload []byte) {
	var q queryer
	err := json.Unmarshal(payload, &q)
	if err != nil {
		return
	}
	q.Do()
}

func Insert(payload []byte) {
	var i inserter
	err := json.Unmarshal(payload, &i)
	if err != nil {
		return
	}
	i.Do()
}

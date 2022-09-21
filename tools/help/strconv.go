package help

import (
	"encoding/json"
	"gitlab.arksec.cn/wpf1118/toolbox/tools/logging"
	"strconv"
)

func StrToInt64(s string) int64 {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return int64(i)
}

func StrToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return 0
	}

	return f
}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return i
}

func BoolToStr(b bool) string {
	if b {
		return "true"
	}

	return "false"
}

func InterfaceToJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		logging.DebugF("translateParamsToJson error: %v", err)
		return ""
	}

	return string(b)
}

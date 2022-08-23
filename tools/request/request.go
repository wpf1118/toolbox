package request

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	defaultPage int64 = 1
	defaultSize int64 = 10
)

type ListReq struct {
	Search map[string]interface{} `json:"search"`
	Sort   map[string]interface{} `json:"sort"`
	Page   int64                  `json:"page"`
	Size   int64                  `json:"size"`
}

func NewListReq() *ListReq {
	return &ListReq{}
}

func (req *ListReq) Parse(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil && err.Error() != "EOF" {
		return err
	}

	if req.Page == 0 {
		req.Page = defaultPage
	}

	if req.Size == 0 {
		req.Size = defaultSize
	}

	return nil
}

func (req *ListReq) GetSearchString(key string) string {
	if v, ok := req.Search[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}

	return ""
}

func (req *ListReq) GetSearchInt(key string) int {
	if v, ok := req.Search[key]; ok {
		if i, ok := v.(int); ok {
			return i
		}

		if i, ok := v.(int64); ok {
			return int(i)
		}

		if i, ok := v.(int32); ok {
			return int(i)
		}

		if i, ok := v.(float64); ok {
			return int(i)
		}

		if i, ok := v.(float32); ok {
			return int(i)
		}
	}

	return 0
}

func (req *ListReq) GetSearchFloat(key string) float64 {
	if v, ok := req.Search[key]; ok {
		if i, ok := v.(float64); ok {
			return i
		}

		if i, ok := v.(int64); ok {
			return float64(i)
		}

		if i, ok := v.(int32); ok {
			return float64(i)
		}

		if i, ok := v.(int); ok {
			return float64(i)
		}

		if i, ok := v.(float32); ok {
			return float64(i)
		}

		if s, ok := v.(string); ok {
			f, _ := strconv.ParseFloat(s, 64)
			return f
		}
	}

	return 0
}

func (req *ListReq) GetSearchSliceString(key string) []string {
	if v, ok := req.Search[key]; ok {
		if i, ok := v.([]string); ok {
			return i
		}

		if list, ok := v.([]interface{}); ok {
			var result []string
			for _, vv := range list {
				result = append(result, vv.(string))
			}

			return result
		}
	}

	return nil
}

func (req *ListReq) GetSearchSliceInt(key string) []int {
	if v, ok := req.Search[key]; ok {
		if i, ok := v.([]int); ok {
			return i
		}

		if list, ok := v.([]interface{}); ok {
			var result []int
			for _, vv := range list {
				if s, ok := vv.(int); ok {
					result = append(result, s)
				}
				if s, ok := vv.(float64); ok {
					result = append(result, int(s))
				}
			}

			return result
		}
	}

	return nil
}

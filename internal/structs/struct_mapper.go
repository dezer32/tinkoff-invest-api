package structs

import (
	"encoding/json"
	"reflect"
)

//type StructMapper interface {
//	Map(dirty interface{}) (interface{}, error)
//}

type StructMapper struct {
	Rules map[string]interface{}
}

func NewStructMapper() *StructMapper {
	return &StructMapper{
		Rules: map[string]interface{}{
			"investapi.SharesResponse": SharesResponse{},
		},
	}
}

func (s *StructMapper) Map(dirty interface{}) (res interface{}, err error) {
	t := reflect.TypeOf(dirty)
	res, ok := s.Rules[t.String()]
	if ok != true {
		return dirty, nil
	}

	dirtyJson, err := json.Marshal(dirty)
	if err != nil {
		return
	}

	err = json.Unmarshal(dirtyJson, &res)
	if err != nil {
		return
	}

	return
}

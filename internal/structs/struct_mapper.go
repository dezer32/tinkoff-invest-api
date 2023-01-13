package structs

import (
	"encoding/json"
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

func (s *StructMapper) Map(dirty any, res any) (err error) {
	dirtyJson, err := json.Marshal(dirty)
	if err != nil {
		return
	}

	err = json.Unmarshal(dirtyJson, res)
	if err != nil {
		return
	}

	return
}

package services

import (
	"net/http"
	"encoding/json"
	"fmt"
	"errors"
)

type FieldPattern map[string](map[string]interface{})

type PostParams struct {
	Request    *http.Request
	Patterns   FieldPattern
	parsedData map[string]interface{}
	valid      bool
}

func (pp *PostParams)doValid(name *string, value *interface{}, con *string, conVal *interface{}) error {
	switch *con{
	case "required":
		if str, ok := (*value).(string); (*conVal).(bool) && (*value == nil || (ok && len(str) == 0)) {
			fmt.Printf(*name, *value, *con, *conVal)
			return errors.New("paramErr.validFail")
		}
	default:
		return errors.New("paramErr.validFail")
	}
	return nil
}

func (pp *PostParams)parse() {
	json.NewDecoder(pp.Request.Body).Decode(&pp.parsedData)
}

func (pp *PostParams) validParam(name string) error {
	conditions := pp.Patterns[name]
	value := pp.parsedData[name]
	for con, conVal := range conditions {
		if err := pp.doValid(&name, &value, &con, &conVal); err != nil {
			return err
		}
	}
	return nil
}

func (pp *PostParams)Valid() error {
	pp.parse()
	fmt.Println(pp.parsedData)
	for key := range pp.Patterns {
		if err := pp.validParam(key); err != nil {
			return err
		}
	}
	pp.valid = true
	return nil
}
func (pp *PostParams)Get(key string) interface{} {
	if pp.Patterns[key] != nil {
		return pp.parsedData[key]
	} else {
		return nil
	}

}
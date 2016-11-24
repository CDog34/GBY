package services

import (
	"net/http"
	"encoding/json"
	"fmt"
	"errors"
	"strings"
)

type FieldRule  map[string]interface{}
type FieldRules map[string]FieldRule

type PostParams struct {
	Request    *http.Request
	Rules      FieldRules
	parsedData map[string]interface{}
	valid      bool
}

func (pp *PostParams)doValid(name *string, value *interface{}, con *string, conVal *interface{}) error {
	switch *con{
	case "required":
		if str, ok := (*value).(string); (*conVal).(bool) && (*value == nil || (ok && len(str) == 0)) {
			return fmt.Errorf("paramErr.validFail.required/*/%s", *name)
		}
	case "type":
		switch (*value).(type){
		case string:
			if *conVal != "string" {
				return fmt.Errorf("paramErr.validFail.type/*/%s", *name)
			}
		case uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
			if *conVal != "number" {
				return fmt.Errorf("paramErr.validFail.type/*/%s", *name)
			}
		case bool:
			if *conVal != "boolean" {
				return fmt.Errorf("paramErr.validFail.type/*/%s", *name)
			}
		case nil:{
			return nil
		}
		default:
			fmt.Println(*name, *value)
			return errors.New("paramErr.validFail/*/Not Support Type")
		}
	default:
		return errors.New("paramErr.validFail/*/Not Support Rule")
	}
	return nil
}

func (pp *PostParams)parse() {
	json.NewDecoder(pp.Request.Body).Decode(&pp.parsedData)
}

func (pp *PostParams) validParam(name string) error {
	conditions := pp.Rules[name]
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
	for key := range pp.Rules {
		if err := pp.validParam(key); err != nil {
			return err
		}
	}
	pp.valid = true
	return nil
}
func (pp *PostParams)GetString(key string) string {
	if str, ok := pp.parsedData[key].(string); ok && pp.Rules[key] != nil {
		return strings.Trim(str, " ")
	} else {
		return ""
	}
}
func (pp *PostParams)GetBool(key string) bool {
	if val, ok := pp.parsedData[key].(bool); ok && pp.Rules[key] != nil {
		return val
	} else {
		return false
	}
}
func (pp *PostParams)GetInt(key string) int {
	if val, ok := pp.parsedData[key].(int); ok && pp.Rules[key] != nil {
		return val
	} else {
		return 0
	}
}
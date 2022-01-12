package validate

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

var (
	JsonType   = "json"
	InputType  = "input"
	StructType = "struct"
)

func Do(opts govalidator.Options, validateType string) (url.Values, map[string]interface{}) {
	data := make(map[string]interface{})
	opts.Data = &data

	v := govalidator.New(opts)

	var e url.Values
	switch {
	case validateType == JsonType:
		e = v.ValidateJSON()
	case validateType == InputType:
		e = v.Validate()
	case validateType == StructType:
		e = v.ValidateStruct()
	}

	return e, data
}

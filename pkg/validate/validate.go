package validate

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

type Validate struct{}

var (
	JsonType   = "json"
	InputType  = "input"
	StructType = "struct"
)

func (Validate) do(opts govalidator.Options, validateType string) (url.Values, map[string]interface{}) {
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

func (v Validate) SuppressValidate(r echo.Context) (url.Values, map[string]interface{}) {
	rules := govalidator.MapData{
		"sender": []string{"required", "between:4,20"},
	}

	opts := govalidator.Options{
		Request:         r.Request(), // request object
		Rules:           rules,       // rules map
		RequiredDefault: true,        // all the field to be pass the rules
	}

	return v.do(opts, InputType)
}

func (v Validate) PublishValidate(r echo.Context) (url.Values, map[string]interface{}) {
	rules := govalidator.MapData{
		"from":    []string{"required", "between:4,20"},
		"to":      []string{"required", "between:4,20"},
		"message": []string{"between:0,250"},
	}

	opts := govalidator.Options{
		Request:         r.Request(), // request object
		Rules:           rules,       // rules map
		RequiredDefault: true,        // all the field to be pass the rules
	}

	return v.do(opts, JsonType)
}

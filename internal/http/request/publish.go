package request

import (
	"Telegraph/pkg/validate"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

func PublishValidate(r echo.Context) (url.Values, map[string]interface{}) {
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

	return validate.Do(opts)
}

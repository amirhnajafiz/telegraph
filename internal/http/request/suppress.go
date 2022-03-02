package request

import (
	"net/url"

	"github.com/amirhnajafiz/Telegraph/pkg/validate"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

func SuppressValidate(r echo.Context) (url.Values, map[string]interface{}) {
	rules := govalidator.MapData{
		"sender": []string{"required", "between:4,20"},
	}

	opts := govalidator.Options{
		Request:         r.Request(), // request object
		Rules:           rules,       // rules map
		RequiredDefault: true,        // all the field to be pass the rules
	}

	return validate.Do(opts, validate.InputType)
}

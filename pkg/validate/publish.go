package validate

import (
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

func ValidatePublish(r echo.Context) map[string]interface{} {
	rules := govalidator.MapData{
		"from":    []string{"between:4,20"},
		"to":      []string{"between:4,20"},
		"message": []string{"between:0,250"},
	}

	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request:         r.Request(), // request object
		Rules:           rules,       // rules map
		Data:            &data,
		RequiredDefault: true, // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}

	return err
}

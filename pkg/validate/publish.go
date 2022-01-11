package validate

import (
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

func ValidatePublish(r echo.Context) map[string]interface{} {
	rules := govalidator.MapData{
		"from":    []string{"required", "min:4", "max:20"},
		"to":      []string{"required", "min:4", "max:20"},
		"message": []string{"max:250"},
	}

	messages := govalidator.MapData{
		"from":    []string{"required: Enter the source name", "min: Source name is too short", "max: Source name is too long"},
		"to":      []string{"required: Enter the destin name", "min: Destin name is too short", "max: Destin name is too long"},
		"message": []string{"max: Input value is to large"},
	}

	opts := govalidator.Options{
		Request:         r.Request(), // request object
		Rules:           rules,       // rules map
		Messages:        messages,    // custom message map (Optional)
		RequiredDefault: true,        // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.Validate()
	err := map[string]interface{}{"validationError": e}

	return err
}

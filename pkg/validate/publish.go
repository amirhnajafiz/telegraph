package validate

import (
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

func ValidatePublish(r echo.Context) map[string]interface{} {
	rules := govalidator.MapData{
		"username": []string{"required", "between:3,8"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
		"phone":    []string{"digits:11"},
		"agree":    []string{"bool"},
		"dob":      []string{"date"},
	}

	messages := govalidator.MapData{
		"username": []string{"required:আপনাকে অবশ্যই ইউজারনেম দিতে হবে", "between:ইউজারনেম অবশ্যই ৩-৮ অক্ষর হতে হবে"},
		"phone":    []string{"digits:ফোন নাম্বার অবশ্যই ১১ নম্বারের হতে হবে"},
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

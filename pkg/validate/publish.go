package validate

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

func Do(opts govalidator.Options) (url.Values, map[string]interface{}) {
	data := make(map[string]interface{})
	opts.Data = data

	v := govalidator.New(opts)
	e := v.ValidateJSON()

	return e, data
}

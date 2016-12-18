package gf

import 	(
		"html"
		"reflect"
		"strconv"
		"strings"
		//
		)

const 	(
		STRING_MAX_LENGTH = 2000
		)

type ValidationFunction func (*Request, string) (bool, interface{})

type ValidationConfig struct {
	model interface{}
	function ValidationFunction
	keys []string
	min int
	max int
}

func (vc *ValidationConfig) Keys() string {

	return strings.Join(vc.keys, "_")
}

func (vc *ValidationConfig) Type() string {

	return reflect.TypeOf(vc.model).String()
}

func (vc *ValidationConfig) Expecting() string {

	return "expecting: " + vc.Type() + " for keys: "+strings.Join(vc.keys, ", ")
}

func NewValidationConfig(validationType interface{}, validationFunc ValidationFunction) *ValidationConfig {

	return &ValidationConfig{
		model: validationType,
		function: validationFunc,
	}
}

// Returns a validation function that checks for a string with a length within optional range
func String(ranges ...int) *ValidationConfig {

	var min, max int

	switch len(ranges) {

		case 0:

			max = STRING_MAX_LENGTH

		case 1:

			min = ranges[0]
			max = ranges[0]

		case 2:

			min = ranges[0]
			max = ranges[1]

	}

	config := NewValidationConfig(
		"",
		func (req *Request, param string) (bool, interface{}) {

			lp := len(param)

			if lp == 0 || lp > 64 { return false, nil }

			return true, req.BlueMonday.Sanitize(html.UnescapeString(param))
		},
	)

	config.min = min
	config.max = max

	return config
}

func SplitString(delimiter string) *ValidationConfig {

	return NewValidationConfig(
		[]string{},
		func (req *Request, param string) (bool, interface{}) {

			lp := len(param)

			if lp == 0 || lp > STRING_MAX_LENGTH { return false, nil }
			
			list := []string{}

			for _, part := range strings.Split(req.BlueMonday.Sanitize(html.UnescapeString(param)), delimiter) {

				if len(part) == 0 { continue }

				list = append(list, part)

			}

			return true, list
		},
	)
}

// Returns a validation function that checks for an int which parses correctly
func Int() *ValidationConfig {

	return NewValidationConfig(
		0,
		func (req *Request, param string) (bool, interface{}) {

			if len(param) == 0 { return false, nil }

			val, err := strconv.Atoi(param)

			return err == nil, val
		},
	)
}

func Int64() *ValidationConfig {

	return NewValidationConfig(
		int64(0),
		func (req *Request, param string) (bool, interface{}) {

			if len(param) == 0 { return false, nil }

			val, err := strconv.ParseInt(param, 10, 64)

			return err == nil, val
		},
	)
}

func Float64() *ValidationConfig {

	return NewValidationConfig(
		float64(0),
		func (req *Request, param string) (bool, interface{}) {

			if len(param) == 0 { return false, nil }

			val, err := strconv.ParseFloat(param, 64)

			return err == nil, val
		},
	)
}

func Bool() *ValidationConfig {

	return NewValidationConfig(
		true,
		func (req *Request, param string) (bool, interface{}) {

			if len(param) == 0 { return false, nil }

			switch param {

				case "true":	return true, true
				case "false":	return true, false

			}

			return false, false
		},
	)
}

//

func CountryISO2() *ValidationConfig {
		
	return NewValidationConfig(
		&Country{},
		func (req *Request, param string) (bool, interface{}) {

			lp := len(param)

			if lp == 0 || lp > 64 { return false, nil }

			param = strings.ToUpper(param)

			country := gfConfig.countries[param]

			return (country != nil), country
		},
	)
}

func LanguageISO2() *ValidationConfig {
		
	return NewValidationConfig(
		&Language{},
		func (req *Request, param string) (bool, interface{}) {

			lp := len(param)

			if lp == 0 || lp > 64 { return false, nil }

			param = strings.ToUpper(param)

			language := gfConfig.languages[param]

			return (language != nil), language
		},
	)
}
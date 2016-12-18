package gf

import 	(
		"reflect"
		"strconv"
		"strings"
		//
		)

const 	(
		STRING_MAX_LENGTH = 2000
		)

type BodyValidationFunction func (RequestInterface, interface{}) (bool, interface{})
type PathValidationFunction func (RequestInterface, string) (bool, interface{})

type ValidationConfig struct {
	model interface{}
	pathFunction PathValidationFunction
	bodyFunction BodyValidationFunction
	keys []string
	min int
	max int
}

func (vc *ValidationConfig) Key() string {

	return vc.keys[0]
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

func NewValidationConfig(validationType interface{}, pathFunction PathValidationFunction, bodyFunction BodyValidationFunction) *ValidationConfig {

	return &ValidationConfig{
		model: validationType,
		pathFunction: pathFunction,
		bodyFunction: bodyFunction,
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
		func (req RequestInterface, param string) (bool, interface{}) {

			lp := len(param)

			if lp < min || lp > max { return false, nil }

			return true, globalNode.Config.Sanitize(param)
		},
		func (req RequestInterface, param interface{}) (bool, interface{}) {

			s, ok := param.(string); if !ok { return false, nil }

			lp := len(s)

			if lp < min || lp > max { return false, nil }

			return true, globalNode.Config.Sanitize(s)
		},
	)

	config.min = min
	config.max = max

	return config
}

func SplitString(delimiter string) *ValidationConfig {

	return NewValidationConfig(
		[]string{},
		func (req RequestInterface, param string) (bool, interface{}) {

			lp := len(param)

			if lp == 0 || lp > STRING_MAX_LENGTH { return false, nil }
			list := []string{}

			for _, part := range strings.Split(globalNode.Config.Sanitize(param), delimiter) {

				if len(part) == 0 { continue }

				list = append(list, part)

			}

			return true, list
		},
		func (req RequestInterface, param interface{}) (bool, interface{}) {

			s, ok := param.(string); if !ok { return false, nil }

			lp := len(s)
			if lp == 0 || lp > STRING_MAX_LENGTH { return false, nil }
			
			list := []string{}

			for _, part := range strings.Split(globalNode.Config.Sanitize(s), delimiter) {

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
		func (req RequestInterface, param string) (bool, interface{}) {

			if len(param) == 0 { return false, nil }

			val, err := strconv.Atoi(param)

			return err == nil, val
		},
		func (req RequestInterface, param interface{}) (bool, interface{}) {

			i, ok := param.(float64)

			return ok, int(i)
		},
	)
}

func Int64() *ValidationConfig {

	return NewValidationConfig(
		int64(0),
		func (req RequestInterface, param string) (bool, interface{}) {

			if len(param) == 0 { return false, nil }

			val, err := strconv.ParseInt(param, 10, 64)

			return err == nil, val
		},
		func (req RequestInterface, param interface{}) (bool, interface{}) {

			i, ok := param.(float64)

			return ok, int64(i)
		},
	)
}

func Float64() *ValidationConfig {

	return NewValidationConfig(
		float64(0),
		func (req RequestInterface, param string) (bool, interface{}) {

			if len(param) == 0 { return false, nil }

			val, err := strconv.ParseFloat(param, 64)

			return err == nil, val
		},
		func (req RequestInterface, param interface{}) (bool, interface{}) {

			i, ok := param.(float64)

			return ok, i
		},
	)
}

func Bool() *ValidationConfig {

	return NewValidationConfig(
		true,
		func (req RequestInterface, param string) (bool, interface{}) {

			if len(param) == 0 { return false, nil }

			switch param {

				case "true":	return true, true
				case "false":	return true, false

			}

			return false, false
		},
		func (req RequestInterface, param interface{}) (bool, interface{}) {

			b, ok := param.(bool); if !ok { return false, nil }

			return false, b
		},
	)
}

func MSI() *ValidationConfig {

	return NewValidationConfig(
		true,
		nil,
		func (req RequestInterface, param interface{}) (bool, interface{}) {

			x, ok := param.(Object); if !ok { return false, nil }

			return false, x
		},
	)
}

func IA() *ValidationConfig {

	return NewValidationConfig(
		true,
		nil,
		func (req RequestInterface, param interface{}) (bool, interface{}) {

			x, ok := param.(Array); if !ok { return false, nil }

			return false, x
		},
	)
}

//

func CountryISO2() *ValidationConfig {
		
	return NewValidationConfig(
		&Country{},
		func (req RequestInterface, param string) (bool, interface{}) {

			lp := len(param)
			if lp == 0 || lp > 64 { return false, nil }

			param = strings.ToUpper(param)

			country := globalNode.Config.countries[param]

			return (country != nil), country
		},
		func (req RequestInterface, param interface{}) (bool, interface{}) {

			s, ok := param.(string); if !ok { return false, nil }

			lp := len(s)
			if lp > 64 { return false, nil }

			country := globalNode.Config.countries[strings.ToUpper(s)]

			return (country != nil), country
		},
	)
}

func LanguageISO2() *ValidationConfig {
		
	return NewValidationConfig(
		&Language{},
		func (req RequestInterface, param string) (bool, interface{}) {

			lp := len(param)
			if lp == 0 || lp > 64 { return false, nil }

			param = strings.ToUpper(param)

			language := globalNode.Config.languages[param]

			return (language != nil), language
		},
		func (req RequestInterface, param interface{}) (bool, interface{}) {

			s, ok := param.(string); if !ok { return false, nil }

			lp := len(s)
			if lp > 64 { return false, nil }

			language := globalNode.Config.languages[strings.ToUpper(s)]

			return (language != nil), language
		},
	)
}
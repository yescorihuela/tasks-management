package validator

import "fmt"

type Validator struct {
	Errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) In(value string, list []string) bool {
	fmt.Println(list)
	for _, v := range list {
		fmt.Println("value == v", value, v)
		if value == v {
			return true
		}
	}
	return false
}

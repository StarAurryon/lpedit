package model

import "github.com/StarAurryon/lpedit-lib/model/pod"

type Parameter struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	Value         string   `json:"value"`
	ValueNumber   float32  `json:"valueNumber"`
	AllowedValues []string `json:"allowedValue"`

	Min int `json:"min"`
	Max int `json:"max"`
}

func (Parameter) WailsTsType() Parameter { return Parameter{} }

func ToParameters(src []pod.Parameter) []Parameter {
	ret := make([]Parameter, len(src))
	for i := range ret {
		ret[i] = ToParameter(src[i])
	}
	return ret
}

func ToParameter(src pod.Parameter) Parameter {
	min, max := src.GetValueRange()
	ret := Parameter{
		Id:            int(src.GetID()),
		Name:          src.GetName(),
		Value:         src.GetValueCurrent(),
		ValueNumber:   src.GetValueCurrent2(),
		AllowedValues: src.GetAllowedValues(),
		Min:           min,
		Max:           max,
	}

	return ret
}

type ListParam struct {
	List  []string `json:"list"`
	Value string   `json:"value"`
}

type RangeParam struct {
	Value    string `json:"value"`
	ValueMin string `json:"valueMin"`
	ValueMax string `json:"valueMax"`
}

package model

import "github.com/StarAurryon/lpedit-lib/model/pod"

type Set struct {
	Id      int                      `json:"id"`
	Name    string                   `json:"name"`
	Presets [pod.PresetPerSet]Preset `json:"presets"`
}

func (Set) WailsTsType() Set { return Set{} }

func ToSet(src *pod.Set) Set {
	return Set{
		Id:      int(src.GetID()),
		Name:    src.GetName(),
		Presets: ToPresets(src),
	}
}

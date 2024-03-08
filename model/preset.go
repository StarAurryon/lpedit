package model

import "github.com/StarAurryon/lpedit-lib/model/pod"

type Preset struct {
	Id        int                `json:"id"`
	IdStr     string             `json:"idStr"`
	DTs       [2]DT              `json:"dts"`
	Items     [12]PedalBoardItem `json:"items"`
	Name      string             `json:"name"`
	Parameter []Parameter        `json:"parameters"`
	SetId     int                `json:"setId"`
}

func (Preset) WailsTsType() Preset { return Preset{} }

func ToPresets(src *pod.Set) (ret [pod.PresetPerSet]Preset) {
	for i := range ret {
		preset := src.GetPreset(uint8(i))
		if preset == nil {
			continue
		}
		ret[i] = ToPreset(preset)
	}
	return ret
}

func ToPreset(src *pod.Preset) Preset {
	return Preset{
		Id:        int(src.GetID()),
		IdStr:     src.GetID2(),
		DTs:       [2]DT{ToDT(src.GetDT(0)), ToDT(src.GetDT(1))},
		Name:      src.GetName(),
		Parameter: ToParameters(src.GetParams()),
		SetId:     int(src.GetSet().GetID()),
	}
}

package model

import "github.com/StarAurryon/lpedit-lib/model/pod"

type Pod struct {
	CurrentPresetId *int               `json:"currentPresetId"`
	CurrentSetId    *int               `json:"currentSetId"`
	Sets            [pod.NumberSet]Set `json:"sets"`
}

func (Pod) WaiWailsTsType() Pod { return Pod{} }

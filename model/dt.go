package model

import "github.com/StarAurryon/lpedit-lib/model/pod"

type DT struct {
	Id       int    `json:"id"`
	AmpId    int    `json:"ampId"`
	Class    string `json:"class"`
	Mode     string `json:"mode"`
	Topology string `json:"topology"`
}

func (DT) WailsTsType() DT { return DT{} }

func ToDT(src *pod.DT) DT {
	return DT{
		Id:       src.GetID(),
		AmpId:    int(src.GetAmpID()),
		Class:    src.GetClass(),
		Mode:     src.GetMode(),
		Topology: src.GetTopology(),
	}
}

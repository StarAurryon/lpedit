package model

type PedalBoardItem struct {
	Active bool `json:"active"`
}

func (PedalBoardItem) WailsTsType() PedalBoardItem { return PedalBoardItem{} }

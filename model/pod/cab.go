/*
 * Copyright (C) 2020 Nicolas SCHWARTZ
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301, USA
 */

package pod

import "fmt"

const (
    noCab uint32 = 17301503
    CabLowCutID uint32 = iota
    CabResLevelID
    CabThumpID
    CabDecayID
    CabERID
    CabMicID
)

type Cab struct {
    id          uint32
    active      bool
    ctype       uint32
    name        string
    params      []Parameter
    pos         uint16
    posType     uint8
    pb          *PedalBoard
    hideParams  bool
}

var cabs = []Cab {
    Cab{ctype: noCab, active: true, name: "No Cab", hideParams: true, params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235968, active: true, name: "212 PhD Ported", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235969, active: true, name: "6x9 Super O", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235970, active: true, name: "112 Celest 12-H", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235971, active: true, name: "112 BF 'Lux", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235972, active: true, name: "112 Field Coil", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235973, active: true, name: "112 Blue Bell", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235974, active: true, name: "212 Blackface Double", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235976, active: true, name: "212 Silver Bell", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235977, active: true, name: "410 Tweed", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235978, active: true, name: "412 Uber", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235979, active: true, name: "412 XXL V-30", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235980, active: true, name: "412 Hiway", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235982, active: true, name: "412 Greenback 25", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235983, active: true, name: "412 Brit T-75", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235984, active: true, name: "412 Tread V-30", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235985, active: true, name: "412 Greenback 30", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
    Cab{ctype: 17235986, active: true, name: "115 Flip Top", params: []Parameter{
        &PerCentParam{GenericParameter: GenericParameter{id: CabERID, name: "E.R."}},
        &FreqParam{GenericParameter: GenericParameter{id: CabLowCutID, name: "Low Cut"}, min: 20, max: 500},
        &PerCentParam{GenericParameter: GenericParameter{id: CabResLevelID, name: "Res Level"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabThumpID, name: "Thump"}},
        &PerCentParam{GenericParameter: GenericParameter{id: CabDecayID, name: "Decay"}},
        &ListParam{GenericParameter: GenericParameter{id: CabMicID, name: "Mic Model"}, binValueType: Int32Type,
            list: []string{"57 On XS", "57 Off XS", "409 Dyn", "421 Dyn", "4038 Rbn",
                "121 Rbn", "67 Cond", "87 Cond"}},
        }},
}

func newNoCab(id uint32, pos uint16, posType uint8, pb *PedalBoard) *Cab {
    c := newCab(id, pos, posType, pb, noCab)
    return c
}

func newCab(id uint32, pos uint16, posType uint8, pb *PedalBoard, ctype uint32) *Cab {
    for _, cab := range cabs {
        if cab.ctype == ctype {
            newCab := new(Cab)
            *newCab = cab
            newCab.id = id
            newCab.pos = pos
            newCab.posType = posType
            newCab.pb = pb
            newCab.params = make([]Parameter, len(cab.params))
            for i := range newCab.params {
                newCab.params[i] = cab.params[i].Copy()
                newCab.params[i].SetParent(newCab)
            }
            return newCab
        }
    }
    return nil
}

func GetCabType() []string {
    m := make([]string, len(cabs))
    for i, a := range cabs {
        m[i] = a.GetName()
    }
    return m
}

func (c *Cab) GetActive() bool {
    return c.active
}

func (c *Cab) GetActive2() uint32 {
    if c.active {
        return 1
    }
    return 0
}

func (c *Cab) GetHideParams() bool {
    return c.hideParams
}

func (c *Cab) GetID() uint32 {
    return c.id
}

func (c *Cab) GetName() string {
    return c.name
}

func (c *Cab) GetParam(id uint32) Parameter {
    for _, param := range c.params {
        if param.GetID() == id {
            return param
        }
    }
    return nil
}

func (c *Cab) GetParams() []Parameter {
    return c.params
}

func (c *Cab) GetParamLen() uint16 {
    return uint16(len(c.params))
}

func (c *Cab) GetPos() (uint16, uint8) {
    return c.pos, c.posType
}

func (c *Cab) GetType() uint32 {
    return c.ctype
}

func (c *Cab) LockData() { c.pb.LockData() }

func (c *Cab) SetActive(active bool){
    c.active = active
}

func (c *Cab) SetPos(pos uint16, posType uint8) {}

func (c *Cab) SetPosWithoutCheck(pos uint16, posType uint8) {
    c.pos = pos
    c.posType = posType
}

func (c *Cab) SetType(ctype uint32) error{
    _c := newCab(c.id, c.pos, c.posType, c.pb, ctype)
    if _c == nil {
        return fmt.Errorf("Cab type not found, code: %d", ctype)
    }
    *c = *_c
    for i := range c.params {
        c.params[i].SetParent(c)
    }
    return nil
}

func (c *Cab) SetType2(name string, none string) {
    for _, _c := range cabs {
        if name == _c.name {
            c.SetType(_c.ctype)
            break
        }
    }
}

func (c *Cab) UnlockData() { c.pb.UnlockData() }

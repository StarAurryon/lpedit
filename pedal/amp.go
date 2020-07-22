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

package pedal

import "fmt"

const (
    disAmpType uint32 = 524287
)

type Amp struct {
    id     uint32
    active bool
    name   string
    params []Parameter
    pb     *PedalBoard
    plist  *[]PedalBoardItem
    cab    *Cab
}

var amps = map[uint32]Amp {
    disAmpType: Amp{active: true, name: "Amp Disabled"},
    458752: Amp{active: true, name: "phD Motorway",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458753: Amp{active: true, name: "Tweed B-Man Normal",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458754: Amp{active: true, name: "Tweed B-Man Bright",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458755: Amp{active: true, name: "Blackface ‘Lux Normal",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458756: Amp{active: true, name: "Blackface ‘Lux Vibrato",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458757: Amp{active: true, name: "Blackface Double Normal",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458758: Amp{active: true, name: "Blackface Double Vibrato",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458760: Amp{active: true, name: "Hiway 100",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458761: Amp{active: true, name: "Brit J-45 Normal",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458762: Amp{active: true, name: "Brit J-45 Bright",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458763: Amp{active: true, name: "Treadplate",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458764: Amp{active: true, name: "Brit P-75 Normal",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458765: Amp{active: true, name: "Brit P-75 Bright",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458766: Amp{active: true, name: "Super O",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458768: Amp{active: true, name: "Class A-15",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458769: Amp{active: true, name: "Class A-30 TB",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458770: Amp{active: true, name: "Divide 9/15",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458772: Amp{active: true, name: "Gibtone 185",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458773: Amp{active: true, name: "Brit J-800",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458775: Amp{active: true, name: "Bomber Uber",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458777: Amp{active: true, name: "Angel F-Ball",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458779: Amp{active: true, name: "phD Motorway Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458780: Amp{active: true, name: "Tweed B-Man Normal Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458781: Amp{active: true, name: "Tweed B-Man Bright Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458782: Amp{active: true, name: "Blackface ‘Lux Normal Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458783: Amp{active: true, name: "Blackface ‘Lux Vibrato Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458784: Amp{active: true, name: "Blackface Double Normal Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458785: Amp{active: true, name: "Blackface Double Vibrato Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458787: Amp{active: true, name: "Hiway 100 Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458788: Amp{active: true, name: "Brit J-45 Normal Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458789: Amp{active: true, name: "Brit J-45 Bright Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458790: Amp{active: true, name: "Treadplate Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458791: Amp{active: true, name: "Brit P-75 Normal Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458792: Amp{active: true, name: "Brit P-75 Bright Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458793: Amp{active: true, name: "Super O Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458795: Amp{active: true, name: "Class A-15 Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458796: Amp{active: true, name: "Class A-30 TB Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458797: Amp{active: true, name: "Divide 9/15 Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458799: Amp{active: true, name: "Gibtone 185 Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458800: Amp{active: true, name: "Brit J-800 Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458802: Amp{active: true, name: "Bomber Uber Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458804: Amp{active: true, name: "Angel F-Ball Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458834: Amp{active: true, name: "Line 6 Elektrik",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458835: Amp{active: true, name: "Line 6 Elektrik Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458837: Amp{active: true, name: "Plexi Lead 100 Normal",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458838: Amp{active: true, name: "Plexi Lead 100 Normal Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458840: Amp{active: true, name: "Plexi Lead 100 Bright",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458841: Amp{active: true, name: "Plexi Lead 100 Bright Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458843: Amp{active: true, name: "Flip Top",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458844: Amp{active: true, name: "Flip Top Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458846: Amp{active: true, name: "Solo 100 Clean",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458847: Amp{active: true, name: "Solo 100 Clean Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458849: Amp{active: true, name: "Solo 100 Crunch",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458850: Amp{active: true, name: "Solo 100 Crunch Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458852: Amp{active: true, name: "Solo 100 OD",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458853: Amp{active: true, name: "Solo 100 OD Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458855: Amp{active: true, name: "Line 6 Doom",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458856: Amp{active: true, name: "Line 6 Doom Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458858: Amp{active: true, name: "Line 6 Epic",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
    458859: Amp{active: true, name: "Line 6 Epic Preamp",
        params: []Parameter{
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Presence", value:0},
            Parameter{name: "Volume", value:0},
            }},
}

func newDisAmp(id uint32, pb *PedalBoard, plist *[]PedalBoardItem, c *Cab) *Amp {
    a := amps[disAmpType]
    a.id = id
    a.pb = pb
    a.cab = c
    *plist = append(*plist, PedalBoardItem(&a))
    a.plist = plist
    return &a
}

func (a *Amp) GetActive() bool {
    return a.active
}

func (a *Amp) GetID() uint32 {
    return a.id
}

func (a *Amp) GetName() string {
    return a.name
}

func (a *Amp) GetParam(id uint16) *Parameter {
    if id >= uint16(len(a.params)) {
        return nil
    }
    return &a.params[id]
}

func (a *Amp) GetParamLen() uint16 {
    return uint16(len(a.params))
}

func (a *Amp) SetActive(active bool){
    a.active = active
}

func (a *Amp) SetLastPos(pos uint16, ctype uint8) error {
    if a.GetID() == 2 {
        return nil
    }
    b, _ := a.pb.GetItem(2).(*Amp)
    switch pos {
    case 0:
        a.remove()
        a.pb.pchan.aAmp = append(a.pb.pchan.aAmp, a)
        a.plist = &a.pb.pchan.aAmp
        b.remove()
        b.pb.pchan.bAmp = append(b.pb.pchan.bAmp, b)
        b.plist = &b.pb.pchan.bAmp
    default:
        p, _ := a.pb.GetPedal(pos).(*Pedal)
        switch p.plist {
        case &a.pb.start:
            a.remove()
            a.pb.startAmp = append(a.pb.startAmp, a)
            a.plist = &a.pb.startAmp
        case &a.pb.end:
            a.remove()
            a.pb.endAmp = append(a.pb.endAmp, a)
            a.plist = &a.pb.endAmp
        default:
            return fmt.Errorf("Wrong metodology in Amp placement")
        }
        b.remove()
        b.pb.bAmp = append(b.pb.bAmp, b)
        b.plist = &a.pb.bAmp
    }
    return nil
}

func (a *Amp) SetType(atype uint32) error{
    _a := *a
    at, found := amps[atype]
    if !found {
        return fmt.Errorf("Amp type not found, code: %d", atype)
    }
    *a = at
    a.id = _a.id
    a.pb = _a.pb
    a.cab = _a.cab
    a.plist = _a.plist
    return nil
}

func (a *Amp) remove() {
    *a.plist = nil
}

func (a Amp) PrintInfo() {
    a.cab.PrintInfo()
    fmt.Printf("Id %d, Amp Info, Name %s, Active %t\n", a.id, a.name, a.active)
    fmt.Printf("Parameters:\n")
    for i, param := range(a.params) {
        fmt.Printf("----%d %s %f\n", i, param.name, param.value)
    }
}

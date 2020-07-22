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
    noCabType uint32 = 17301503
)

type Cab struct {
    id     uint32
    active bool
    name   string
    pb     *PedalBoard
}

var cabs = map[uint32]Cab {
    noCabType: Cab{active: true, name: "No Cab"},
    17235968: Cab{active: true, name: "212 PhD Ported"},
    17235969: Cab{active: true, name: "6x9 Super O"},
    17235970: Cab{active: true, name: "112 Celest 12-H"},
    17235971: Cab{active: true, name: "112 BF 'Lux"},
    17235972: Cab{active: true, name: "112 Field Coil"},
    17235973: Cab{active: true, name: "112 Blue Bell"},
    17235974: Cab{active: true, name: "212 Blackface Double"},
    17235976: Cab{active: true, name: "212 Silver Bell"},
    17235977: Cab{active: true, name: "410 Tweed"},
    17235978: Cab{active: true, name: "412 Uber"},
    17235979: Cab{active: true, name: "412 XXL V-30"},
    17235980: Cab{active: true, name: "412 Hiway"},
    17235982: Cab{active: true, name: "412 Greenback 25"},
    17235983: Cab{active: true, name: "412 Brit T-75"},
    17235984: Cab{active: true, name: "412 Tread V-30"},
    17235985: Cab{active: true, name: "412 Greenback 30"},
    17235986: Cab{active: true, name: "115 Flip Top"},
}

func newNoneCab(id uint32, pb *PedalBoard, plist *[]PedalBoardItem) *Cab {
    c := cabs[noCabType]
    c.id = id
    c.pb = pb
    *plist = append(*plist, PedalBoardItem(&c))
    return &c
}

func (c *Cab) GetActive() bool {
    return c.active
}

func (c *Cab) GetID() uint32 {
    return c.id
}

func (c *Cab) GetName() string {
    return c.name
}

func (c *Cab) GetParam(id uint16) *Parameter {
    return nil
}

func (c *Cab) GetParamLen() uint16 {
    return 0
}

func (c *Cab) SetActive(active bool){
    c.active = active
}

func (c *Cab) SetLastPos(pos uint16, ctype uint8) error {
    return nil
}

func (c *Cab) SetType(ctype uint32) error{
    _c := *c
    ct, found := cabs[ctype]
    if !found {
        return fmt.Errorf("Cab type not found, code: %d", ctype)
    }
    *c = ct
    c.id = _c.id
    c.pb = _c.pb
    return nil
}

func (c *Cab) remove() {}

func (c Cab) PrintInfo() {
    fmt.Printf("Id %d, Cab Info, Name %s, Active %t\n", c.id, c.name, c.active)
}

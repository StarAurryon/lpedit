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
import "log"

type Cab struct {
    id     uint32
    active bool
    ctype  uint32
    name   string
    pb     *PedalBoard
}

var cabs = []Cab {
    Cab{ctype: 17301503, active: true, name: "No Cab"},
    Cab{ctype: 17235968, active: true, name: "212 PhD Ported"},
    Cab{ctype: 17235969, active: true, name: "6x9 Super O"},
    Cab{ctype: 17235970, active: true, name: "112 Celest 12-H"},
    Cab{ctype: 17235971, active: true, name: "112 BF 'Lux"},
    Cab{ctype: 17235972, active: true, name: "112 Field Coil"},
    Cab{ctype: 17235973, active: true, name: "112 Blue Bell"},
    Cab{ctype: 17235974, active: true, name: "212 Blackface Double"},
    Cab{ctype: 17235976, active: true, name: "212 Silver Bell"},
    Cab{ctype: 17235977, active: true, name: "410 Tweed"},
    Cab{ctype: 17235978, active: true, name: "412 Uber"},
    Cab{ctype: 17235979, active: true, name: "412 XXL V-30"},
    Cab{ctype: 17235980, active: true, name: "412 Hiway"},
    Cab{ctype: 17235982, active: true, name: "412 Greenback 25"},
    Cab{ctype: 17235983, active: true, name: "412 Brit T-75"},
    Cab{ctype: 17235984, active: true, name: "412 Tread V-30"},
    Cab{ctype: 17235985, active: true, name: "412 Greenback 30"},
    Cab{ctype: 17235986, active: true, name: "115 Flip Top"},
}

func newNoneCab(id uint32, pb *PedalBoard, plist *[]PedalBoardItem) *Cab {
    c := cabs[0]
    c.id = id
    c.pb = pb
    *plist = append(*plist, PedalBoardItem(&c))
    return &c
}

func newCab(ctype uint32) *Cab {
    for _, newCab := range cabs {
        if newCab.ctype == ctype {
            return &newCab
        }
    }
    return nil
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

func (c *Cab) GetParam(id uint16) Parameter {
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
    _c := newCab(ctype)
    if _c == nil {
        return fmt.Errorf("Cab type not found, code: %d", ctype)
    }
    _c.id = c.id
    _c.pb = c.pb
    *c = *_c
    return nil
}

func (c *Cab) remove() {}

func (c Cab) LogInfo() {
    log.Printf("Id %d, Cab Info, Name %s, Active %t\n", c.id, c.name, c.active)
}

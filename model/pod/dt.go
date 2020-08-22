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

type DT struct {
    id       int
    ampID    uint32
    pb       *PedalBoard
    class    bool
    mode     bool
    topology uint8
}

var allowedTopology = []string{"I", "II", "III", "IV"}

func newDT(id int, ampID uint32, pb *PedalBoard) *DT {
    var dt DT
    dt.id = id
    dt.ampID = ampID
    dt.pb = pb
    return &dt
}

func GetAllowedTopology() []string {
    return allowedTopology
}

func (dt *DT) GetAmpID() uint32 {
    return dt.ampID
}

func (dt *DT) GetBinClass() uint8 {
    if dt.class {
        return 0x7f
    }
    return 0
}

func (dt *DT) GetBinMode() uint8 {
    if dt.mode {
        return 0x7f
    }
    return 0
}

func (dt *DT) GetBinTopology() uint8 {
    return dt.topology
}

func (dt *DT) GetClass() string {
    if dt.class {
        return "A/B"
    }
    return "A"
}

func (dt *DT) GetID() int {
    return dt.id
}

func (dt *DT) GetMode() string {
    if dt.mode {
        return "Pent"
    }
    return "Tri"
}

func (dt *DT) GetTopology() string {
    return allowedTopology[dt.topology]
}

func (dt *DT) SetBinClass(data byte) error {
    switch data {
    case 0:
        dt.class = false
    case 0x7f:
        dt.class = true
    default:
        return fmt.Errorf("Wrong bin value for class, received: 0x%x", data)
    }
    return nil
}

func (dt *DT) SetBinMode(data byte) error {
    switch data {
    case 0:
        dt.mode = false
    case 0x7f:
        dt.mode = true
    default:
        return fmt.Errorf("Wrong bin value for mode, received: 0x%x", data)
    }
    return nil
}

func (dt *DT) SetBinTopology(data byte) error {
    if data > 3 {
        return fmt.Errorf("Wrong bin value for topology, received: 0x%x", data)
    }
    dt.topology = data
    return nil
}

func (dt *DT) SetClass(s string) error {
    switch s {
    case "A":
        dt.class = false
    case "A/B":
        dt.class = true
    default:
        return fmt.Errorf("Value must be \"A\" or \"A\\B\", received \"%s\"", s)
    }
    return nil
}

func (dt *DT) SetMode(s string) error {
    switch s {
    case "Tri":
        dt.mode = false
    case "Pent":
        dt.mode = true
    default:
        return fmt.Errorf("Value must be \"A\" or \"A\\B\", received \"%s\"", s)
    }
    return nil
}

func (dt *DT) SetTopology(s string) error {
    for i, _s := range allowedTopology {
        if s == _s {
            dt.topology = uint8(i)
            return nil
        }
    }
    return fmt.Errorf("Value incorrect, received \"%s\"", s)
}

func (dt *DT) LockData() { dt.pb.LockData() }
func (dt *DT) UnlockData() { dt.pb.UnlockData() }

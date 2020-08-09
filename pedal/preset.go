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

import "strconv"

const (
    NumberSet = 8
    PresetPerSet = 64
)

type set struct {
    id         uint8
    name       string
    presetList [PresetPerSet]*preset
}

func newSet(id uint8, name string) *set {
    s := set{id: id, name: name}
    for i := 0; i < PresetPerSet; i++ {
        s.presetList[i] = &preset{id: uint8(i), name: "New Tone"}
    }
    return &s
}

func (s *set) GetID() uint8 {
    return s.id
}

func (s *set) GetName() string {
    return s.name
}

func (s *set) SetName(name string) {
    s.name = name
}

type preset struct {
    id   uint8
    name string
}

func (p *preset) GetID() uint8 {
    return p.id
}

func (p *preset) GetID2() string {
    id := p.GetID()
    section := (id / 4) + 1
    letter := string([]byte{uint8((id % 4)  + 65)})
    return strconv.Itoa(int(section)) + letter
}

func (p *preset) GetName() string {
    return p.name
}

func (p *preset) SetName(name string) {
    p.name = name
}

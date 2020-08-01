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
    numberSet = 8
    presetPerSet = 64
)

type set struct {
    id         uint32
    name       string
    presetList [presetPerSet]*preset
}

func newSet(id uint32, name string) *set {
    s := set{id: id, name: name}
    for i := 0; i < presetPerSet; i++ {
        s.presetList[i] = &preset{id: uint32(i), name: "New Tone"}
    }
    return &s
}

func (s *set) GetID() uint32 {
    return s.id
}

func (s *set) GetName() string {
    return s.name
}

type preset struct {
    id   uint32
    name string
}

func (p *preset) GetID() uint32 {
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

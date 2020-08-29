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

type Set struct {
    id      uint8
    name    string
    pod     *Pod
    presets [PresetPerSet]*Preset
}

func newSet(id uint8, name string, p *Pod) *Set {
    s := Set{id: id, name: name, pod: p}
    for i := range s.presets {
        s.presets[i] = newPreset(uint8(i), &s)
        s.presets[i].SetName2("New Tone")
    }
    return &s
}

func (s *Set) GetID() uint8 {
    return s.id
}

func (s *Set) GetName() string {
    return s.name
}

func (s *Set) GetPod() *Pod {
    return s.pod
}

func (s *Set) GetPreset(id uint8) *Preset {
    for _, preset := range s.presets {
        if preset.GetID() == id {
            return preset
        }
    }
    return nil
}

func (s *Set) GetPresetList() [][]string {
    ret := make([][]string, PresetPerSet)
    for i, preset := range s.presets {
        ret[i] = preset.GetName3()
    }
    return ret
}

func (s *Set) LockData() {
    s.pod.LockData()
}

func (s *Set) SetName(name string) {
    s.name = name
}

func (s *Set) UnlockData() {
    s.pod.UnlockData()
}

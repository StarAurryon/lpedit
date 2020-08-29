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
import "sync"

const (
    NumberSet = 8
    PresetPerSet = 64
)

type Pod struct {
    currentPreset *Preset
    currentSet    *Set
    mux           sync.Mutex
    sets          [NumberSet]*Set
}

func NewPod() *Pod {
    p := Pod{}
    for i := range p.sets {
        p.sets[i] = newSet(uint8(i), fmt.Sprintf("Set %d", i), &p)
    }
    p.currentSet = p.sets[0]
    p.currentPreset = p.sets[0].presets[0]

    return &p
}

func (p *Pod) GetCurrentPreset() *Preset {
    return p.currentPreset
}

func (p *Pod) GetCurrentSet() *Set {
    return p.currentSet
}

func (p *Pod) GetSet(id uint8) *Set{
    for _, set := range p.sets {
        if set.GetID() == id {
            return set
        }
    }
    return nil
}

func (p *Pod) GetSetList() []string {
    ret := make([]string, NumberSet)
    for i, set := range p.sets {
        ret[i] = set.GetName()
    }
    return ret
}

func (p *Pod) LockData() {
    p.mux.Lock()
}

func (p *Pod) SetCurrentPreset(id uint8) {
    if p.currentSet != nil {
        p.currentPreset = p.currentSet.GetPreset(id)
    } else {
        p.currentPreset = nil
    }
}

func (p *Pod) SetCurrentSet(id uint8) {
    p.currentSet = p.GetSet(id)
    p.currentPreset = nil
}

func (p *Pod) UnlockData() {
    p.mux.Unlock()
}

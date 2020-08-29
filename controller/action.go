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

package controller

import "bytes"
import "encoding/binary"
import "fmt"

import "github.com/StarAurryon/lpedit/model/pod"
import "github.com/StarAurryon/lpedit/model/pod/message"

func (c *Controller) InitPOD() {
    f := func() {
        c.QueryAllSets(false)
        c.QueryAllPresets(false)
        c.QueryCurrentSetID(false)
        c.QueryCurrentPresetID(false)
        c.QueryCurrentPreset(false)
        c.notify(nil, sg.StatusInitDone(), nil)
    }
    go f()
}

func (c *Controller) QueryAllPresets(async bool) {
    if !c.started { return }
    f := func() {
        c.syncMode = true
        max := pod.NumberSet * pod.PresetPerSet
        p := c.GetPod()
        for i := 0; i < pod.NumberSet; i++ {
            p.LockData()
            p.SetCurrentSet(uint8(i))
            p.UnlockData()
            for j := 0; j < pod.PresetPerSet; j++ {
                p.LockData()
                p.SetCurrentPreset(uint8(j))
                p.UnlockData()
                c.QueryPreset(false, uint16(j), uint16(i))
                <- c.syncModeChan
                progress := (((i * pod.PresetPerSet) + (j + 1)) * 100) / max
                c.notify(nil, sg.StatusProgress(), progress)
            }
        }
        c.syncMode = false
    }
    if async {
        go f()
    } else {
        f()
    }
}

func (c *Controller) QueryAllSets(async bool) {
    if !c.started { return }
    f := func() {
        p := c.GetPod()
        c.syncMode = true
        for i := 0; i < pod.NumberSet; i++ {
            p.LockData()
            p.SetCurrentSet(uint8(i))
            p.UnlockData()
            m := message.GenSetQuery(uint32(i))
            c.writeMessage(m, 0, 0)
            <- c.syncModeChan
            progress := ((i + 1) * 100) / pod.NumberSet
            c.notify(nil, sg.StatusProgress(), progress)
        }
        c.syncMode = false
    }
    if async {
        go f()
    } else {
        f()
    }
}

func (c *Controller) QueryCurrentPreset(async bool) {
    if !c.started { return }
    c.QueryPreset(async, message.CurrentPreset, message.CurrentSet)
}

func (c *Controller) QueryCurrentPresetID(async bool) {
    f := func() {
        if !c.started { return }
        m := message.GenStatusQueryPresetID()
        c.syncMode = true
        c.writeMessage(m, 0, 0)
        <- c.syncModeChan
        c.syncMode = false
    }
    if async {
        go f()
    } else {
        f()
    }
}

func (c *Controller) QueryCurrentSetID(async bool) {
    f := func() {
        if !c.started { return }
        m := message.GenStatusQuerySetID()
        c.syncMode = true
        c.writeMessage(m, 0, 0)
        <- c.syncModeChan
        c.syncMode = false
    }
    if async {
        go f()
    } else {
        f()
    }
}

func (c *Controller) QueryPreset(async bool, presetID uint16, setID uint16) {
    if !c.started { return }
    f := func() {
        m := message.GenPresetQuery(presetID, setID)
        c.writeMessage(m, 0, 0)
    }
    if async {
        go f()
    } else {
        f()
    }
}

func (c *Controller) ReloadPreset() {
    f := func() {
        if !c.started { return }
        c.pod.LockData()
        defer c.pod.UnlockData()
        set := c.pod.GetCurrentSet()
        if set == nil { return }
        preset := c.pod.GetCurrentPreset()
        if preset == nil { return }
        c.QueryPreset(false, uint16(preset.GetID()), uint16(set.GetID()))
    }
    go f()
}

func (c *Controller) SavePreset() {
    f := func() {
        if !c.started { return }

        c.syncMode = true
        c.QueryCurrentPreset(false)
        <- c.syncModeChan
        c.syncMode = false

        c.pod.LockData()
        defer c.pod.UnlockData()
        set := c.pod.GetCurrentSet()
        if set == nil { return }
        preset := c.pod.GetCurrentPreset()
        if preset == nil { return }

        m := message.GenStatusQuerySave()
        m2 := message.GenPresetSet(preset, c.lastLoadPreset, uint16(preset.GetID()), uint16(preset.GetSet().GetID()))
        c.writeMessage(m, 0, 0)
        c.writeMessage(m2, 0, 0)
    }
    go f()
}

func (c *Controller) SetDTClass(dtID int, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    dt := c.pod.GetCurrentPreset().GetDT(dtID)
    if dt == nil {
        return fmt.Errorf("DT not found ID:%d", dtID)
    }
    return c.setDTClass(dt, value)
}

func (c *Controller) SetDTClass2(ampID uint32, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    dt := c.pod.GetCurrentPreset().GetDT2(ampID)
    if dt == nil {
        return fmt.Errorf("DT not found AmpID:%d", ampID)
    }
    return c.setDTClass(dt, value)
}

func (c *Controller) setDTClass(dt *pod.DT, value string) error {
    err := dt.SetClass(value)
    if err != nil {
        return err
    }
    m := message.GenDTClassChange(dt)
    go c.writeMessage(m, 0, 0)
    return nil
}

func (c *Controller) SetDTMode(dtID int, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    dt := c.pod.GetCurrentPreset().GetDT(dtID)
    if dt == nil {
        return fmt.Errorf("DT not found ID:%d", dtID)
    }
    return c.setDTMode(dt, value)
}

func (c *Controller) SetDTMode2(ampID uint32, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    dt := c.pod.GetCurrentPreset().GetDT2(ampID)
    if dt == nil {
        return fmt.Errorf("DT not found AmpID:%d", ampID)
    }
    return c.setDTMode(dt, value)
}

func (c *Controller) setDTMode(dt *pod.DT, value string) error {
    err := dt.SetMode(value)
    if err != nil {
        return err
    }
    m := message.GenDTModeChange(dt)
    go c.writeMessage(m, 0, 0)
    return nil
}

func (c *Controller) SetDTTopology(dtID int, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    dt := c.pod.GetCurrentPreset().GetDT(dtID)
    if dt == nil {
        return fmt.Errorf("DT not found ID:%d", dtID)
    }
    return c.setDTTopology(dt, value)
}

func (c *Controller) SetDTTopology2(ampID uint32, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    dt := c.pod.GetCurrentPreset().GetDT2(ampID)
    if dt == nil {
        return fmt.Errorf("DT not found AmpID:%d", ampID)
    }
    return c.setDTTopology(dt, value)
}

func (c *Controller) setDTTopology(dt *pod.DT, value string) error {
    err := dt.SetTopology(value)
    if err != nil {
        return err
    }
    m := message.GenDTTopologyChange(dt)
    go c.writeMessage(m, 0, 0)
    return nil
}

func (c *Controller) SetAmpParameterValue(id uint32, pid uint32, value string) error {
    return c.SetPedalBoardItemParameterValue(id*2, pid, value)
}

func (c *Controller) SetCabParameterValue(id uint32, pid uint32, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    pbi := c.pod.GetCurrentPreset().GetItem((id*2) + 1);
    if pbi == nil {
        return nil
    }
    p := pbi.GetParam(pid)
    if p == nil {
        return nil
    }
    err := p.SetValueCurrent(value)
    if err != nil {
        return err
    }
    m := message.GenParameterCabChange(p)
    go c.writeMessage(m, 0, 0)
    c.notify(nil, sg.StatusParameterChange(), p)
    return nil
}

func (c *Controller) SetPedalParameterValue(id uint32, pid uint32, value string) error {
    return c.SetPedalBoardItemParameterValue(id+4, pid, value)
}

func (c *Controller) SetPedalBoardParameterValue(pid uint32, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    p := c.pod.GetCurrentPreset().GetParam(pid)
    if p == nil {
        return nil
    }
    err := p.SetValueCurrent(value)
    if err != nil {
        return err
    }
    m := message.GenParameterPedalBoardChange(p)
    go c.writeMessage(m, 0, 0)
    c.notify(nil, sg.StatusParameterChange(), p)
    return nil
}

func (c *Controller) SetPedalBoardItemParameterValue(id uint32, pid uint32, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    pbi := c.pod.GetCurrentPreset().GetItem(id);
    if pbi == nil {
        return nil
    }
    p := pbi.GetParam(pid)
    if p == nil {
        return nil
    }
    err := p.SetValueCurrent(value)
    if err != nil {
        return err
    }
    switch p2 := p.(type) {
    case *pod.TempoParam:
        switch p2.GetID() {
        case 0x3F100000:
            m := message.GenParameterTempoChange(p2)
            c.writeMessage(m, 0, 0)
        case 0x3F100002:
            m := message.GenParameterTempoChange2(p2)
            c.writeMessage(m, 0, 0)
        }
        binValue := p2.GetBinValueCurrent()
        var value float32
        binary.Read(bytes.NewReader(binValue[:]), binary.LittleEndian, &value)
        if value <= 1 {
            m := message.GenParameterChange(p)
            c.writeMessage(m, 0, 0)
        }
    default:
        m := message.GenParameterChange(p)
        go c.writeMessage(m, 0, 0)
    }
    c.notify(nil, sg.StatusParameterChange(), p)
    return nil
}

func (c *Controller) SetPedalBoardItemParameterValueMin(id uint32, pid uint32, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    pbi := c.pod.GetCurrentPreset().GetItem(id);
    if pbi == nil {
        return nil
    }
    p := pbi.GetParam(pid)
    if p == nil {
        return nil
    }
    err := p.SetValueMin(value)
    if err != nil {
        return err
    }
    m := message.GenParameterChangeMin(p)
    go c.writeMessage(m, 0, 0)
    return nil
}

func (c *Controller) SetPedalBoardItemParameterValueMax(id uint32, pid uint32, value string) error {
    if !c.started { return nil }
    c.pod.LockData()
    defer c.pod.UnlockData()
    pbi := c.pod.GetCurrentPreset().GetItem(id);
    if pbi == nil {
        return nil
    }
    p := pbi.GetParam(pid)
    if p == nil {
        return nil
    }
    err := p.SetValueMax(value)
    if err != nil {
        return err
    }
    m := message.GenParameterChangeMax(p)
    go c.writeMessage(m, 0, 0)
    return nil
}

func (c *Controller) SetPedalBoardItemPosition(id uint32, pos uint16, posType uint8) {
    f := func() {
        if !c.started { return }

        c.syncMode = true
        c.QueryCurrentPreset(false)
        <- c.syncModeChan
        c.syncMode = false

        c.pod.LockData()
        defer c.pod.UnlockData()
        preset := c.pod.GetCurrentPreset()
        pbi := preset.GetItem(id);
        if pbi == nil {
            return
        }
        fmt.Println("Setting pos")
        pbi.SetPos(pos, posType)
        c.setCurrentPreset(preset)
        fmt.Println("Setting pos done")
    }
    go f()
}

func (c *Controller) SetAmpActive(id uint32, active bool) {
    c.SetPedalBoardItemActive(id*2, active)
}

func (c *Controller) SetPedalActive(id uint32, active bool) {
    c.SetPedalBoardItemActive(id+4, active)
}

func (c *Controller) SetPedalBoardItemActive(id uint32, active bool) {
    f := func() {
        if !c.started { return }
        c.pod.LockData()
        pbi := c.pod.GetCurrentPreset().GetItem(id)
        if pbi == nil {
            c.pod.UnlockData()
            return
        }
        pbi.SetActive(active)
        m := message.GenActiveChange(pbi)
        c.pod.UnlockData()
        c.writeMessage(m, 0, 0)
    }
    go f()
}

func (c *Controller) SetAmpType(id uint32, name string) {
    c.SetPedalBoardItemType(id*2, name, "")
}

func (c *Controller) SetCabType(id uint32, name string) {
    c.SetPedalBoardItemType(id*2+1, name, "")
}

func (c *Controller) SetPedalType(id uint32, fxType string, fxModel string) {
    c.SetPedalBoardItemType(id+4, fxType, fxModel)
}

func (c *Controller) SetPedalBoardItemType(id uint32, fxType string, fxModel string) {
    if !c.started { return }
    c.pod.LockData()
    pbi := c.pod.GetCurrentPreset().GetItem(id)
    if pbi == nil {
        c.pod.UnlockData()
        return
    }
    pbi.SetType2(fxType, fxModel)
    m := message.GenTypeChange(pbi)
    m2 := message.GenPresetQuery(message.CurrentPreset, message.CurrentSet)
    c.pod.UnlockData()
    f := func() {
        c.writeMessage(m, 0, 0)
        c.writeMessage(m2, 0, 0)
    }
    go f()
}

func (c *Controller) SetPreset(presetID uint8, setID uint8) {
    f := func() {
        if !c.started { return }

        c.pod.LockData()
        c.pod.SetCurrentSet(setID)
        c.pod.SetCurrentPreset(presetID)
        c.pod.UnlockData()

        c.syncMode = true
        c.QueryPreset(false, uint16(presetID), uint16(setID))
        <- c.syncModeChan
        c.syncMode = false

        m := message.GenSetChange(setID)
        m2 := message.GenPresetChange(presetID)
        c.writeMessage(m, 0, 0)
        c.writeMessage(m2, 0, 0)

        c.pod.LockData()
        c.setCurrentPreset(c.pod.GetCurrentPreset())
        c.pod.UnlockData()
    }
    go f()
}

func (c *Controller) setCurrentPreset(p *pod.Preset) {
    m := message.GenPresetSet(p, c.lastLoadPreset, message.CurrentPreset, message.CurrentSet)
    c.writeMessage(m, 0, 0)
}

func (c *Controller) SetCurrentPresetName(name string) {
    f := func() {
        if !c.started { return }

        c.syncMode = true
        c.QueryCurrentPreset(false)
        <- c.syncModeChan
        c.syncMode = false

        c.pod.LockData()
        defer c.pod.UnlockData()

        preset := c.pod.GetCurrentPreset()
        preset.SetName2(name)
        c.setCurrentPreset(preset)
        c.QueryCurrentPreset(false)
    }
    go f()
}

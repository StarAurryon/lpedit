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
    nonePedal uint32 = 34471935
)

type Pedal struct {
    id      uint32
    active  bool
    name    string
    ptype   uint32
    stype   string
    pos     uint16
    posType uint8
    params  []Parameter
    pb      *PedalBoard
}

var pedals = []Pedal {
    Pedal{ptype: nonePedal, active: true, stype: "None", name: "None"},
    /*
     * Dynamics section
     */
    Pedal{ptype: 33554443, active: true, stype: "Dynamics", name: "Red Comp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Sustain"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Level"}},
            }},
    Pedal{ptype: 33554444, active: true, stype: "Dynamics", name: "Blue Comp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Sustain"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Level"}},
            }},
    Pedal{ptype: 33554445, active: true, stype: "Dynamics", name: "Blue Comp Treble",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Level"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Sustain"}},
            }},
    Pedal{ptype: 33554446, active: true, stype: "Dynamics", name: "Tube Comp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Threshold"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Level"}},
            }},
    Pedal{ptype: 33554447, active: true, stype: "Dynamics", name: "Vetta Comp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Sensitivity"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Level"}},
            }},
    Pedal{ptype: 33554448, active: true, stype: "Dynamics", name: "Vetta Juice",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Amount"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Level"}},
            }},
    Pedal{ptype: 33554449, active: true, stype: "Dynamics", name: "Noise Gate",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Threshold"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Decay"}},
            }},
    Pedal{ptype: 33554450, active: true, stype: "Dynamics", name: "Boost Comp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Comp"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33554451, active: true, stype: "Dynamics", name: "Hard Gate",
        params: []Parameter{
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Open Threshold"}, max: 0, min: -96},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Close Threshold"}, max: 0, min: -96},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Hold"}, maxMs: 800},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Decay"}, maxMs: 4000},
            }},
    /*
     * Delay section
     */
    Pedal{ptype: 33685521, active: true, stype: "Delay", name: "Digital Delay",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685522, active: true, stype: "Delay", name: "Digital Delay W/Mod",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "ModSpd"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685523, active: true, stype: "Delay", name: "Stereo Delay",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "L Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "L-FDBK"}},
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "R Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "R-FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685524, active: true, stype: "Delay", name: "Analog Echo",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685525, active: true, stype: "Delay", name: "Analog W/Mod",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "ModSpd"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685530, active: true, stype: "Delay", name: "Multi-Head",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Heads 1-2"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Heads 3-4"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685533, active: true, stype: "Delay", name: "Low Res Delay",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Res"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685536, active: true, stype: "Delay", name: "Ping Pong",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Offset"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Spread"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685537, active: true, stype: "Delay", name: "Reverse",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "ModSpd"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685538, active: true, stype: "Delay", name: "Dynamic Delay",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Thresh"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Ducking"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685539, active: true, stype: "Delay", name: "Auto-Volume Echo",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685546, active: true, stype: "Delay", name: "Tube Echo",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Wow/Flt"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685547, active: true, stype: "Delay", name: "Tube Echo Dry",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Wow/Flt"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685548, active: true, stype: "Delay", name: "Tape Echo",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685549, active: true, stype: "Delay", name: "Tape Echo Dry",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685550, active: true, stype: "Delay", name: "Sweep Echo",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Swp Spd"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Swp Dep"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685551, active: true, stype: "Delay", name: "Sweep Echo Dry",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Swp Spd"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Swp Dep"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685552, active: true, stype: "Delay", name: "Echo Platter",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Wow/Flt"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    Pedal{ptype: 33685553, active: true, stype: "Delay", name: "Echo Platter Dry",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Time"}, min: 20, max: 2000},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "FDBK"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Wow/Flt"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    /*
     * Modulation section
     */
    Pedal{ptype: 33751072, active: true, stype: "Modulation", name: "Opto Tremolo",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Shape"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "VolSens"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751073, active: true, stype: "Modulation", name: "Bias Tremolo",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Shape"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "VolSens"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751074, active: true, stype: "Modulation", name: "Phaser",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Fdbk"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Stages"}, binValueType: float32Type,
                list: []string{"STG 4", "STG 8", "STG 12", "STG 16"},
            },
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751075, active: true, stype: "Modulation", name: "Dual Phaser",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Fdbk"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "LfoShp"}, binValueType: float32Type,
                list: []string{"Sine", "Square"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751076, active: true, stype: "Modulation", name: "Panned Phaser",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Pan"}, binValueType: float32Type,
                list: []string{"Left", "Center", "Right"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "PanSpd"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751077, active: true, stype: "Modulation", name: "U-Vibe",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Fdbk"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "VolSens"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751078, active: true, stype: "Modulation", name: "Rotary Drum",
        params: []Parameter{
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, binValueType: float32Type,
                list: []string{"Slow", "Fast"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751079, active: true, stype: "Modulation", name: "Rotary Drum/Hrn",
        params: []Parameter{
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, binValueType: float32Type,
                list: []string{"Slow", "Fast"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Horn Dep"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751080, active: true, stype: "Modulation", name: "Analog Flanger",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Fdbk"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Manual"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751081, active: true, stype: "Modulation", name: "Jet Flanger",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Fdbk"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Manual"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751082, active: true, stype: "Modulation", name: "Analog Chorus",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Ch Vib"}, binValueType: float32Type,
                list: []string{"Chorus", "Vibrato"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751083, active: true, stype: "Modulation", name: "Dimension",
        params: []Parameter{
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Sw1"}, binValueType: float32Type,
                list: []string{"SW1 OFF", "SW1 ON"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Sw2"}, binValueType: float32Type,
                list: []string{"SW2 OFF", "SW2 ON"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Sw3"}, binValueType: float32Type,
                list: []string{"SW3 OFF", "SW3 ON"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Sw4"}, binValueType: float32Type,
                list: []string{"SW4 OFF", "SW4 ON"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751084, active: true, stype: "Modulation", name: "Tri Chorus",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Depth2"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Depth3"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751085, active: true, stype: "Modulation", name: "Pitch Vibrato",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Rise"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "VolSens"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751086, active: true, stype: "Modulation", name: "Ring Modulator",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Shape"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "AM/FM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Speed"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751087, active: true, stype: "Modulation", name: "Panner",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Depth"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Shape"}, binValueType: float32Type,
                list: []string{"Triangle", "Sine", "Square"},
            },
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "VolSens"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751104, active: true, stype: "Modulation", name: "Barberpole Phaser",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Fdbk"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Mode"}, binValueType: float32Type,
                list: []string{"Up", "Down", "Stereo"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751106, active: true, stype: "Modulation", name: "Frequency Shifter",
        params: []Parameter{
            &FreqParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}, min: 0, max: 3520},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Mode"}, binValueType: float32Type,
                list: []string{"Up", "Down", "Stereo"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33751107, active: true, stype: "Modulation", name: "Pattern Tremolo",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Step1"}, binValueType: float32Type,
                list: []string{"Mute", "1", "2", "3", "4", "5", "6", "7",
                     "8", "9", "10", "11", "12", "13", "14", "15", "16",
                      "Full", "Skip"},
                  },
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Step2"}, binValueType: float32Type,
                list: []string{"Mute", "1", "2", "3", "4", "5", "6", "7",
                     "8", "9", "10", "11", "12", "13", "14", "15", "16",
                      "Full", "Skip"},
                  },
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Step3"}, binValueType: float32Type,
                list: []string{"Mute", "1", "2", "3", "4", "5", "6", "7",
                     "8", "9", "10", "11", "12", "13", "14", "15", "16",
                      "Full", "Skip"},
                  },
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Step4"}, binValueType: float32Type,
                list: []string{"Mute", "1", "2", "3", "4", "5", "6", "7",
                     "8", "9", "10", "11", "12", "13", "14", "15", "16",
                      "Full", "Skip"},
                  },
            }},
    Pedal{ptype: 33751109, active: true, stype: "Modulation", name: "Script Phase",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            }},
    Pedal{ptype: 33751111, active: true, stype: "Modulation", name: "AC Flanger",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Width"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Regen"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Manual"}},
            }},
    Pedal{ptype: 33751113, active: true, stype: "Modulation", name: "80A Flanger",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Range"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Enhance"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Manual"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Even Odd"}, binValueType: float32Type,
                list: []string{"Even", "Odd"}},
            }},
    /*
     * Reverb section
     */
    Pedal{ptype: 33816604, active: true, stype: "Reverb", name: "'63 Spring",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816605, active: true, stype: "Reverb", name: "Spring",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816606, active: true, stype: "Reverb", name: "Plate",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816607, active: true, stype: "Reverb", name: "Room",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816608, active: true, stype: "Reverb", name: "Chamber",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816609, active: true, stype: "Reverb", name: "Hall",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816610, active: true, stype: "Reverb", name: "Ducking",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816611, active: true, stype: "Reverb", name: "Octo",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816612, active: true, stype: "Reverb", name: "Cave",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816613, active: true, stype: "Reverb", name: "Tile",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816614, active: true, stype: "Reverb", name: "Echo",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Decay"}},
            &TimeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Predelay"}, maxMs: 200},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33816615, active: true, stype: "Reverb", name: "Particle Verb",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Dwell"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Condition"}, binValueType: float32Type,
                list: []string{"Stable", "Critical", "Hazard"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Gain"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    /*
     * Distortion section
     */
    Pedal{ptype: 33882122, active: true, stype: "Distortion", name: "Jet Fuzz",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Fdbk"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Speed"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882123, active: true, stype: "Distortion", name: "Classic Dist",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Filter"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882124, active: true, stype: "Distortion", name: "Octave Fuzz",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882125, active: true, stype: "Distortion", name: "Tube Drive",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882126, active: true, stype: "Distortion", name: "Screamer",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882127, active: true, stype: "Distortion", name: "Fuzz Pi",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882128, active: true, stype: "Distortion", name: "Overdrive",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882129, active: true, stype: "Distortion", name: "Facial Fuzz",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882130, active: true, stype: "Distortion", name: "Line 6 Distortion",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882131, active: true, stype: "Distortion", name: "Line 6 Drive",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882132, active: true, stype: "Distortion", name: "Sub Octave Fuzz",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Sub"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882133, active: true, stype: "Distortion", name: "Buzz Saw",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882134, active: true, stype: "Distortion", name: "Heavy Dist",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882135, active: true, stype: "Distortion", name: "Jumbo Fuzz",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    Pedal{ptype: 33882136, active: true, stype: "Distortion", name: "Color Drive",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Output"}},
            }},
    /*
     * Wah section
     */
    Pedal{ptype: 33947659, active: true, stype: "Wah", name: "Vetta Wah",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Position"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33947660, active: true, stype: "Wah", name: "Fassel",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Position"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33947661, active: true, stype: "Wah", name: "Weeper",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Position"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33947662, active: true, stype: "Wah", name: "Chrome",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Position"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33947663, active: true, stype: "Wah", name: "Chrome Custom",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Position"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33947664, active: true, stype: "Wah", name: "Throaty",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Position"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33947665, active: true, stype: "Wah", name: "Conductor",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Position"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 33947666, active: true, stype: "Wah", name: "Colorful",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Position"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    /*
     * Volume/Pan section
     */
    Pedal{ptype: 34013188, active: true, stype: "Volume/Pan", name: "Volume",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Volume Level"}},
            }},
    Pedal{ptype: 34013189, active: true, stype: "Volume/Pan", name: "Pan",
        params: []Parameter{
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Pan L-R Balance"}, min: -100, max: 100},
            }},
    /*
     * FX Loop
     */
    Pedal{ptype: 34078720, active: true, stype: "FX Loop", name: "FX Loop",
        params: []Parameter{
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Send"}, min: -80, max: 0},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Return"}, min: 0, max: 24},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mix"}},
            }},
    /*
     * Pitch section
     */
    Pedal{ptype: 34144258, active: true, stype: "Pitch", name: "Pitch Glide",
        params: []Parameter{
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Pitch"}, min: -24, max: 24},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F010001, name: "Mix"}},
            }},
    Pedal{ptype: 34144259, active: true, stype: "Pitch", name: "Smart Harmony",
        params: []Parameter{
            &ListParam{GenericParameter: GenericParameter{id: 0x3F001801, name: "Key"}, binValueType: Int32Type,
                list: []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F",
                    "F#", "G", "G#"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F001802, name: "Shift"}, binValueType: Int32Type,
                maxIDShift: -8,
                list: []string{"-9th", "-8th", "-7th", "-6th", "-5th", "-4th",
                    "-3rd", "-2nd", "None", "2nd", "3rd", "4th", "5th", "6th",
                    "7th", "8th", "9th"},
            },
            &ListParam{GenericParameter: GenericParameter{id: 0x3F001803, name: "Scale"}, binValueType: Int32Type,
                list: []string{"Major", "Minor", "Maj. Pent.", "Min. Pent",
                    "Harm. Min.", "Mel. Min.", "Whole", "Whole D"},
            },
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F010001, name: "Mix"}},
            }},
    Pedal{ptype: 34144260, active: true, stype: "Pitch", name: "Bass Octaver",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Tone"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Normal"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Octave"}},
            }},
    /*
     * Preamp+EQ section
     */
    Pedal{ptype: 34209807, active: true, stype: "Filter", name: "Tron Up",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Range"}, binValueType: float32Type,
                list: []string{"Low", "High"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Type"}, binValueType: float32Type,
                list: []string{"LP", "BP", "HP"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209808, active: true, stype: "Filter", name: "Tron Down",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Range"}, binValueType: float32Type,
                list: []string{"Low", "High"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Type"}, binValueType: float32Type,
                list: []string{"LP", "BP", "HP"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209809, active: true, stype: "Filter", name: "Seeker",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Steps"}, binValueType: float32Type,
                maxIDShift: 1, list: []string{"2 Steps", "3 Steps",
                "4 Steps", "5 Steps", "6 Steps", "7 Steps", "8 Steps", "9 Steps"},
            },
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209810, active: true, stype: "Filter", name: "Obi Wah",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Type"}, binValueType: float32Type,
                list: []string{"LP", "BP", "HP"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209811, active: true, stype: "Filter", name: "Slow Filter",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Speed"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mode"}, binValueType: float32Type,
                list: []string{"Up", "Down"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209812, active: true, stype: "Filter", name: "Q Filter",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Gain"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Type"}, binValueType: float32Type,
                list: []string{"LP", "BP", "HP"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209813, active: true, stype: "Filter", name: "Throbber",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Wave"}, binValueType: float32Type,
                list: []string{"Ramp Up", "Ramp Down", "Triangle", "Square",},
            },
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209814, active: true, stype: "Filter", name: "Spin Cycle",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "VolSens"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209815, active: true, stype: "Filter", name: "Comet Trails",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Gain"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209816, active: true, stype: "Filter", name: "Octisynth",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Depth"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209817, active: true, stype: "Filter", name: "Growler",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Pitch"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209818, active: true, stype: "Filter", name: "Synth O Matic",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Q"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Wave"}, binValueType: float32Type,
                list: []string{"Wave 1", "Wave 2", "Wave 3", "Wave 4", "Wave 5",
                    "Wave 6", "Wave 7", "Wave 8"},
            },
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Pitch"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209819, active: true, stype: "Filter", name: "Attack Synth",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Wave"}, binValueType: float32Type,
                list: []string{"Square", "PWM", "Ramp"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Speed"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Pitch"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209820, active: true, stype: "Filter", name: "Synth String",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Attack"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Pitch"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209821, active: true, stype: "Filter", name: "Voice Box",
        params: []Parameter{
            &TempoParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Speed"}, min: 0.10, max: 15},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Start"}, binValueType: float32Type,
                list: []string{"A", "E", "I", "O", "U"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "End"}, binValueType: float32Type,
                list: []string{"A", "E", "I", "O", "U"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Auto"}, binValueType: float32Type,
                list: []string{"Auto 1", "Auto 2", "Auto 3", "Auto 4"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209822, active: true, stype: "Filter", name: "V-Tron",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Speed"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Start"}, binValueType: float32Type,
                list: []string{"A", "E", "I", "O", "U"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "End"}, binValueType: float32Type,
                list: []string{"A", "E", "I", "O", "U"}},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Mode"}, binValueType: float32Type,
                list: []string{"Up", "Up/Down"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    Pedal{ptype: 34209830, active: true, stype: "Filter", name: "Vocoder",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Input"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Mic"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Decay"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Mix"}},
            }},
    /*
     * Preamp+EQ section
     */
    Pedal{ptype: 34340873, active: true, stype: "Preamp+EQ", name: "Graphic EQ",
        params: []Parameter{
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "80Hz"}, min: -12, max: 12},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "220Hz"}, min: -12, max: 12},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "480Hz"}, min: -12, max: 12},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "1.1kHz"}, min: -12, max: 12},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "2.2kHz"}, min: -12, max: 12},
            }},
    Pedal{ptype: 34340874, active: true, stype: "Preamp+EQ", name: "Studio EQ",
        params: []Parameter{
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Low Freq"}, binValueType: float32Type,
                maxIDShift: 1, list: []string{"75Hz", "150Hz",
                "180Hz", "240Hz", "500Hz", "700Hz", "1000Hz", "1400Hz"}},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Low Gain"}, min: -11, max: 11},
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Hi Freq"}, binValueType: float32Type,
                maxIDShift: 1, list: []string{"200Hz", "300Hz", "400Hz", "800Hz",
                    "1500Hz", "3000Hz", "5000Hz", "8000Hz"}},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Hi Gain"}, min: -11, max: 11},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Gain"}, min: -18, max: 18},
            }},
    Pedal{ptype: 34340875, active: true, stype: "Preamp+EQ", name: "Parametric EQ",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Lows"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Highs"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Q"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Gain"}},
            }},
    Pedal{ptype: 34340876, active: true, stype: "Preamp+EQ", name: "4 Band Shift EQ",
        params: []Parameter{
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Low"}, min: -12, max: 12},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Low Mid"}, min: -12, max: 12},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Hi Mid"}, min: -12, max: 12},
            &RangeParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Hi"}, min: -12, max: 12},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Shift"}},
            }},
    Pedal{ptype: 34340877, active: true, stype: "Preamp+EQ", name: "Mid Focus EQ",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Hi Pass Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Hi Pass Q"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Low Pass Freq"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Low Pass Q"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Gain"}},
            }},
    Pedal{ptype: 34340878, active: true, stype: "Preamp+EQ", name: "Vintage Pre",
        params: []Parameter{
            &ListParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Phase"}, binValueType: float32Type,
                list: []string{"0", "180"},
            },
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Gain"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Output"}},
            &FreqParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Hi Pass Filter"}, min: 20, max: 500},
            &FreqKParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Low pass Filter"}, min:5, max: 20},
            }},
}

func newNonePedal(id uint32, pos uint16, posType uint8, pb *PedalBoard) *Pedal{
    p := newPedal(id, pos, posType, pb, nonePedal)
    return p
}

func newPedal(id uint32, pos uint16, posType uint8, pb *PedalBoard, ptype uint32 ) *Pedal {
    for _, newPedal := range pedals {
        if newPedal.ptype == ptype {
            newPedal.id = id
            newPedal.pos = pos
            newPedal.posType = posType
            newPedal.pb = pb
            for i := range newPedal.params {
                newPedal.params[i] = newPedal.params[i].Copy()
                newPedal.params[i].SetParent(&newPedal)
            }
            return &newPedal
        }
    }
    return nil
}

func GetPedalType() map[string][]string {
    m := map[string][]string{}
    for _, p := range pedals {
        m[p.stype] = append(m[p.stype], p.GetName())
    }
    return m
}

func (p *Pedal) GetActive() bool {
    return p.active
}

func (p *Pedal) GetActive2() uint32 {
    if p.active {
        return 1
    }
    return 0
}

func (p *Pedal) GetID() uint32 {
    return p.id
}

func (p *Pedal) GetName() string {
    return p.name
}

func (p *Pedal) GetParam(id uint32) Parameter {
    for _, param := range p.params {
        if param.GetID() == id {
            return param
        }
    }
    return nil
}

func (p *Pedal) GetParams() []Parameter {
    return p.params
}

func (p *Pedal) GetParamLen() uint16 {
    return uint16(len(p.params)) //parameter start at 1
}

func (p *Pedal) GetPos() (uint16, uint8) {
    return p.pos, p.posType
}

func (p *Pedal) GetType() uint32 {
    return p.ptype
}

func (p *Pedal) GetSType() string {
    return p.stype
}

func (p *Pedal) LockData() { p.pb.LockData() }

func (p *Pedal) SetPos(pos uint16, posType uint8) {
    if posType == AmpAPos || posType == AmpBPos || pos >= 8 { return }
    var incr int
    if int(pos) - int(p.pos) > 0 {
        incr = 1
    } else {
        incr = -1
    }
    for i := int(p.pos) + incr; i != int(pos) + incr ; i += incr {
        p2 := p.pb.GetPedal(uint16(i))
        p2Pos, p2PosType := p2.GetPos()
        p2.SetPosWithoutCheck(uint16(int(p2Pos)-incr), p2PosType)
        p2Pos, _ = p2.GetPos()
    }
    p.pos = pos
    p.posType = posType
}

func (p *Pedal) SetPosWithoutCheck(pos uint16, posType uint8) {
    if posType == AmpAPos || posType == AmpBPos { return }
    p.pos = pos
    p.posType = posType
}

func (p *Pedal) SetActive(active bool){
    p.active = active
}

func (p *Pedal) SetType(ptype uint32) error {
    _p := newPedal(p.id, p.pos, p.posType, p.pb, ptype)
    if _p == nil {
        return fmt.Errorf("Pedal type not found, code: %d", ptype)
    }
    *p = *_p
    for i := range p.params {
        p.params[i].SetParent(p)
    }
    return nil
}

func (p *Pedal) SetType2(stype string, name string) {
    for _, _p := range pedals {
        if stype == _p.stype && name == _p.name {
            p.SetType(_p.ptype)
            break
        }
    }
}

func (p *Pedal) UnlockData() { p.pb.UnlockData() }

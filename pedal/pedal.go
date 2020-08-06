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
            &NullParam{},
            &PerCentParam{name: "Sustain", value:0},
            &PerCentParam{name: "Level", value:0},
            }},
    Pedal{ptype: 33554444, active: true, stype: "Dynamics", name: "Blue Comp",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Sustain", value:0},
            &PerCentParam{name: "Level", value:0},
            }},
    Pedal{ptype: 33554445, active: true, stype: "Dynamics", name: "Blue Comp Treble",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Level", value:0},
            &PerCentParam{name: "Sustain", value:0},
            }},
    Pedal{ptype: 33554446, active: true, stype: "Dynamics", name: "Tube Comp",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Threshold", value:0},
            &PerCentParam{name: "Level", value:0},
            }},
    Pedal{ptype: 33554447, active: true, stype: "Dynamics", name: "Vetta Comp",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Sensitivity", value:0},
            &PerCentParam{name: "Level", value:0},
            }},
    Pedal{ptype: 33554448, active: true, stype: "Dynamics", name: "Vetta Juice",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Amount", value:0},
            &PerCentParam{name: "Level", value:0},
            }},
    Pedal{ptype: 33554449, active: true, stype: "Dynamics", name: "Noise Gate",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Threshold", value:0},
            &PerCentParam{name: "Decay", value:0},
            }},
    Pedal{ptype: 33554450, active: true, stype: "Dynamics", name: "Boost Comp",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Comp", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33554451, active: true, stype: "Dynamics", name: "Hard Gate",
        params: []Parameter{
            &RangeParam{name: "Open Threshold", max: 0, min: -96, increment: 1, value:0},
            &RangeParam{name: "Close Threshold", max: 0, min: -96, increment: 1, value:0},
            &TimeParam{name: "Hold", value:0, maxMs: 800},
            &TimeParam{name: "Decay", value:0, maxMs: 4000},
            }},
    /*
     * Delay section
     */
    Pedal{ptype: 33685521, active: true, stype: "Delay", name: "Digital Delay",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685522, active: true, stype: "Delay", name: "Digital Delay W/Mod",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "ModSpd", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685523, active: true, stype: "Delay", name: "Stereo Delay",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "L Time", value:0},
            &PerCentParam{name: "L-FDBK", value:0},
            &PerCentParam{name: "R Time", value:0},
            &PerCentParam{name: "R-FDBK", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685524, active: true, stype: "Delay", name: "Analog Echo",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685525, active: true, stype: "Delay", name: "Analog W/Mod",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "ModSpd", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685530, active: true, stype: "Delay", name: "Multi-Head",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Heads 1-2", value:0},
            &PerCentParam{name: "Heads 3-4", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685533, active: true, stype: "Delay", name: "Low Res Delay",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Tone", value:0},
            &PerCentParam{name: "Res", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685536, active: true, stype: "Delay", name: "Ping Pong",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Offset", value:0},
            &PerCentParam{name: "Spread", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685537, active: true, stype: "Delay", name: "Reverse",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "ModSpd", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685538, active: true, stype: "Delay", name: "Dynamic Delay",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Thresh", value:0},
            &PerCentParam{name: "Ducking", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685539, active: true, stype: "Delay", name: "Auto-Volume Echo",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685546, active: true, stype: "Delay", name: "Tube Echo",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Wow/Flt", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685547, active: true, stype: "Delay", name: "Tube Echo Dry",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Wow/Flt", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685548, active: true, stype: "Delay", name: "Tape Echo",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685549, active: true, stype: "Delay", name: "Tape Echo Dry",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685550, active: true, stype: "Delay", name: "Sweep Echo",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Swp Spd", value:0},
            &PerCentParam{name: "Swp Dep", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685551, active: true, stype: "Delay", name: "Sweep Echo Dry",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Swp Spd", value:0},
            &PerCentParam{name: "Swp Dep", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685552, active: true, stype: "Delay", name: "Echo Platter",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Wow/Flt", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33685553, active: true, stype: "Delay", name: "Echo Platter",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Time", value:0},
            &PerCentParam{name: "FDBK", value:0},
            &PerCentParam{name: "Wow/Flt", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    /*
     * Modulation section
     */
    Pedal{ptype: 33751072, active: true, stype: "Modulation", name: "Opto Tremolo",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Shape", value:0},
            &PerCentParam{name: "VolSens", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751073, active: true, stype: "Modulation", name: "Bias Tremolo",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Shape", value:0},
            &PerCentParam{name: "VolSens", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751074, active: true, stype: "Modulation", name: "Phaser",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Fdbk", value:0},
            &ListParam{name: "Stages", value:0, list: []string{"STG 4", "STG 8",
                "STG 12", "STG 16"},
            },
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751075, active: true, stype: "Modulation", name: "Dual Phaser",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Fdbk", value:0},
            &ListParam{name: "LfoShp", value:0, list: []string{"Sine", "Square"}},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751076, active: true, stype: "Modulation", name: "Panned Phaser",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &ListParam{name: "Pan", value:0, list: []string{"Left", "Center", "Right"}},
            &PerCentParam{name: "PanSpd", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751077, active: true, stype: "Modulation", name: "U-Vibe",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Fdbk", value:0},
            &PerCentParam{name: "VolSens", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751078, active: true, stype: "Modulation", name: "Rotary Drum",
        params: []Parameter{
            &ListParam{name: "Speed", value:0, list: []string{"Slow", "Fast"}},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Tone", value:0},
            &PerCentParam{name: "Drive", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751079, active: true, stype: "Modulation", name: "Rotary Drum/Hrn",
        params: []Parameter{
            &ListParam{name: "Speed", value:0, list: []string{"Slow", "Fast"}},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Horn Dep", value:0},
            &PerCentParam{name: "Drive", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751080, active: true, stype: "Modulation", name: "Analog Flanger",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Fdbk", value:0},
            &PerCentParam{name: "Manual", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751081, active: true, stype: "Modulation", name: "Jet Flanger",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Fdbk", value:0},
            &PerCentParam{name: "Manual", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751082, active: true, stype: "Modulation", name: "Analog Chorus",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &ListParam{name: "Ch Vib", value:0, list: []string{"Chorus", "Vibrato"}},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751083, active: true, stype: "Modulation", name: "Dimension",
        params: []Parameter{
            &NullParam{},
            &ListParam{name: "Sw1", value:0, list: []string{"SW1 OFF", "SW1 ON"}},
            &ListParam{name: "Sw2", value:0, list: []string{"SW2 OFF", "SW2 ON"}},
            &ListParam{name: "Sw3", value:0, list: []string{"SW3 OFF", "SW3 ON"}},
            &ListParam{name: "Sw4", value:0, list: []string{"SW4 OFF", "SW4 ON"}},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751084, active: true, stype: "Modulation", name: "Tri Chorus",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Depth2", value:0},
            &PerCentParam{name: "Depth3", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751085, active: true, stype: "Modulation", name: "Pitch Vibrato",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Rise", value:0},
            &PerCentParam{name: "VolSens", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751086, active: true, stype: "Modulation", name: "Ring Modulator",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Depth", value:0},
            &PerCentParam{name: "Shape", value:0},
            &PerCentParam{name: "AM/FM", value:0},
            &PerCentParam{name: "Speed", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751087, active: true, stype: "Modulation", name: "Panner",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Depth", value:0},
            &ListParam{name: "Shape", value:0, list: []string{"Triangle",
                "Sine", "Square"},
            },
            &PerCentParam{name: "VolSens", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751104, active: true, stype: "Modulation", name: "Barberpole Phaser",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Fdbk", value:0},
            &ListParam{name: "Mode", value:0, list: []string{"Up", "Down", "Stereo"}},
            &NullParam{},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751106, active: true, stype: "Modulation", name: "Frequency Shifter",
        params: []Parameter{
            &NullParam{},
            &FreqParam{name: "Freq", value:0, min: 0, max: 3520},
            &ListParam{name: "Mode", value:0, list: []string{"Up", "Down", "Stereo"}},
            &NullParam{},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33751107, active: true, stype: "Modulation", name: "Pattern Tremolo",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &ListParam{name: "Step1", value:0,
                list: []string{"Mute", "1", "2", "3", "4", "5", "6", "7",
                     "8", "9", "10", "11", "12", "13", "14", "15", "16",
                      "Full", "Skip"},
                  },
            &ListParam{name: "Step2", value:0,
                list: []string{"Mute", "1", "2", "3", "4", "5", "6", "7",
                     "8", "9", "10", "11", "12", "13", "14", "15", "16",
                      "Full", "Skip"},
                  },
            &ListParam{name: "Step3", value:0,
                list: []string{"Mute", "1", "2", "3", "4", "5", "6", "7",
                     "8", "9", "10", "11", "12", "13", "14", "15", "16",
                      "Full", "Skip"},
                  },
            &ListParam{name: "Step4", value:0,
                list: []string{"Mute", "1", "2", "3", "4", "5", "6", "7",
                     "8", "9", "10", "11", "12", "13", "14", "15", "16",
                      "Full", "Skip"},
                  },
            }},
    Pedal{ptype: 33751109, active: true, stype: "Modulation", name: "Script Phase",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            }},
    Pedal{ptype: 33751111, active: true, stype: "Modulation", name: "AC Flanger",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Width", value:0},
            &PerCentParam{name: "Regen", value:0},
            &PerCentParam{name: "Manual", value:0},
            }},
    Pedal{ptype: 33751113, active: true, stype: "Modulation", name: "80A Flanger",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Range", value:0},
            &PerCentParam{name: "Enhance", value:0},
            &PerCentParam{name: "Manual", value:0},
            &ListParam{name: "Even Odd", value:0, list: []string{"Even", "Odd"}},
            }},
    /*
     * Reverb section
     */
    Pedal{ptype: 33816604, active: true, stype: "Reverb", name: "'63 Spring",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816605, active: true, stype: "Reverb", name: "Spring",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816606, active: true, stype: "Reverb", name: "Plate",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816607, active: true, stype: "Reverb", name: "Room",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816608, active: true, stype: "Reverb", name: "Chamber",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816609, active: true, stype: "Reverb", name: "Hall",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816610, active: true, stype: "Reverb", name: "Ducking",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816611, active: true, stype: "Reverb", name: "Octo",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816612, active: true, stype: "Reverb", name: "Cave",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816613, active: true, stype: "Reverb", name: "Tile",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816614, active: true, stype: "Reverb", name: "Echo",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Decay", value:0},
            &TimeParam{name: "Predelay", value:0, maxMs: 200},
            &PerCentParam{name: "Tone", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33816615, active: true, stype: "Reverb", name: "Particle Verb",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Dwell", value:0},
            &PerCentParam{name: "Condition", value:0},
            &PerCentParam{name: "Gain", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    /*
     * Distortion section
     */
    Pedal{ptype: 33882122, active: true, stype: "Distortion", name: "Jet Fuzz",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Fdbk", value:0},
            &PerCentParam{name: "Tone", value:0},
            &PerCentParam{name: "Speed", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882123, active: true, stype: "Distortion", name: "Classic Dist",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Filter", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882124, active: true, stype: "Distortion", name: "Octave Fuzz",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882125, active: true, stype: "Distortion", name: "Tube Drive",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882126, active: true, stype: "Distortion", name: "Screamer",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Tone", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882127, active: true, stype: "Distortion", name: "Fuzz Pi",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882128, active: true, stype: "Distortion", name: "Overdrive",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882129, active: true, stype: "Distortion", name: "Facial Fuzz",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882130, active: true, stype: "Distortion", name: "Line 6 Distortion",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882131, active: true, stype: "Distortion", name: "Line 6 Drive",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882132, active: true, stype: "Distortion", name: "Sub Octave Fuzz",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Sub", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882133, active: true, stype: "Distortion", name: "Buzz Saw",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882134, active: true, stype: "Distortion", name: "Heavy Dist",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882135, active: true, stype: "Distortion", name: "Jumbo Fuzz",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    Pedal{ptype: 33882136, active: true, stype: "Distortion", name: "Color Drive",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Output", value:0},
            }},
    /*
     * Wah section
     */
    Pedal{ptype: 33947659, active: true, stype: "Wah", name: "Vetta Wah",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Position", value:0},
            &NullParam{},
            &NullParam{},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33947660, active: true, stype: "Wah", name: "Fassel",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Position", value:0},
            &NullParam{},
            &NullParam{},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33947661, active: true, stype: "Wah", name: "Weeper",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Position", value:0},
            &NullParam{},
            &NullParam{},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33947662, active: true, stype: "Wah", name: "Chrome",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Position", value:0},
            &NullParam{},
            &NullParam{},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33947663, active: true, stype: "Wah", name: "Chrome Custom",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Position", value:0},
            &NullParam{},
            &NullParam{},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33947664, active: true, stype: "Wah", name: "Throaty",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Position", value:0},
            &NullParam{},
            &NullParam{},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33947665, active: true, stype: "Wah", name: "Conductor",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Position", value:0},
            &NullParam{},
            &NullParam{},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 33947666, active: true, stype: "Wah", name: "Colorful",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Position", value:0},
            &NullParam{},
            &NullParam{},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    /*
     * Volume/Pan section
     */
    Pedal{ptype: 34013188, active: true, stype: "Volume/Pan", name: "Volume",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Volume Level", value:0},
            }},
    Pedal{ptype: 34013189, active: true, stype: "Volume/Pan", name: "Pan",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Pan L-R Balance", value:0},
            }},
    /*
     * FX Loop
     */
    Pedal{ptype: 34078720, active: true, stype: "FX Loop", name: "FX Loop",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Send", value:0},
            &PerCentParam{name: "Return", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    /*
     * Pitch section
     */
    Pedal{ptype: 34144258, active: true, stype: "Pitch", name: "Pitch Glide",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            &PerCentParam{name: "Pitch", value:0},
            &NullParam{},
            &NullParam{},
            &NullParam{},
            }},
    Pedal{ptype: 34144259, active: true, stype: "Pitch", name: "Smart Harmony",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Key", value:0},
            &PerCentParam{name: "Shift", value:0},
            &PerCentParam{name: "Scale", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34144260, active: true, stype: "Pitch", name: "Bass Octaver",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Tone", value:0},
            &PerCentParam{name: "Normal", value:0},
            &PerCentParam{name: "Octave", value:0},
            }},
    /*
     * Preamp+EQ section
     */
    Pedal{ptype: 34209807, active: true, stype: "Filter", name: "Tron Up",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &ListParam{name: "Range", value:0, list: []string{"Low", "High"}},
            &ListParam{name: "Type", value:0, list: []string{"LP", "BP", "HP"}},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209808, active: true, stype: "Filter", name: "Tron Down",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &ListParam{name: "Range", value:0, list: []string{"Low", "High"}},
            &ListParam{name: "Type", value:0, list: []string{"LP", "BP", "HP"}},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209809, active: true, stype: "Filter", name: "Seeker",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &ListParam2{name: "Steps", value:0, list: []string{"2 Steps", "3 Steps",
                "4 Steps", "5 Steps", "6 Steps", "7 Steps", "8 Steps", "9 Steps"},
            },
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209810, active: true, stype: "Filter", name: "Obi Wah",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &ListParam{name: "Type", value:0, list: []string{"LP", "BP", "HP"}},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209811, active: true, stype: "Filter", name: "Slow Filter",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &PerCentParam{name: "Speed", value:0},
            &ListParam{name: "Mode", value:0, list: []string{"Up", "Down"}},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209812, active: true, stype: "Filter", name: "Q Filter",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &PerCentParam{name: "Gain", value:0},
            &ListParam{name: "Type", value:0, list: []string{"LP", "BP", "HP"}},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209813, active: true, stype: "Filter", name: "Throbber",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &ListParam{name: "Wave", value:0, list: []string{"Ramp Up", "Ramp Down",
                "Triangle", "Square",
                }},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209814, active: true, stype: "Filter", name: "Spin Cycle",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &PerCentParam{name: "VolSens", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209815, active: true, stype: "Filter", name: "Comet Trails",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &PerCentParam{name: "Gain", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209816, active: true, stype: "Filter", name: "Octisynth",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &PerCentParam{name: "Depth", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209817, active: true, stype: "Filter", name: "Growler",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &PerCentParam{name: "Pitch", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209818, active: true, stype: "Filter", name: "Synth O Matic",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &ListParam{name: "Wave", value:0, list: []string{"Wave 1", "Wave 2",
                "Wave 3", "Wave 4", "Wave 5", "Wave 6", "Wave 7", "Wave 8"},
            },
            &PerCentParam{name: "Pitch", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209819, active: true, stype: "Filter", name: "Attack Synth",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Freq", value:0},
            &ListParam{name: "Wave", value:0, list: []string{"Square", "PWM", "Ramp"}},
            &PerCentParam{name: "Speed", value:0},
            &PerCentParam{name: "Pitch", value:0},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209820, active: true, stype: "Filter", name: "Synth String",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Attack", value:0},
            &PerCentParam{name: "Pitch", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209821, active: true, stype: "Filter", name: "Voice Box",
        params: []Parameter{
            &TempoParam{name: "Speed", value:0},
            &ListParam{name: "Start", value:0, list: []string{"A", "E", "I", "O", "U"}},
            &ListParam{name: "End", value:0, list: []string{"A", "E", "I", "O", "U"}},
            &ListParam{name: "Auto", value:0, list: []string{"Auto 1", "Auto 2", "Auto 3", "Auto 4"}},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209822, active: true, stype: "Filter", name: "V-Tron",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Speed", value:0},
            &ListParam{name: "Start", value:0, list: []string{"A", "E", "I", "O", "U"}},
            &ListParam{name: "End", value:0, list: []string{"A", "E", "I", "O", "U"}},
            &ListParam{name: "Mode", value:0, list: []string{"Up", "Up/Down"}},
            &PerCentParam{name: "Mix", value:0},
            }},
    Pedal{ptype: 34209830, active: true, stype: "Filter", name: "Vocoder",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Input", value:0},
            &PerCentParam{name: "Mic", value:0},
            &PerCentParam{name: "Decay", value:0},
            &NullParam{},
            &PerCentParam{name: "Mix", value:0},
            }},
    /*
     * Preamp+EQ section
     */
    Pedal{ptype: 34340873, active: true, stype: "Preamp+EQ", name: "Graphic EQ",
        params: []Parameter{
            &NullParam{},
            &RangeParam{name: "80Hz", value:0, min: -12, max: 12},
            &RangeParam{name: "220Hz", value:0, min: -12, max: 12},
            &RangeParam{name: "480Hz", value:0, min: -12, max: 12},
            &RangeParam{name: "1.1kHz", value:0, min: -12, max: 12},
            &RangeParam{name: "2.2kHz", value:0, min: -12, max: 12},
            }},
    Pedal{ptype: 34340874, active: true, stype: "Preamp+EQ", name: "Studio EQ",
        params: []Parameter{
            &NullParam{},
            &ListParam2{name: "Low Freq", value:0, list: []string{"75Hz", "150Hz",
                "180Hz", "240Hz", "500Hz", "700Hz", "1000Hz", "1400Hz"}},
            &RangeParam{name: "Low Gain", value:0, min: -11, max: 11},
            &ListParam2{name: "Hi Freq", value:0, list: []string{"200Hz", "300Hz",
                "400Hz", "800Hz", "1500Hz", "3000Hz", "5000Hz", "8000Hz"}},
            &RangeParam{name: "Hi Gain", value:0, min: -11, max: 11},
            &RangeParam{name: "Gain", value:0, min: -18, max: 18},
            }},
    Pedal{ptype: 34340875, active: true, stype: "Preamp+EQ", name: "Parametric EQ",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Lows", value:0},
            &PerCentParam{name: "Highs", value:0},
            &PerCentParam{name: "Freq", value:0},
            &PerCentParam{name: "Q", value:0},
            &PerCentParam{name: "Gain", value:0},
            }},
    Pedal{ptype: 34340876, active: true, stype: "Preamp+EQ", name: "4 Band Shift EQ",
        params: []Parameter{
            &NullParam{},
            &RangeParam{name: "Low", value:0, min: -12, max: 12},
            &RangeParam{name: "Low Mid", value:0, min: -12, max: 12},
            &RangeParam{name: "Hi Mid", value:0, min: -12, max: 12},
            &RangeParam{name: "Hi", value:0, min: -12, max: 12},
            &PerCentParam{name: "Shift", value:0},
            }},
    Pedal{ptype: 34340877, active: true, stype: "Preamp+EQ", name: "Mid Focus EQ",
        params: []Parameter{
            &NullParam{},
            &PerCentParam{name: "Hi Pass Freq", value:0},
            &PerCentParam{name: "Hi Pass Q", value:0},
            &PerCentParam{name: "Low Pass Freq", value:0},
            &PerCentParam{name: "Low Pass Q", value:0},
            &PerCentParam{name: "Gain", value:0},
            }},
    Pedal{ptype: 34340878, active: true, stype: "Preamp+EQ", name: "Vintage Pre",
        params: []Parameter{
            &NullParam{},
            &ListParam{name: "Phase", value:0, list: []string{"0", "180"}},
            &PerCentParam{name: "Gain", value:0},
            &PerCentParam{name: "Output", value:0},
            &FreqParam{name: "Hi Pass Filter", value:0, min: 20, max: 500},
            &FreqKParam{name: "Low pass Filter", value:0, min:5, max: 20},
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

func (p *Pedal) GetParam(id uint16) Parameter {
    if id > p.GetParamLen() {
        return nil
    }
    return p.params[id]
}

func (p *Pedal) GetParams() []Parameter {
    return p.params
}

func (p *Pedal) GetParamID(param Parameter) (error, uint16) {
    for i, _param := range p.params {
        if _param == param {
            return nil, uint16(i)
        }
    }
    return fmt.Errorf("Parameter %s not found", p.GetName()), 0
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

func (p Pedal) LogInfo() {
    log.Printf("Id %d, Pedal Info, Name %s, Type %s, Active %t\n", p.id, p.name,
         p.stype, p.active)
    log.Printf("Parameters:\n")
    for i, param := range(p.params) {
        log.Printf("----%d %s %s\n", i+1, param.GetName(), param.GetValue())
    }
}

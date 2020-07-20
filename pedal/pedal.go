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

type PedalType uint32

const (
    nonePedalType PedalType = 34471935
)

var pedals = map[PedalType]Pedal {
    nonePedalType: Pedal{active: true, ptype: "None", name: "None"},
    /*
     * Dynamics section
     */
    33554443:      Pedal{active: true, ptype: "Dynamics", name: "Red Comp",
        params: []Parameter{
            Parameter{name: "Sustain", value:0},
            Parameter{name: "Level", value:0},
            }},
    33554444:      Pedal{active: true, ptype: "Dynamics", name: "Blue Comp",
        params: []Parameter{
            Parameter{name: "Sustain", value:0},
            Parameter{name: "Level", value:0},
            }},
    33554445:      Pedal{active: true, ptype: "Dynamics", name: "Blue Comp Treble",
        params: []Parameter{
            Parameter{name: "Sustain", value:0},
            Parameter{name: "Level", value:0},
            }},
    33554446:      Pedal{active: true, ptype: "Dynamics", name: "Tube Comp",
        params: []Parameter{
            Parameter{name: "Threshold", value:0},
            Parameter{name: "Level", value:0},
            }},
    33554447:      Pedal{active: true, ptype: "Dynamics", name: "Vetta Comp",
        params: []Parameter{
            Parameter{name: "Sensitivity", value:0},
            Parameter{name: "Level", value:0},
            }},
    33554448:      Pedal{active: true, ptype: "Dynamics", name: "Vetta Juice",
        params: []Parameter{
            Parameter{name: "Amount", value:0},
            Parameter{name: "Level", value:0},
            }},
    33554449:      Pedal{active: true, ptype: "Dynamics", name: "Noise Gate",
        params: []Parameter{
            Parameter{name: "Threshold", value:0},
            Parameter{name: "Decay", value:0},
            }},
    33554450:      Pedal{active: true, ptype: "Dynamics", name: "Boost Comp",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Comp", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33554451:      Pedal{active: true, ptype: "Dynamics", name: "Hard Gate",
        params: []Parameter{
            Parameter{name: "Open Threshold", value:0},
            Parameter{name: "Close Threshold", value:0},
            Parameter{name: "Hold", value:0},
            Parameter{name: "Decay", value:0},
            }},
    /*
     * Delay section
     */
    33685521:      Pedal{active: true, ptype: "Delay", name: "Digital Delay",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685522:      Pedal{active: true, ptype: "Delay", name: "Digital Delay W/Mod",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "ModSpd", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685523:      Pedal{active: true, ptype: "Delay", name: "Stereo Delay",
        params: []Parameter{
            Parameter{name: "L Time", value:0},
            Parameter{name: "L-FDBK", value:0},
            Parameter{name: "R Time", value:0},
            Parameter{name: "R-FDBK", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685524:      Pedal{active: true, ptype: "Delay", name: "Analog Echo",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685525:      Pedal{active: true, ptype: "Delay", name: "Analog W/Mod",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "ModSpd", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685530:      Pedal{active: true, ptype: "Delay", name: "Multi-Head",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Heads 1-2", value:0},
            Parameter{name: "Heads 3-4", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685533:      Pedal{active: true, ptype: "Delay", name: "Low Res Delay",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Res", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685536:      Pedal{active: true, ptype: "Delay", name: "Ping Pong",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Offset", value:0},
            Parameter{name: "Spread", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685537:      Pedal{active: true, ptype: "Delay", name: "Reverse",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "ModSpd", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685538:      Pedal{active: true, ptype: "Delay", name: "Dynamic Delay",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Thresh", value:0},
            Parameter{name: "Ducking", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685539:      Pedal{active: true, ptype: "Delay", name: "Auto-Volume Echo",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685546:      Pedal{active: true, ptype: "Delay", name: "Tube Echo",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Wow/Flt", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685547:      Pedal{active: true, ptype: "Delay", name: "Tube Echo Dry",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Wow/Flt", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685548:      Pedal{active: true, ptype: "Delay", name: "Tape Echo",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685549:      Pedal{active: true, ptype: "Delay", name: "Tape Echo Dry",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685550:      Pedal{active: true, ptype: "Delay", name: "Sweep Echo",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Swp Spd", value:0},
            Parameter{name: "Swp Dep", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685551:      Pedal{active: true, ptype: "Delay", name: "Sweep Echo Dry",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Swp Spd", value:0},
            Parameter{name: "Swp Dep", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685552:      Pedal{active: true, ptype: "Delay", name: "Echo Platter",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Wow/Flt", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33685553:      Pedal{active: true, ptype: "Delay", name: "Echo Platter",
        params: []Parameter{
            Parameter{name: "Time", value:0},
            Parameter{name: "FDBK", value:0},
            Parameter{name: "Wow/Flt", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Mix", value:0},
            }},
    /*
     * Modulation section
     */
    33751072:      Pedal{active: true, ptype: "Modulation", name: "Opto Tremolo",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Shape", value:0},
            Parameter{name: "VolSens", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751073:      Pedal{active: true, ptype: "Modulation", name: "Bias Tremolo",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Shape", value:0},
            Parameter{name: "VolSens", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751074:      Pedal{active: true, ptype: "Modulation", name: "Phaser",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Fdbk", value:0},
            Parameter{name: "Stages", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751075:      Pedal{active: true, ptype: "Modulation", name: "Dual Phaser",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Fdbk", value:0},
            Parameter{name: "LfoShp", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751076:      Pedal{active: true, ptype: "Modulation", name: "Panned Phaser",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Pan", value:0},
            Parameter{name: "PanSpd", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751077:      Pedal{active: true, ptype: "Modulation", name: "U-Vibe",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Fdbk", value:0},
            Parameter{name: "VolSens", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751078:      Pedal{active: true, ptype: "Modulation", name: "Rotary Drum",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751079:      Pedal{active: true, ptype: "Modulation", name: "Rotary Drum/Hrn",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Horn Dep", value:0},
            Parameter{name: "Drive", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751080:      Pedal{active: true, ptype: "Modulation", name: "Analog Flanger",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Fdbk", value:0},
            Parameter{name: "Manual", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751081:      Pedal{active: true, ptype: "Modulation", name: "Jet Flanger",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Fdbk", value:0},
            Parameter{name: "Manual", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751082:      Pedal{active: true, ptype: "Modulation", name: "Analog Chorus",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Ch Vib", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751083:      Pedal{active: true, ptype: "Modulation", name: "Dimension",
        params: []Parameter{
            Parameter{name: "Sw1", value:0},
            Parameter{name: "Sw2", value:0},
            Parameter{name: "Sw3", value:0},
            Parameter{name: "Sw4", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751084:      Pedal{active: true, ptype: "Modulation", name: "Tri Chorus",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Depth2", value:0},
            Parameter{name: "Depth3", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751085:      Pedal{active: true, ptype: "Modulation", name: "Pitch Vibrato",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Rise", value:0},
            Parameter{name: "VolSens", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751086:      Pedal{active: true, ptype: "Modulation", name: "Ring Modulator",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Shape", value:0},
            Parameter{name: "AM/FM", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751087:      Pedal{active: true, ptype: "Modulation", name: "Panner",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Shape", value:0},
            Parameter{name: "VolSens", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751104:      Pedal{active: true, ptype: "Modulation", name: "Barberpole Phaser",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{},
            Parameter{name: "Fdbk", value:0},
            Parameter{name: "Mode", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33751106:      Pedal{active: true, ptype: "Modulation", name: "Frequency Shifter",
        params: []Parameter{
            Parameter{name: "Freq", value:0},
            Parameter{name: "Mode", value:0},
            Parameter{},
            Parameter{},
            Parameter{name: "Mix", value:0},
            }},
    33751107:      Pedal{active: true, ptype: "Modulation", name: "Pattern Tremolo",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Step1", value:0},
            Parameter{name: "Step2", value:0},
            Parameter{name: "Step3", value:0},
            Parameter{name: "Step4", value:0},
            }},
    33751109:      Pedal{active: true, ptype: "Modulation", name: "Script Phase",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            }},
    33751111:      Pedal{active: true, ptype: "Modulation", name: "AC Flanger",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Width", value:0},
            Parameter{name: "Regen", value:0},
            Parameter{name: "Manual", value:0},
            }},
    33751113:      Pedal{active: true, ptype: "Modulation", name: "80A Flanger",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Range", value:0},
            Parameter{name: "Enhance", value:0},
            Parameter{name: "Manual", value:0},
            Parameter{name: "Even Odd", value:0},
            }},
    /*
     * Reverb section
     */
    33816604:      Pedal{active: true, ptype: "Reverb", name: "'63 Spring",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816605:      Pedal{active: true, ptype: "Reverb", name: "Spring",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816606:      Pedal{active: true, ptype: "Reverb", name: "Plate",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816607:      Pedal{active: true, ptype: "Reverb", name: "Room",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816608:      Pedal{active: true, ptype: "Reverb", name: "Chamber",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816609:      Pedal{active: true, ptype: "Reverb", name: "Hall",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816610:      Pedal{active: true, ptype: "Reverb", name: "Ducking",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816611:      Pedal{active: true, ptype: "Reverb", name: "Octo",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816612:      Pedal{active: true, ptype: "Reverb", name: "Cave",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816613:      Pedal{active: true, ptype: "Reverb", name: "Tile",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816614:      Pedal{active: true, ptype: "Reverb", name: "Echo",
        params: []Parameter{
            Parameter{name: "Decay", value:0},
            Parameter{name: "Predelay", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33816615:      Pedal{active: true, ptype: "Reverb", name: "Particle Verb",
        params: []Parameter{
            Parameter{name: "Dwell", value:0},
            Parameter{name: "Condition", value:0},
            Parameter{name: "Gain", value:0},
            Parameter{name: "Mix", value:0},
            }},
    /*
     * Distortion section
     */
    33882122:      Pedal{active: true, ptype: "Distortion", name: "Jet Fuzz",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Fdbk", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Speed", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882123:      Pedal{active: true, ptype: "Distortion", name: "Classic Dist",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Filter", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882124:      Pedal{active: true, ptype: "Distortion", name: "Octave Fuzz",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882125:      Pedal{active: true, ptype: "Distortion", name: "Tube Drive",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882126:      Pedal{active: true, ptype: "Distortion", name: "Screamer",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Tone", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882127:      Pedal{active: true, ptype: "Distortion", name: "Fuzz Pi",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882128:      Pedal{active: true, ptype: "Distortion", name: "Overdrive",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882129:      Pedal{active: true, ptype: "Distortion", name: "Facial Fuzz",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882130:      Pedal{active: true, ptype: "Distortion", name: "Line 6 Distortion",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882131:      Pedal{active: true, ptype: "Distortion", name: "Line 6 Drive",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882132:      Pedal{active: true, ptype: "Distortion", name: "Sub Octave Fuzz",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Sub", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882133:      Pedal{active: true, ptype: "Distortion", name: "Buzz Saw",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882134:      Pedal{active: true, ptype: "Distortion", name: "Heavy Dist",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882135:      Pedal{active: true, ptype: "Distortion", name: "Jumbo Fuzz",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    33882136:      Pedal{active: true, ptype: "Distortion", name: "Color Drive",
        params: []Parameter{
            Parameter{name: "Drive", value:0},
            Parameter{name: "Bass", value:0},
            Parameter{name: "Mid", value:0},
            Parameter{name: "Treble", value:0},
            Parameter{name: "Output", value:0},
            }},
    /*
     * Wah section
     */
    33947659:      Pedal{active: true, ptype: "Wah", name: "Vetta Wah",
        params: []Parameter{
            Parameter{name: "Position", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33947660:      Pedal{active: true, ptype: "Wah", name: "Fassel",
        params: []Parameter{
            Parameter{name: "Position", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33947661:      Pedal{active: true, ptype: "Wah", name: "Weeper",
        params: []Parameter{
            Parameter{name: "Position", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33947662:      Pedal{active: true, ptype: "Wah", name: "Chrome",
        params: []Parameter{
            Parameter{name: "Position", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33947663:      Pedal{active: true, ptype: "Wah", name: "Chrome Custom",
        params: []Parameter{
            Parameter{name: "Position", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33947664:      Pedal{active: true, ptype: "Wah", name: "Throaty",
        params: []Parameter{
            Parameter{name: "Position", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33947665:      Pedal{active: true, ptype: "Wah", name: "Conductor",
        params: []Parameter{
            Parameter{name: "Position", value:0},
            Parameter{name: "Mix", value:0},
            }},
    33947666:      Pedal{active: true, ptype: "Wah", name: "Colorful",
        params: []Parameter{
            Parameter{name: "Position", value:0},
            Parameter{name: "Mix", value:0},
            }},
    /*
     * Volume/Pan section
     */
    34013188:      Pedal{active: true, ptype: "Volume/Pan", name: "Volume",
        params: []Parameter{
            Parameter{name: "Volume Level", value:0},
            }},
    34013189:      Pedal{active: true, ptype: "Volume/Pan", name: "Pan",
        params: []Parameter{
            Parameter{name: "Pan L-R Balance", value:0},
            }},
    /*
     * FX Loop
     */
    34078720:      Pedal{active: true, ptype: "FX Loop", name: "FX Loop",
        params: []Parameter{
            Parameter{name: "Send", value:0},
            Parameter{name: "Return", value:0},
            Parameter{name: "Mix", value:0},
            }},
    /*
     * Pitch section
     */
    34144258:      Pedal{active: true, ptype: "Pitch", name: "Pitch Glide",
        params: []Parameter{
            Parameter{name: "Pitch", value:0},
            Parameter{},
            Parameter{},
            Parameter{},
            Parameter{name: "Mix", value:0},
            }},
    34144259:      Pedal{active: true, ptype: "Pitch", name: "Smart Harmony",
        params: []Parameter{
            Parameter{name: "Key", value:0},
            Parameter{name: "Shift", value:0},
            Parameter{name: "Scale", value:0},
            Parameter{},
            Parameter{name: "Mix", value:0},
            }},
    34144260:      Pedal{active: true, ptype: "Pitch", name: "Bass Octaver",
        params: []Parameter{
            Parameter{name: "Tone", value:0},
            Parameter{name: "Normal", value:0},
            Parameter{name: "Octave", value:0},
            }},
    /*
     * Preamp+EQ section
     */
    34209807:      Pedal{active: true, ptype: "Filter", name: "Tron Up",
        params: []Parameter{
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Range", value:0},
            Parameter{name: "Type", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209808:      Pedal{active: true, ptype: "Filter", name: "Tron Down",
        params: []Parameter{
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Range", value:0},
            Parameter{name: "Type", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209809:      Pedal{active: true, ptype: "Filter", name: "Seeker",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Steps", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209810:      Pedal{active: true, ptype: "Filter", name: "Obi Wah",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Type", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209811:      Pedal{active: true, ptype: "Filter", name: "Slow Filter",
        params: []Parameter{
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Speed", value:0},
            Parameter{name: "Mode", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209812:      Pedal{active: true, ptype: "Filter", name: "Q Filter",
        params: []Parameter{
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Gain", value:0},
            Parameter{name: "Type", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209813:      Pedal{active: true, ptype: "Filter", name: "Throbber",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Wave", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209814:      Pedal{active: true, ptype: "Filter", name: "Spin Cycle",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "VolSens", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209815:      Pedal{active: true, ptype: "Filter", name: "Comet Trails",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Gain", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209816:      Pedal{active: true, ptype: "Filter", name: "Octisynth",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Depth", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209817:      Pedal{active: true, ptype: "Filter", name: "Growler",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Pitch", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209818:      Pedal{active: true, ptype: "Filter", name: "Synth O Matic",
        params: []Parameter{
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Wave", value:0},
            Parameter{name: "Pitch", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209819:      Pedal{active: true, ptype: "Filter", name: "Attack Synth",
        params: []Parameter{
            Parameter{name: "Freq", value:0},
            Parameter{name: "Wave", value:0},
            Parameter{name: "Speed", value:0},
            Parameter{name: "Pitch", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209820:      Pedal{active: true, ptype: "Filter", name: "Synth String",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Freq", value:0},
            Parameter{name: "Attack", value:0},
            Parameter{name: "Pitch", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209821:      Pedal{active: true, ptype: "Filter", name: "Voice Box",
        params: []Parameter{
            Parameter{name: "Speed", value:0},
            Parameter{name: "Start", value:0},
            Parameter{name: "End", value:0},
            Parameter{name: "Auto", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209822:      Pedal{active: true, ptype: "Filter", name: "V-Tron",
        params: []Parameter{
            Parameter{name: "Start", value:0},
            Parameter{name: "End", value:0},
            Parameter{name: "Speed", value:0},
            Parameter{name: "Mode", value:0},
            Parameter{name: "Mix", value:0},
            }},
    34209830:      Pedal{active: true, ptype: "Filter", name: "Vocoder",
        params: []Parameter{
            Parameter{name: "Mic", value:0},
            Parameter{name: "Input", value:0},
            Parameter{},
            Parameter{name: "Decay", value:0},
            Parameter{name: "Mix", value:0},
            }},
    /*
     * Preamp+EQ section
     */
    34340873:      Pedal{active: true, ptype: "Preamp+EQ", name: "Graphic EQ",
        params: []Parameter{
            Parameter{name: "80Hz", value:0},
            Parameter{name: "220Hz", value:0},
            Parameter{name: "480Hz", value:0},
            Parameter{name: "1.1kHz", value:0},
            Parameter{name: "2.2kHz", value:0},
            }},
    34340874:      Pedal{active: true, ptype: "Preamp+EQ", name: "Studio EQ",
        params: []Parameter{
            Parameter{name: "Low Freq", value:0},
            Parameter{name: "Low Gain", value:0},
            Parameter{name: "Hi Freq", value:0},
            Parameter{name: "Hi Gain", value:0},
            Parameter{name: "Gain", value:0},
            }},
    34340875:      Pedal{active: true, ptype: "Preamp+EQ", name: "Parametric EQ",
        params: []Parameter{
            Parameter{name: "Lows", value:0},
            Parameter{name: "Highs", value:0},
            Parameter{name: "Freq", value:0},
            Parameter{name: "Q", value:0},
            Parameter{name: "Gain", value:0},
            }},
    34340876:      Pedal{active: true, ptype: "Preamp+EQ", name: "4 Band Shift EQ",
        params: []Parameter{
            Parameter{name: "Low", value:0},
            Parameter{name: "Low Mid", value:0},
            Parameter{name: "Hi Mid", value:0},
            Parameter{name: "Hi", value:0},
            Parameter{name: "Shift", value:0},
            }},
    34340877:      Pedal{active: true, ptype: "Preamp+EQ", name: "Mid Focus EQ",
        params: []Parameter{
            Parameter{name: "Hi Pass Freq", value:0},
            Parameter{name: "Hi Pass Q", value:0},
            Parameter{name: "Low Pass Freq", value:0},
            Parameter{name: "Low Pass Q", value:0},
            Parameter{name: "Gain", value:0},
            }},
    34340878:      Pedal{active: true, ptype: "Preamp+EQ", name: "Vintage Pre",
        params: []Parameter{
            Parameter{name: "Gain", value:0},
            Parameter{name: "Output", value:0},
            Parameter{name: "Phase", value:0},
            Parameter{name: "Hi Pass Filter", value:0},
            Parameter{name: "Low pass Filter", value:0},
            }},
}

type Parameter struct {
    name  string
    value float32
}

type Pedal struct {
    id     uint32
    active bool
    name   string
    ptype  string
    params []Parameter
    pb     *PedalBoard
    plist  *[]*Pedal
}

func NewNonePedal(id uint32, pb *PedalBoard, plist *[]*Pedal) {
    p := pedals[nonePedalType]
    p.id = id
    p.pb = pb
    p.plist = plist
    *plist = append(*plist, &p)
}

func (p Pedal) PrintInfo() {
    fmt.Printf("Pedal Info\n")
    fmt.Printf("Id %d\n", p.id)
    fmt.Printf("Name %s, type %s\n", p.name, p.ptype)
    fmt.Printf("Active %t\n", p.active)
    fmt.Printf("Parameters:\n")
    for i, param := range(p.params) {
        fmt.Printf("----%d %s %f\n", i, param.name, param.value)
    }
}

func (p *Pedal) SetActive(active bool){
    p.active = active
}

func (p *Pedal) GetParam(id uint32) *Parameter {
    if id >= uint32(len(p.params)) {
        return nil
    }
    return &p.params[id]
}

func (p *Pedal) GetParamLen() uint32 {
    return uint32(len(p.params))
}

func (p *Pedal) remove() {
    for i, _p := range *p.plist {
        if _p == p {
            *p.plist = append((*p.plist)[:i], (*p.plist)[i+1:]...)
            return
        }
    }
}

func (p *Pedal) SetLastPos(pos uint16, ptype uint8) error {
    switch ptype {
    case pedalPosStart:
        p.remove()
        p.pb.start = append(p.pb.start, p)
        p.plist = &p.pb.start
    case pedalPosA:
        p.remove()
        p.pb.pchan.a = append(p.pb.pchan.a, p)
        p.plist = &p.pb.pchan.a
    case pedalPosB:
        p.remove()
        p.pb.pchan.b = append(p.pb.pchan.b, p)
        p.plist = &p.pb.pchan.b
    case pedalPosEnd:
        p.remove()
        p.pb.end = append(p.pb.end, p)
        p.plist = &p.pb.end
    default:
        return fmt.Errorf("Type %d is unsupported for pedal location", ptype)
    }
    return nil
}

func (p *Parameter) SetValue(v float32) {
    if len(p.name) != 0 {
        p.value = v
    }
}

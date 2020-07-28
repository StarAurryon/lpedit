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

type Amp struct {
    id     uint32
    atype  uint32
    active bool
    name   string
    params []Parameter
    pb     *PedalBoard
    plist  *[]PedalBoardItem
    cab    *Cab
}

var amps = []Amp {
    Amp{atype: 524287, active: true, name: "Amp Disabled"},
    Amp{atype: 458752, active: true, name: "phD Motorway",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458753, active: true, name: "Tweed B-Man Normal",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458754, active: true, name: "Tweed B-Man Bright",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458755, active: true, name: "Blackface ‘Lux Normal",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458756, active: true, name: "Blackface ‘Lux Vibrato",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458757, active: true, name: "Blackface Double Normal",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458758, active: true, name: "Blackface Double Vibrato",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458760, active: true, name: "Hiway 100",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458761, active: true, name: "Brit J-45 Normal",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458762, active: true, name: "Brit J-45 Bright",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458763, active: true, name: "Treadplate",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458764, active: true, name: "Brit P-75 Normal",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458765, active: true, name: "Brit P-75 Bright",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458766, active: true, name: "Super O",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458768, active: true, name: "Class A-15",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458769, active: true, name: "Class A-30 TB",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458770, active: true, name: "Divide 9/15",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458772, active: true, name: "Gibtone 185",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458773, active: true, name: "Brit J-800",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458775, active: true, name: "Bomber Uber",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458777, active: true, name: "Angel F-Ball",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458779, active: true, name: "phD Motorway Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458780, active: true, name: "Tweed B-Man Normal Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458781, active: true, name: "Tweed B-Man Bright Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458782, active: true, name: "Blackface ‘Lux Normal Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458783, active: true, name: "Blackface ‘Lux Vibrato Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458784, active: true, name: "Blackface Double Normal Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458785, active: true, name: "Blackface Double Vibrato Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458787, active: true, name: "Hiway 100 Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458788, active: true, name: "Brit J-45 Normal Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458789, active: true, name: "Brit J-45 Bright Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458790, active: true, name: "Treadplate Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458791, active: true, name: "Brit P-75 Normal Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458792, active: true, name: "Brit P-75 Bright Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458793, active: true, name: "Super O Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458795, active: true, name: "Class A-15 Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458796, active: true, name: "Class A-30 TB Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458797, active: true, name: "Divide 9/15 Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458799, active: true, name: "Gibtone 185 Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458800, active: true, name: "Brit J-800 Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458802, active: true, name: "Bomber Uber Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458804, active: true, name: "Angel F-Ball Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458834, active: true, name: "Line 6 Elektrik",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458835, active: true, name: "Line 6 Elektrik Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458837, active: true, name: "Plexi Lead 100 Normal",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458838, active: true, name: "Plexi Lead 100 Normal Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458840, active: true, name: "Plexi Lead 100 Bright",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458841, active: true, name: "Plexi Lead 100 Bright Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458843, active: true, name: "Flip Top",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458844, active: true, name: "Flip Top Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458846, active: true, name: "Solo 100 Clean",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458847, active: true, name: "Solo 100 Clean Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458849, active: true, name: "Solo 100 Crunch",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458850, active: true, name: "Solo 100 Crunch Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458852, active: true, name: "Solo 100 OD",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458853, active: true, name: "Solo 100 OD Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458855, active: true, name: "Line 6 Doom",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458856, active: true, name: "Line 6 Doom Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458858, active: true, name: "Line 6 Epic",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
    Amp{atype: 458859, active: true, name: "Line 6 Epic Preamp",
        params: []Parameter{
            &PerCentParam{name: "Bass", value:0},
            &PerCentParam{name: "Mid", value:0},
            &PerCentParam{name: "Treble", value:0},
            &PerCentParam{name: "Drive", value:0},
            &PerCentParam{name: "Presence", value:0},
            &PerCentParam{name: "Volume", value:0},
            }},
}

func newDisAmp(id uint32, pb *PedalBoard, plist *[]PedalBoardItem, c *Cab) *Amp {
    a := amps[0]
    a.id = id
    a.pb = pb
    a.cab = c
    *plist = append(*plist, PedalBoardItem(&a))
    a.plist = plist
    return &a
}

func newAmp(atype uint32) *Amp {
    for _, newAmp := range amps {
        if newAmp.atype == atype {
            for i := range newAmp.params {
                newAmp.params[i] = newAmp.params[i].Copy()
            }
            return &newAmp
        }
    }
    return nil
}

func (a *Amp) GetActive() bool {
    return a.active
}

func (a *Amp) GetID() uint32 {
    return a.id
}

func (a *Amp) GetName() string {
    return a.name
}

func (a *Amp) GetParam(id uint16) Parameter {
    if id >= uint16(len(a.params)) {
        return nil
    }
    return a.params[id]
}

func (a *Amp) GetParamLen() uint16 {
    return uint16(len(a.params))
}

func (a *Amp) SetActive(active bool){
    a.active = active
}

func (a *Amp) SetLastPos(pos uint16, ctype uint8) error {
    if a.GetID() == 2 {
        return nil
    }
    b, _ := a.pb.GetItem(2).(*Amp)
    switch pos {
    case 0:
        a.remove()
        a.pb.pchan.aAmp = append(a.pb.pchan.aAmp, a)
        a.plist = &a.pb.pchan.aAmp
        b.remove()
        b.pb.pchan.bAmp = append(b.pb.pchan.bAmp, b)
        b.plist = &b.pb.pchan.bAmp
    default:
        p, _ := a.pb.GetPedal(pos).(*Pedal)
        switch p.plist {
        case &a.pb.start:
            a.remove()
            a.pb.startAmp = append(a.pb.startAmp, a)
            a.plist = &a.pb.startAmp
        case &a.pb.end:
            a.remove()
            a.pb.endAmp = append(a.pb.endAmp, a)
            a.plist = &a.pb.endAmp
        default:
            return fmt.Errorf("Wrong metodology in Amp placement")
        }
        b.remove()
        b.pb.bAmp = append(b.pb.bAmp, b)
        b.plist = &a.pb.bAmp
    }
    return nil
}

func (a *Amp) SetType(atype uint32) error{
    _a := newAmp(atype)
    if _a == nil {
        return fmt.Errorf("Amp type not found, code: %d", atype)
    }
    _a.id = a.id
    _a.pb = a.pb
    _a.cab = a.cab
    _a.plist = a.plist
    *a = *_a
    return nil
}

func (a *Amp) remove() {
    *a.plist = nil
}

func (a Amp) LogInfo() {
    a.cab.LogInfo()
    log.Printf("Id %d, Amp Info, Name %s, Active %t\n", a.id, a.name, a.active)
    log.Printf("Parameters:\n")
    for i, param := range(a.params) {
        log.Printf("----%d %s %s\n", i, param.GetName(), param.GetValue())
    }
}

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

type PedalBoardItem interface {
    LPEObject
    GetActive() bool
    GetActive2() uint32
    GetID() uint32
    GetType() uint32
    GetPos() (uint16, uint8)
    LockData()
    SetActive(bool)
    SetPos(uint16, uint8)
    SetPosWithoutCheck(uint16, uint8)
    SetType(uint32) error
    SetType2(string, string)
}

type SortablePosPBI []PedalBoardItem

func (s SortablePosPBI) Len() int           { return len(s) }

func (s SortablePosPBI) Less(i, j int) bool {
    posI, _ := s[i].GetPos()
    posJ, _ := s[j].GetPos()
    return posI < posJ
}

func (s SortablePosPBI) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

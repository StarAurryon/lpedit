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

package message

type ChangeType struct {}

func (ChangeType) StatusActiveChange() int { return 0 }
func (ChangeType) StatusNone() int { return 1 }
func (ChangeType) StatusParameterChange() int { return 2 }
func (ChangeType) StatusParameterChangeMin() int { return 3 }
func (ChangeType) StatusParameterChangeMax() int { return 4 }
func (ChangeType) StatusPresetChange() int { return 4 }
func (ChangeType) StatusPresetLoad() int { return 5 }
func (ChangeType) StatusSetChange() int { return 6 }
func (ChangeType) StatusSetLoad() int { return 7 }
func (ChangeType) StatusTempoChange() int { return 8 }
func (ChangeType) StatusTypeChange() int { return 9 }
func (ChangeType) StatusWarning() int { return 10 }

var ct ChangeType

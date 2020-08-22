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

import "github.com/StarAurryon/lpedit/model/pod/message"

type Signal struct {
    message.ChangeType
}

func (Signal) StatusError() int { return 0 + (1 << 8) }
func (Signal) StatusErrorStop() int { return 1 + (1 << 8) }
func (Signal) StatusInitDone() int { return 2 + (1 << 8) }
func (Signal) StatusNormalStop() int { return 3 + (1 << 8) }
func (Signal) StatusNormalStart() int { return 4 + (1 << 8) }
func (Signal) StatusProgress() int { return 5 + (1 << 8) }

var sg Signal

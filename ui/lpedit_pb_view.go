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

package ui

import "github.com/StarAurryon/qt/core"
import "github.com/StarAurryon/qt/gui"
import "github.com/StarAurryon/qt/widgets"

import "strconv"
import "fmt"
import "unsafe"

import "github.com/StarAurryon/lpedit/model/pod"

type LineView struct {
    *widgets.QFrame
    lpedit *LPEdit
    pos uint16
    posType uint8
}

type PedalView struct {
    *widgets.QLabel
    pbiID uint32
    pos uint16
}

func NewLineView(parent widgets.QWidget_ITF, ff core.Qt__WindowType, s widgets.QFrame__Shape,
        l *LPEdit, pos uint16, posType uint8) *LineView {
    ret := LineView{QFrame: widgets.NewQFrame(parent, ff), lpedit: l, pos: pos, posType: posType}
    ret.SetFrameShape(s)
    ret.SetFrameShadow(widgets.QFrame__Sunken)
    ret.SetMinimumSize2(0, 40)
    ret.ConnectDragEnterEvent(func(event *gui.QDragEnterEvent) {
        event.Accept(event.AnswerRect())
    })
    ret.ConnectDropEvent(ret.dropEventHandler)
    return &ret
}

func NewPedalView(parent widgets.QWidget_ITF, ff core.Qt__WindowType, pbiID uint32, pos uint16) *PedalView {
    ret := PedalView{QLabel: widgets.NewQLabel(parent, ff), pbiID: pbiID, pos: pos}
    ret.ConnectMouseMoveEvent(ret.dragPedal)
    return &ret
}

func (l *LineView) dropEventHandler(event *gui.QDropEvent) {
    ptr, _ := strconv.ParseUint(event.MimeData().Text(), 16, 64)
    p := (*PedalView)(unsafe.Pointer(uintptr(ptr)))

    newPos := l.pos
    newPosType := l.posType
    if l.pos > p.pos {
        newPos--
    }

    l.lpedit.ctrl.SetPedalBoardItemPosition(p.pbiID, newPos, newPosType)
}

func (p *PedalView) dragPedal(event *gui.QMouseEvent) {
    if 0 == (event.Buttons() & core.Qt__LeftButton) { return }
    pixmap := gui.NewQPixmap2(p.Size())

    mimedata := core.NewQMimeData()
    mimedata.SetText(strconv.FormatUint(uint64(uintptr(unsafe.Pointer(p))), 16))

    painter := gui.NewQPainter2(pixmap)
    painter.DrawPixmap10(p.Rect(), p.Grab(core.NewQRect()))
    painter.End()

    drag := gui.NewQDrag(p)
    drag.SetMimeData(mimedata)
    drag.SetPixmap(pixmap)
    drag.SetHotSpot(event.Pos())
    drag.Exec(core.Qt__CopyAction|core.Qt__MoveAction)
}

func (l *LPEdit) generatePedalBoardTopology(pb *pod.PedalBoard) (start []pod.PedalBoardItem,
    bot []pod.PedalBoardItem, top[]pod.PedalBoardItem, end []pod.PedalBoardItem) {

    start = pb.GetItems(pod.PedalPosStart)
    ampA := pb.GetItems(pod.AmpAPos)[0]
    ampB := pb.GetItems(pod.AmpBPos)[0]
    aStart := pb.GetItems(pod.PedalPosAStart)
    aEnd := pb.GetItems(pod.PedalPosAEnd)
    bStart := pb.GetItems(pod.PedalPosBStart)
    bEnd := pb.GetItems(pod.PedalPosBEnd)
    end = pb.GetItems(pod.PedalPosEnd)
    posA, _ := ampA.GetPos()

    fmt.Printf("ID %d, PosA %d\n", ampA.GetID(), posA)

    if posA == 5 { start = append(start, ampA) }
    if posA == 7 { end = append([]pod.PedalBoardItem{ampA}, end...) }

    if posA != 0 {
        top = append(aStart, aEnd...)
        bot = append(bStart, bEnd...)
    } else {
        top = append(append(aStart, ampA), aEnd...)
        bot = append(append(bStart, ampB), bEnd...)
    }

    return
}

func (l *LPEdit) fillPedalBoardView(lay *widgets.QGridLayout, x int, y int, maxX int,
        pbis []pod.PedalBoardItem, initPos uint16, posType uint8) (int, uint16) {
    pos := initPos
    for i := 0; i < (maxX + 1); i++ {
        line := NewLineView(l.PedalBoardView, core.Qt__Widget, widgets.QFrame__HLine, l, pos, posType)
        lay.AddWidget2(line, y, x, 0)
        line.SetAcceptDrops(true)
        x++
        if i < maxX {
            if i < len(pbis) {
                pbi := pbis[i]
                pbiPos, _ := pbi.GetPos()
                pbiUI := NewPedalView(l.PedalBoardView, core.Qt__Widget, pbi.GetID(), pbiPos)
                pbiUI.SetText(pbi.GetName())
                lay.AddWidget2(pbiUI, y, x, 0)
                switch pbi.(type) {
                case *pod.Amp:
                    if posType == pod.PedalPosAStart {
                        posType = pod.PedalPosAEnd
                    } else if posType == pod.PedalPosBStart {
                        posType = pod.PedalPosBEnd
                    }
                default:
                    _, posType = pbi.GetPos()
                    pos++
                }
            } else {
                line := NewLineView(l.PedalBoardView, core.Qt__Widget, widgets.QFrame__HLine, l, pos, posType)
                lay.AddWidget2(line, y, x, 0)
                line.SetAcceptDrops(true)
            }
            x++
        }
    }
    return x, pos
}

func (l *LPEdit) updatePedalBoardView(pb *pod.PedalBoard) {
    for _, item := range l.PedalBoardView.Children(){
        item.DestroyQObject()
    }

    layout := widgets.NewQGridLayout(l.PedalBoardView)
    l.PedalBoardView.SetLayout(layout)

    start, bot, top, end := l.generatePedalBoardTopology(pb)

    j, pos := l.fillPedalBoardView(layout, 0, 1, len(start), start, 0, pod.PedalPosStart)
    max := Max(len(bot), len(top))
    _, pos = l.fillPedalBoardView(layout, j, 2, max, bot, pos, pod.PedalPosBStart)
    j, pos = l.fillPedalBoardView(layout, j, 0, max, top, pos, pod.PedalPosAStart)

    mainMix := widgets.NewQLabel(l.PedalBoardView, core.Qt__Widget)
    mainMix.SetText("Main Mix")
    AddWidget(layout, mainMix, j, 1)
    j++

    l.fillPedalBoardView(layout, j, 1, len(end), end, pos, pod.PedalPosEnd)
}

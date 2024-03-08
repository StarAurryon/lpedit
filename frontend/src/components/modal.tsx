import React, { MouseEventHandler, ReactNode } from "react";

interface ModalType {
    children?: ReactNode;
    isOpen: boolean;
    toggle?: () => void;
  }
  

export const Modal = ({children, isOpen, toggle = () => {}}: ModalType) => {
    const stopPropagate = (e: React.MouseEvent) => {
        e.stopPropagation()
    }

    return (
      <div className={`top-0 left-0 absolute h-screen w-screen bg-black/25 z-10 ${isOpen ? '' : 'hidden'} flex`} onClick={toggle}>
        <div className="bg-gray-700 p-4 rounded-xl m-auto" onClick={stopPropagate}>
          {children}
        </div>
      </div>
    )
}

export default Modal;
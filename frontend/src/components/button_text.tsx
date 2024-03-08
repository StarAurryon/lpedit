import React, { MouseEventHandler } from 'react';
import DefaultButton from './button';

interface Props {
    className?: string;
    textClassName?: string;
    disabled?: boolean;
    onClick?: MouseEventHandler;
    children?: string;
};

export const ButtonText: React.FC<Props> = ({
    className,
    textClassName,
    disabled,
    onClick,
    children,
}) => (
    <DefaultButton className={className} disabled={disabled} onClick={onClick}>
        <span className={`text-white font-semibold ${disabled ? "text-gray-400" : "" } ${textClassName}`}>{children}</span>
    </DefaultButton>
)

export default ButtonText;
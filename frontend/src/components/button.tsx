import { MouseEventHandler } from 'react';

interface Props {
    children?: React.ReactNode;
    className?: string;
    disabled?: boolean;
    onClick?: MouseEventHandler;
};

export const Button: React.FC<Props> = ({
    className,
    children,
    disabled,
    onClick,
}) => (
    <button 
        className={`rounded-full ${className}  ${disabled ? "bg-gray-800" : "bg-gray-600 hover:bg-gray-500"} px-4 py-2`}
        disabled={disabled} 
        onClick={(e) => disabled ? null : onClick ? onClick(e) : ""}
        >
        {children}
    </button>
)

export default Button;
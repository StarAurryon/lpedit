import { ChangeEvent } from 'react';

interface Props {
  className?: string;
  disabled?: boolean;
  onChange?: (value: string) => void;
  onPressEnter? : () => void;
  value? : string;
}

export const Input: React.FC<Props> = ({
  className,
  disabled,
  onChange,
  onPressEnter,
  value,
}) => (
  <input
    className={`rounded-md ${className}  ${disabled ? 'bg-gray-800' : 'bg-gray-800 hover:bg-gray-500'} p-2`}
    disabled={disabled}
    onChange={(e: ChangeEvent<HTMLInputElement>) =>
      disabled
        ? null
        : onChange
          ? onChange(e.target.value)
          : ''
    }
    onKeyUp={(e) => {e.key === "Enter" ? onPressEnter ? onPressEnter() : null : null}}
    value={disabled ? undefined : value}
  />
);

export default Input;

import React, { ReactEventHandler } from 'react';

interface Props<T = string | number> {
  className?: string;
  disabled?: boolean;
  onChange?: (key: string) => void;
  list?: Array<Item<T>>;
  defaultValue?: string;
  value?: string;
}

interface Item<T = string | number> {
  key: T;
  name: string;
}

export const Select = ({
  className = '',
  disabled,
  onChange = () => {},
  list = [],
  value,
  defaultValue = '',
}: Props<string | number>) => (
  <select
    className={`${className} appearance-none p-2 rounded-lg bg-gray-800`}
    disabled={disabled}
    onChange={(i) => onChange(i.target.value)}
    defaultValue={value ? undefined : defaultValue}
    value={value}>
    {list?.map((item) => (
      <option key={item.key} value={item.key} >
        {item.name}
      </option>
    ))}
  </select>
);

export default Select;

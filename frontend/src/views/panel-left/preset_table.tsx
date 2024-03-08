import { usePod } from '@/provider/pod-provider';
import { useEffect, useMemo, useRef } from 'react';

interface TableItem {
  id: number;
  idStr: string;
  name: string;
}

interface Props {
  setId?: number;
}

const PresetTable = ({ setId }: Props) => {
  const ref = useRef<null | HTMLTableRowElement>(null);
  const { pod, setPreset } = usePod();


  const items = useMemo((): TableItem[] => {
    if (setId === undefined) return [];
    if (!pod.sets[setId]) return [];

    return pod.sets[setId].presets.map((preset) => ({
      id: preset.id,
      idStr: preset.idStr,
      name: preset.name,
    }));
  }, [pod, setId]);

  useEffect(() => {
    ref.current?.scrollIntoView({ behavior: 'smooth' });
  }, [items]);

  const handleClick = (id: number) => {
    if (setId === undefined) return;
    setPreset(id, setId);
  };

  return (
    <table className="flex flex-col bg-gray-600 h-0 flex-1">
      <thead>
        <tr className="flex bg-gray-700">
          <th className="border border-slate-500 text-left w-12">Id</th>
          <th className="border border-slate-500 text-left flex-1">Name</th>
        </tr>
      </thead>
      <tbody className="flex-1 h-0 overflow-y-scroll flex flex-col">
        {items.map((item) => {
          const isSelected =
            item.id === pod.currentPresetId && setId === pod.currentSetId;

          return (
            <tr
              className={`flex ${isSelected ? 'bg-gray-700' : 'bg-gray-600'} hover:${isSelected ? 'bg-gray-600' : 'bg-gray-700'}`}
              key={item.id}
              ref={item.id === pod.currentPresetId ? ref : undefined}
              onClick={() => handleClick(item.id)}>
              <td className="border border-slate-500 text-left w-12">
                {item.idStr}
              </td>
              <td className="border border-slate-500 text-left flex-1">
                {item.name}
              </td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
};

export default PresetTable;

import { useEffect, useState } from 'react';
import { ButtonText, Select } from '@/components';
import AboutModal from '../modal/about-modal';
import { usePod } from '@/provider/pod-provider';
import PresetTable from './preset_table';

const PanelBot = () => {
  const [podModalOpen, setPodModalOpen] = useState<boolean>(false);
  const togglePodModal = () => {
    return setPodModalOpen(!podModalOpen);
  };

  const { pod } = usePod();

  const [currentSetIndex, setCurrentSetIndex] = useState<number | undefined>();
  useEffect(() => {
    if (pod.currentSetId === undefined) {
      setCurrentSetIndex(undefined);
      return;
    }
    setCurrentSetIndex(pod.sets[pod.currentSetId].id);
  }, [pod]);

  const handleChange = (key: string) => {
    if (Number.isNaN(key)) return;
    setCurrentSetIndex(Number(key));
  };

  return (
    <div className="flex flex-col h-0 flex-1 gap-4 w-full text-white font-semibold pt-4">
      <div className="flex flex-col gap-1">
        <span>Set List</span>
        <Select
          list={pod?.sets?.map((set) => ({ key: set.id, name: set.name }))}
          value={String(currentSetIndex)}
          onChange={handleChange}></Select>
      </div>
      <div className="flex flex-col h-0 flex-1 gap-1">
        <span className="">Preset</span>
        <PresetTable setId={currentSetIndex} />
      </div>
      <AboutModal isOpen={podModalOpen} toggle={togglePodModal}></AboutModal>
      <ButtonText className="w-full" onClick={togglePodModal}>
        About LPEdit
      </ButtonText>
    </div>
  );
};

export default PanelBot;

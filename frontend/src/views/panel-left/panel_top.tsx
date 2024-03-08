import { useEffect, useMemo, useState } from 'react';
import { ButtonText, Select } from '@/components';

import PodModal from '@/views/modal/pod-modal';
import { usePod } from '@/provider/pod-provider';
import { ReloadPreset, SavePreset } from '@/../wailsjs/go/main/App';
import Input from '@/components/input';

const PanelTop = () => {
  const [podModalOpen, setPodModalOpen] = useState<boolean>(false);
  const togglePodModal = () => {
    return setPodModalOpen(!podModalOpen);
  };

  const { currentPreset, setPedalBoardParameterValue } = usePod();

  const [currentTempo, setCurrentTempo] = useState<undefined | string>();
  const resetTempo = () => {
    setCurrentTempo(currentPreset?.parameters[3].value);
  };

  useEffect(() => resetTempo(), [currentPreset]);

  const handleTempoChange = (value: string) => setCurrentTempo(value);
  const handleTempoSet = () => {
    if (!currentTempo) return;

    const param = currentPreset?.parameters[3];
    if (!param) return;

    const value = parseFloat(currentTempo);
    if (isNaN(value)) {
      resetTempo();
      return;
    }

    if (value > param.max || value < param.min) {
      resetTempo();
      return;
    }

    setPedalBoardParameterValue(3, currentTempo);
  };

  const [tapTempo, setTapTempo] = useState<Date | undefined>();
  const handleTapTempoClick = () => {
    const previousTapTempo = tapTempo;
    const newTapTempo = new Date();
    setTapTempo(newTapTempo);

    if (!previousTapTempo) return;

    const param = currentPreset?.parameters[3];
    if (!param) return;

    const value = 60000 / (newTapTempo.getTime() - previousTapTempo.getTime()) 
    if (value > param.max || value < param.min) return;
    setPedalBoardParameterValue(3, value.toString());
  }

  return (
    <div className="flex flex-col gap-4 w-full text-white font-semibold pb-4">
      <ButtonText className="w-full" onClick={togglePodModal}>
        Manage Pod
      </ButtonText>
      <PodModal isOpen={podModalOpen} toggle={togglePodModal}></PodModal>
      <div className="flex flex-row w-full gap-2">
        <ButtonText className="w-0 flex-grow" onClick={ReloadPreset}>
          Discard changes
        </ButtonText>
        <ButtonText className="w-0 flex-grow" onClick={SavePreset}>
          Save
        </ButtonText>
      </div>
      <div className="flex flex-col gap-1">
        <span>Input 1 Source</span>
        <Select
          list={currentPreset?.parameters[0]?.allowedValue?.map((item) => ({
            key: item,
            name: item,
          }))}
          value={currentPreset?.parameters[0]?.value}
          onChange={(key) => setPedalBoardParameterValue(0, key)}
        />
      </div>
      <div className="flex flex-col gap-1">
        <span>Input 2 Source</span>
        <Select
          list={currentPreset?.parameters[1]?.allowedValue?.map((item) => ({
            key: item,
            name: item,
          }))}
          value={currentPreset?.parameters[1]?.value}
          onChange={(key) => setPedalBoardParameterValue(1, key)}
        />
      </div>
      <div className="flex flex-row w-full gap-2">
        <div className="w-0 flex-grow flex flex-col gap-1">
          <span>Guitar In-Z</span>
          <Select
            list={currentPreset?.parameters[2]?.allowedValue?.map((item) => ({
              key: item,
              name: item,
            }))}
            value={currentPreset?.parameters[2]?.value}
            onChange={(key) => setPedalBoardParameterValue(2, key)}
          />
        </div>
        <div className="w-0 flex-grow flex flex-col gap-1">
          <span>Input Setup</span>
          <Select></Select>
        </div>
      </div>
      <div className="flex flex-col gap-1">
        <span>Tempo</span>
        <div className="flex flex-row gap-4">
          <Input
            className="w-0 flex-1"
            onChange={handleTempoChange}
            onPressEnter={handleTempoSet}
            value={currentTempo}
          />
          <ButtonText className="" onClick={handleTapTempoClick}>Tap Tempo</ButtonText>
        </div>
      </div>
    </div>
  );
};

export default PanelTop;

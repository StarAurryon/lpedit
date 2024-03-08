import {
  ListDevices,
  Start,
  Stop,
  GetCurrentDevice,
} from '@/../wailsjs/go/main/App';
import { useEffect, useState } from 'react';
import { ButtonText, Modal, ProgressBar, Select } from '@/components';
import { usePod } from '@/provider/pod-provider';

interface Props {
  isOpen: boolean;
  toggle: () => void;
}

const podModal = ({ isOpen, toggle }: Props) => {
  const [devices, setDevices] = useState<string[][]>([]);
  const [selectedDevice, setSelectedDevice] = useState<string>('');

  const { isStarted, stopError } = usePod();

  useEffect(() => {
    const setData = async () => {
      setDevices(await ListDevices());
      setSelectedDevice(await GetCurrentDevice());
    };

    setData();
  }, [isOpen]);

  const selectDevice = (value: string) => setSelectedDevice(value);

  const start = () => Start(selectedDevice).then(() => toggle());

  const stop = () => Stop();

  return (
    <Modal isOpen={isOpen}>
      <div className="w-96 flex flex-col gap-2">
        <span>Select your pedal board device</span>
        <Select
          disabled={isStarted}
          list={devices.map((i) => {
            return { key: i[0], name: i[1] };
          })}
          defaultValue={selectedDevice}
          value={isStarted ? selectedDevice : undefined}
          onChange={selectDevice}></Select>
        <span className={stopError ? 'hidden' : 'text-red-600'}>
          {stopError}
        </span>
        <div className="flex flex-row gap-2">
          <ButtonText
            className="w-0 flex-grow"
            disabled={isStarted}
            onClick={start}>
            Start
          </ButtonText>
          <ButtonText
            className="w-0 flex-grow"
            disabled={!isStarted}
            onClick={stop}>
            Stop
          </ButtonText>
          <ButtonText className="w-0 flex-grow" onClick={toggle}>
            Close
          </ButtonText>
        </div>
      </div>
    </Modal>
  );
};

export default podModal;

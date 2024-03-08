import { Modal, ProgressBar } from '@/components';
import { usePod } from '@/provider/pod-provider';

const podModal = () => {
  const {progress} = usePod();

  return (
    <Modal isOpen={progress ? true : false}>
      <div className="w-96 flex flex-col gap-4">
        <span className="text-white font-semibold">Loading your POD data</span>
        <ProgressBar progress={progress ? progress : 0}></ProgressBar>
      </div>
    </Modal>
  );
};

export default podModal;

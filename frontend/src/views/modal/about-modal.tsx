import { Modal } from '@/components';

interface Props {
  isOpen: boolean;
  toggle: () => void;
}

const AboutModal = ({ isOpen, toggle }: Props) => {
  return (
    <Modal isOpen={isOpen} toggle={toggle}>
      <div className="w-96 flex flex-col text-center">
        <span className="pb-4">LPEdit PoC</span>
        <span className="font-normal">Copyright (c) 2020 Nicolas SCHWARTZ</span>
        <span className="font-normal pb-4">Released under GNU GPLv2</span>
        <a className="font-normal hover:font-semibold underline underline-offset-4" href="https://github.com/StarAurryon/lpedit">https://github.com/StarAurryon/lpedit</a>
      </div>
    </Modal>
  );
};

export default AboutModal;

import { PanelLeft, PanelRight } from '@/views';
import { PodProvider } from './provider/pod-provider';
import ProgressModal from './views/modal/progress-modal';

const App = () => {
  return (
    <PodProvider>
      <ProgressModal />
      <div id="App" className="flex flex-row">
        <PanelLeft></PanelLeft>
        <PanelRight></PanelRight>
      </div>
    </PodProvider>
  );
};

export default App;

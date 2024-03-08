import {
  createContext,
  PropsWithChildren,
  useContext,
  useEffect,
  useState,
} from 'react';
import { EventsOn, EventsOff } from '@/../wailsjs/runtime';

import { SetPreset } from '@/../wailsjs/go/main/App';
import { model } from '@/../wailsjs/go/models';
import { SetPedalBoardParameterValue } from '@/../wailsjs/go/main/App';

interface PodStatus {
  isStarted: boolean;
  progress?: number;
  stopError?: string;
  pod: model.Pod;
  currentPreset?: model.Preset;

  setPedalBoardParameterValue: (parameterId: number, value: string) => void;
  setPreset: (presetId: number, setId: number) => void;
}

const PodContext = createContext<PodStatus>({
  isStarted: false,
  pod: new model.Pod(),
  setPedalBoardParameterValue: () => {},
  setPreset: () => {},
});

export const usePod = () => useContext(PodContext);

export const PodProvider = ({ children }: PropsWithChildren) => {
  const [pod, setPod] = useState<model.Pod>(new model.Pod({ sets: [] }));
  const [currentPreset, setCurrentPreset] = useState<
    model.Preset | undefined
  >();
  const [isStarted, setStarted] = useState<boolean>(false);
  const [progress, setProgress] = useState<number | undefined>();
  const [stopError, setStopError] = useState<string | undefined>(undefined);

  const updateSet = (set: model.Set) => {
    console.log(set);

    let newPod = pod;
    if (!newPod.sets) newPod.sets = [];
    if (newPod.sets.length > set.id) {
      newPod.sets[set.id] = set;
    } else {
      newPod.sets.push(set);
    }
    newPod.currentSetId = set.id;
    updateData(newPod);
  };

  const updatePreset = (preset: model.Preset) => {
    let newPod = pod;

    if (!newPod.sets) newPod.sets = [];

    if (newPod.sets.length <= preset.setId) {
      newPod.sets.push(new model.Set({ presets: [] }));
    }

    if (newPod.sets[preset.setId].presets.length > preset.id) {
      newPod.sets[preset.setId].presets[preset.id] = preset;
    } else {
      newPod.sets[preset.setId].presets.push(preset);
    }
    newPod.currentPresetId = preset.id;
    newPod.currentSetId = preset.setId;
    updateData(newPod);
  };

  const updateParameter = (parameter: model.Parameter) => {
    let newPod = pod;
    if (!newPod.currentPresetId || !newPod.currentSetId) return;

    newPod.sets[newPod.currentSetId].presets[newPod.currentPresetId].parameters[
      parameter.id
    ] = parameter;
    updateData(newPod);
  };

  const updateData = (newPod: model.Pod) => {
    setCurrentPreset(
      !newPod.currentPresetId || !newPod.currentSetId
        ? undefined
        : {
            ...newPod.sets[newPod.currentSetId].presets[newPod.currentPresetId],
            convertValues: () => {},
          }
    );

    setPod({ ...newPod, convertValues: () => {} });
  };

  const setPedalBoardParameterValue = (parameterId: number, value: string) => {
    SetPedalBoardParameterValue(parameterId, value);

    let newPod = pod;
    if (!newPod.currentPresetId || !newPod.currentSetId) return;
    newPod.sets[newPod.currentSetId].presets[newPod.currentPresetId].parameters[
      parameterId
    ].value = value;

    updateData(newPod);
  };

  const setPreset = (presetId: number, setId: number) =>
    SetPreset(presetId, setId);

  useEffect(() => {
    EventsOn('start', () => {
      setStarted(true);
      setStopError(undefined);
      setPod(new model.Pod({ sets: [] }));
    });
    EventsOn('stop', (err: string | null) => {
      console.log(`Pod Stopped, err: ${err}`);
      if (err) setStopError(err);
      updateData({ sets: [], convertValues: () => {} });
      setStarted(false);
    });
    EventsOn('statusProgress', (p: number) => {
      setProgress(p);
    });
    EventsOn('initDone', () => {
      setProgress(undefined);
      console.log(pod);
    });
    EventsOn('parameterChange', updateParameter);
    EventsOn('presetChange', updatePreset);
    EventsOn('setChange', updateSet);
    EventsOn('typeChange', () => {
      if (!isStarted) return;
    });

    return () => {
      EventsOff('start');
      EventsOff('stop');
      EventsOff('statusProgress');
      EventsOff('initDone');
      EventsOff('parameterChange');
      EventsOff('presetChange');
      EventsOff('setChange');
      EventsOff('typeChange');
    };
  }, []);

  const podStatus: PodStatus = {
    isStarted,
    progress,
    stopError,

    pod,
    currentPreset,
    setPedalBoardParameterValue,
    setPreset,
  };

  return (
    <PodContext.Provider value={podStatus}>{children}</PodContext.Provider>
  );
};

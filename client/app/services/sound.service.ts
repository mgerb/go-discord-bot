import { SoundType } from '../components/sound-list/sound-list';
import { axios } from './axios.service';

const playSound = (sound: SoundType): Promise<any> => {
  return axios.post('/api/sound/play', { name: sound.name });
};

export const SoundService = {
  playSound,
};

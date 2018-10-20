import { SoundType } from '../components/sound-list/sound-list';
import { ISound } from '../model';
import { axios } from './axios.service';

const playSound = (sound: SoundType): Promise<any> => {
  return axios.post('/api/sound/play', { name: sound.name });
};

const getSounds = async (): Promise<ISound[]> => {
  const res = await axios.get('/api/sound');
  return res.data.data;
};

export const SoundService = {
  getSounds,
  playSound,
};

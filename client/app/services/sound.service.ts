import { ISound, SoundListType, SoundType } from '../model';
import { axios } from './axios.service';

const playSound = (sound: SoundType): Promise<any> => {
  return axios.post('/api/sound/play', { name: sound.name });
};

const getSounds = async (): Promise<ISound[]> => {
  const res = await axios.get('/api/sound');
  return res.data.data;
};

export const playAudioInBrowser = (sound: SoundType, type: SoundListType) => {
  const t = type === 'favorites' ? 'sounds' : type;
  const url = `/public/${t.toLowerCase()}/` + sound.name + '.' + sound.extension;
  const audio = new Audio(url);
  audio.play();
};

export const SoundService = {
  getSounds,
  playSound,
  playAudioInBrowser,
};

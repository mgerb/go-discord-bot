import { SoundType } from '../model';

const clear = () => {
  localStorage.clear();
};

const setJWT = (token: string) => {
  localStorage.setItem('jwt', token);
};

const getJWT = (): string | null => {
  return localStorage.getItem('jwt');
};

const getFavorites = (): SoundType[] => {
  const f = localStorage.getItem('favorites');
  return f ? JSON.parse(f) : [];
};

const setFavorites = (f: SoundType[]): void => {
  localStorage.setItem('favorites', JSON.stringify(f));
};

export const StorageService = {
  clear,
  getJWT,
  setJWT,
  getFavorites,
  setFavorites,
};

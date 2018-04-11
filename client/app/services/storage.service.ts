const clear = () => {
  localStorage.clear();
};

const setJWT = (token: string) => {
  localStorage.setItem('jwt', token);
};

const getJWT = (): string | null => {
  return localStorage.getItem('jwt');
};

export const StorageService = {
  clear,
  getJWT,
  setJWT,
};

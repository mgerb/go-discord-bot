const setJWT = (token: string) => {
  localStorage.setItem('jwt', token);
};

const getJWT = (): string | null => {
  return localStorage.getItem('jwt');
};

export const storage = {
  getJWT,
  setJWT,
};

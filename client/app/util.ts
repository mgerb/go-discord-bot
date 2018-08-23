const getScreenWidth = () => {
  const w = window,
    d = document,
    e = d.documentElement,
    g = d.getElementsByTagName('body')[0];

  return w.innerWidth || e.clientWidth || g.clientWidth;
};

const isMobileScreen = () => {
  return getScreenWidth() < 768;
};

export const Util = {
  getScreenWidth,
  isMobileScreen,
};

import ax from 'axios';

export const axios = ax.create();

axios.interceptors.request.use(
  config => {
    // Do something before request is sent
    return config;
  },
  function(error) {
    // Do something with request error
    return Promise.reject(error);
  },
);

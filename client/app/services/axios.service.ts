import ax from 'axios';
import nprogress from 'nprogress';

nprogress.configure({ showSpinner: false });

export const axios = ax.create();

axios.interceptors.request.use(
  config => {
    nprogress.start();
    // Do something before request is sent
    return config;
  },
  error => {
    // Do something with request error
    return Promise.reject(error);
  },
);

axios.interceptors.response.use(
  config => {
    nprogress.done();
    return config;
  },
  error => {
    nprogress.done();
    return Promise.reject(error);
  },
);

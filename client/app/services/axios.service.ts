import ax from 'axios';
import { StorageService } from './storage.service';

export const axios = ax.create();

axios.interceptors.request.use(
  config => {
    const jwt = StorageService.getJWT();
    if (jwt) {
      config.headers['Authorization'] = `Bearer ${jwt}`;
    }
    // Do something before request is sent
    return config;
  },
  function(error) {
    // Do something with request error
    return Promise.reject(error);
  },
);

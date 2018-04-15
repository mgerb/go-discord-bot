import { axios } from './axios.service';

const redirectUrl = window.location.origin + '/oauth';

const getClientID = async (): Promise<string> => {
  const res = await axios.get('/api/config/client_id');
  return res.data['id'];
};

const getOauthUrl = async (): Promise<string> => {
  const clientID = await getClientID();
  return `https://discordapp.com/api/oauth2/authorize?client_id=${clientID}&redirect_uri=${redirectUrl}&response_type=code&scope=email%20identify`;
};

export const OauthService = {
  getClientID,
  getOauthUrl,
};

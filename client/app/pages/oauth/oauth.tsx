import React from 'react';
import { get } from 'lodash';
import axios from 'axios';
import { storage } from '../../storage';

interface Props {}

interface State {}

export class Oauth extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
  }

  componentDidMount() {
    const code = get(this, 'props.location.query.code');

    if (code) {
      // do stuff here
      this.fetchOauth(code as string);
    }
  }

  private async fetchOauth(code: string) {
    try {
      const res = await axios.post('/api/oauth', { code });
      storage.setJWT(res.data);
      window.location.href = '/';
    } catch (e) {
      console.error(e);
    }
  }

  render() {
    return <div />;
  }
}

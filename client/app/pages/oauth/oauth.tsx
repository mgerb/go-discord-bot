import React from 'react';
import { axios, StorageService } from '../../services';
import queryString from 'query-string';
import { RouteComponentProps } from 'react-router-dom';

interface Props extends RouteComponentProps<any> {}

interface State {}

export class Oauth extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
  }

  componentDidMount() {
    const params = queryString.parse(this.props.location.search);

    if (params['code']) {
      // do stuff here
      this.fetchOauth(params['code']);
    }
  }

  private async fetchOauth(code: string) {
    try {
      const res = await axios.post('/api/oauth', { code });
      StorageService.setJWT(res.data);
      window.location.href = '/';
    } catch (e) {
      console.error(e);
    }
  }

  render() {
    return <div />;
  }
}

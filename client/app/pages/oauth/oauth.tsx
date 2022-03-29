import queryString, { ParsedQuery } from 'query-string';
import React from 'react';
import { RouteComponentProps } from 'react-router-dom';
import { axios, StorageService } from '../../services';

export class Oauth extends React.Component<RouteComponentProps<unknown>, unknown> {
  constructor(props: RouteComponentProps<unknown>) {
    super(props);
  }

  componentDidMount() {
    const params: ParsedQuery<string> = queryString.parse(this.props.location.search);

    if (params['code']) {
      this.fetchOauth(params['code'] as string);
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

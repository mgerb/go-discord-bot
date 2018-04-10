import React from 'react';
import { NavLink } from 'react-router-dom';
import jwt_decode from 'jwt-decode';

import './Navbar.scss';
import { storage } from '../../storage';

let oauthUrl: string;

if (!process.env.NODE_ENV) {
  // dev
  oauthUrl = `https://discordapp.com/api/oauth2/authorize?client_id=410818759746650140&redirect_uri=https%3A%2F%2Flocalhost%2Foauth&response_type=code&scope=identify%20guilds`;
} else {
  // prod
  oauthUrl = `https://discordapp.com/api/oauth2/authorize?client_id=271998875802402816&redirect_uri=https%3A%2F%2Fcashdiscord.com%2Foauth&response_type=code&scope=identify%20guilds%20email`;
}

interface Props {}

interface State {
  token: string | null;
  email?: string;
}

export class Navbar extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      token: null,
    };
  }

  componentDidMount() {
    const token = storage.getJWT();

    if (token) {
      const claims: any = jwt_decode(token!);
      console.log(claims);
      const email = claims['email'];
      this.setState({ token, email });
    }
  }

  private logout = () => {
    localStorage.clear();
    window.location.href = '/';
  };

  render() {
    return (
      <div className="Navbar">
        <div className="Navbar__header">Go Discord Bot</div>
        <NavLink exact to="/" className="Navbar__item" activeClassName="Navbar__item--active">
          Home
        </NavLink>
        <NavLink to="/soundboard" className="Navbar__item" activeClassName="Navbar__item--active">
          Soundboard
        </NavLink>
        <NavLink to="/downloader" className="Navbar__item" activeClassName="Navbar__item--active">
          Youtube Downloader
        </NavLink>
        <NavLink to="/clips" className="Navbar__item" activeClassName="Navbar__item--active">
          Clips
        </NavLink>

        {!this.state.token ? (
          <a href={oauthUrl} className="Navbar__item">
            Login
          </a>
        ) : (
          <a className="Navbar__item" onClick={this.logout}>
            Logout
          </a>
        )}

        {this.state.email && <div className="Navbar__email">{this.state.email}</div>}
      </div>
    );
  }
}

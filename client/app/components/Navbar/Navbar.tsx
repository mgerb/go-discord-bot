import React from 'react';
import { NavLink } from 'react-router-dom';
import jwt_decode from 'jwt-decode';
import { StorageService } from '../../services';
import './Navbar.scss';

const baseUrl = window.location.origin + '/oauth';

const oauthUrl = `https://discordapp.com/api/oauth2/authorize?client_id=410818759746650140&redirect_uri=${baseUrl}&response_type=code&scope=identify%20guilds`;

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
    const token = StorageService.getJWT();

    if (token) {
      const claims: any = jwt_decode(token!);
      const email = claims['email'];
      this.setState({ token, email });
    }
  }

  private logout = () => {
    StorageService.clear();
    window.location.href = '/';
  };

  render() {
    return (
      <div className="Navbar">
        <div className="Navbar__header">Sound Bot</div>
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
        <NavLink to="/stats" className="Navbar__item" activeClassName="Navbar__item--active">
          Stats
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

import React from 'react';
import { NavLink } from 'react-router-dom';
import jwt_decode from 'jwt-decode';
import { OauthService, StorageService } from '../../services';
import './Navbar.scss';

interface Props {}

interface State {
  token: string | null;
  email?: string;
  oauthUrl?: string;
}

export class Navbar extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      token: null,
    };
  }

  componentDidMount() {
    this.loadOauthUrl();
    const token = StorageService.getJWT();

    if (token) {
      const claims: any = jwt_decode(token);
      const email = claims['email'];
      this.setState({ token, email });
    }
  }

  async loadOauthUrl() {
    try {
      const oauthUrl = await OauthService.getOauthUrl();
      this.setState({ oauthUrl });
    } catch (e) {
      console.error(e);
    }
  }

  private logout = () => {
    StorageService.clear();
    window.location.href = '/';
  };

  renderLoginButton() {
    if (!this.state.oauthUrl) {
      return null;
    }

    return !this.state.token ? (
      <a href={this.state.oauthUrl} className="Navbar__item">
        Login
      </a>
    ) : (
      <a className="Navbar__item" onClick={this.logout}>
        Logout
      </a>
    );
  }

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

        {this.renderLoginButton()}

        {this.state.email && <div className="Navbar__email">{this.state.email}</div>}
      </div>
    );
  }
}

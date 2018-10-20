import React from 'react';
import { NavLink } from 'react-router-dom';
import { IClaims, Permissions } from '../../model';
import { OauthService, StorageService } from '../../services';
import './navbar.scss';

interface Props {
  claims?: IClaims;
  open: boolean;
  onNavClick: () => void;
}

interface State {
  oauthUrl?: string;
}

export class Navbar extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {};
  }

  componentDidMount() {
    this.loadOauthUrl();
  }

  async loadOauthUrl() {
    const oauthUrl = await OauthService.getOauthUrl();
    if (oauthUrl) {
      this.setState({ oauthUrl });
    }
  }

  private logout = () => {
    StorageService.clear();
    window.location.href = '/';
  };

  renderLoginButton() {
    const { claims } = this.props;

    if (!this.state.oauthUrl) {
      return null;
    }

    return !claims ? (
      <a href={this.state.oauthUrl} className="navbar__item">
        Login
      </a>
    ) : (
      <a className="navbar__item" onClick={this.logout}>
        Logout
      </a>
    );
  }

  renderNavLink = (title: string, to: string, params?: any) => {
    return (
      <NavLink
        {...params}
        to={to}
        className="navbar__item"
        activeClassName="navbar__item--active"
        onClick={this.props.onNavClick}
      >
        {title}
      </NavLink>
    );
  };

  render() {
    const { claims, open } = this.props;
    const openClass = open ? 'navbar--open' : '';
    return (
      <div className={'navbar ' + openClass}>
        {this.renderNavLink('Soundboard', '/', { exact: true })}
        {this.renderNavLink('Video Archive', '/video-archive')}
        {this.renderNavLink('Youtube Downloader', '/downloader')}
        {this.renderNavLink('Clips', '/clips')}
        {this.renderNavLink('Stats', '/stats')}
        {claims &&
          claims.permissions &&
          claims.permissions === Permissions.Admin &&
          this.renderNavLink('Admin', '/admin')}
        {this.renderLoginButton()}

        {claims && claims.email && <div className="navbar__email">{claims.email}</div>}
      </div>
    );
  }
}

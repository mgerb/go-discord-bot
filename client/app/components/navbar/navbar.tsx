import React from 'react';
import { NavLink } from 'react-router-dom';
import { OauthService, StorageService } from '../../services';
import { AppStore } from '../../stores';
import './navbar.scss';

interface Props {
  appStore: AppStore;
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
    const { claims } = this.props.appStore;

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

  renderNavLink = (title: string, to: string, params?: unknown) => {
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

  renderAuthLinks = () => {
    const { hasAdminPermissions } = this.props.appStore;

    if (hasAdminPermissions()) {
      return (
        <>
          {this.renderNavLink('User Event Log', '/user-event-log')}
          {this.renderNavLink('Users', '/users')}
        </>
      );
    }

    return null;
  };

  render() {
    const { claims, navbarOpen, hasModPermissions } = this.props.appStore;
    const openClass = navbarOpen ? 'navbar--open' : '';
    return (
      <div className={'navbar ' + openClass}>
        {this.renderNavLink('Soundboard', '/', { exact: true })}
        {hasModPermissions() && this.renderNavLink('Upload History', '/upload-history')}
        {this.renderNavLink('Video Archive', '/video-archive')}
        {this.renderNavLink('Youtube Downloader', '/downloader')}
        {this.renderNavLink('Clips', '/clips')}
        {this.renderNavLink('Stats', '/stats')}
        {this.renderAuthLinks()}
        {this.renderLoginButton()}

        {claims && claims.email && <div className="navbar__email">{claims.email}</div>}
      </div>
    );
  }
}

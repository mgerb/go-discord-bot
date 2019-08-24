import jwt_decode from 'jwt-decode';
import { action, observable } from 'mobx';
import { IClaims, Permissions } from '../model';
import { axios, StorageService } from '../services';
import { Util } from '../util';

export class AppStore {
  @observable
  public navbarOpen: boolean = false;
  @observable
  public jwt?: string;
  @observable
  public claims?: IClaims;

  constructor() {
    const jwt = StorageService.getJWT();
    this.setJWT(jwt as string);
    this.initNavbar();
  }

  private initNavbar() {
    if (!Util.isMobileScreen()) {
      this.navbarOpen = true;
    }
  }

  private setJWT(jwt?: string) {
    if (!jwt) {
      return;
    }
    axios.defaults.headers['Authorization'] = `Bearer ${jwt}`;
    this.jwt = jwt;
    const claims = jwt_decode(jwt);
    if (claims) {
      this.claims = claims as IClaims;
    }
  }

  @action
  public toggleNavbar = () => {
    this.navbarOpen = !this.navbarOpen;
  };

  public hasModPermissions = (): boolean => {
    return !!this.claims && this.claims.permissions >= Permissions.Mod;
  };

  public hasAdminPermissions = (): boolean => {
    return !!this.claims && this.claims.permissions >= Permissions.Admin;
  };
}

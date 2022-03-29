import jwt_decode from 'jwt-decode';
import { filter, uniqBy } from 'lodash';
import { action, observable } from 'mobx';
import { IClaims, Permissions, SoundType } from '../model';
import { axios, StorageService } from '../services';
import { Util } from '../util';

export class AppStore {
  @observable
  public navbarOpen = false;
  @observable
  public jwt?: string;
  @observable
  public claims?: IClaims;
  @observable
  private favorites: SoundType[] = [];

  constructor() {
    const jwt = StorageService.getJWT();
    this.favorites = StorageService.getFavorites();
    this.setJWT(jwt as string);
    this.initNavbar();
  }

  private initNavbar() {
    if (!Util.isMobileScreen()) {
      this.navbarOpen = true;
    }
  }

  private setJWT = (jwt?: string) => {
    if (!jwt) {
      return;
    }
    axios.defaults.headers['Authorization'] = `Bearer ${jwt}`;
    this.jwt = jwt;
    const claims = jwt_decode(jwt);
    if (claims) {
      this.claims = claims as IClaims;
    }
  };

  public getFavorites = (): SoundType[] => {
    return this.favorites;
  };

  public addFavorite = (f: SoundType): void => {
    this.favorites = uniqBy([...this.favorites, f], 'name');
    StorageService.setFavorites(this.favorites);
  };

  public removeFavorite = (f: SoundType): void => {
    this.favorites = filter(this.favorites, (fa) => fa.name !== f.name);
    StorageService.setFavorites(this.favorites);
  };

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

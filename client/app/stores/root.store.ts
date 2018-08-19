import { AppStore } from './app.store';

export class RootStore {
  public appStore = new AppStore();
}

export const rootStoreInstance = new RootStore();

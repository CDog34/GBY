import {observable, action} from 'mobx';

export class AuthStore {
  @observable auth = false;
  static instance = null;

  static getInstance() {
    if (AuthStore.instance) return AuthStore.instance;
    const newInstance = new AuthStore();
    AuthStore.instance = newInstance;
    return newInstance;
  }


  @action
  setState(state) {
    this.auth = state
  }

}

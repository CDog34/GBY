import {observable, computed, action} from 'mobx';
import _ from 'lodash';

export class NetworkStore {
  @observable requests = [];
  @observable err = null;
  timer = null;
  static instance = null;

  @computed
  get isBusy() {
    return this.requests.length > 0;
  }

  @computed
  get isError() {
    return this.err !== null;
  }

  @action
  displayError(err) {
    clearTimeout(this.timer);
    this.err = err;
    this.timer = setTimeout(action(() => this.err = null), 5500);
  }

  static getInstance() {
    if (NetworkStore.instance) return NetworkStore.instance;
    const newInstance = new NetworkStore();
    NetworkStore.instance = newInstance;
    return newInstance;
  }


  @action
  requestStart(url) {
    this.requests.push(url)
  }

  @action
  requestComplete(url) {
    _.remove(this.requests, (val) => val === url);
  }

  @action
  requestError(url, err) {
    _.remove(this.requests, (val) => val === url);
    this.displayError(err);
  }


}

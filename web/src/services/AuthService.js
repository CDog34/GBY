import {UserModel} from '../resources/User';
import {LocalStorageService} from '../services/LocalStorageService';
import {Resource} from '../services/Resource';
import {AuthStore} from '../stores/AuthStore';

const store = AuthStore.getInstance();
export class AuthService {

  static async login(userName, password) {
    try {
      const resRaw = await UserModel.login({}, {email: userName, password: password});
      if (resRaw.status === 204) {
        store.setState(true);
        return true;
      }
      const res = await resRaw.json();
      LocalStorageService.set('user', res.user);
      LocalStorageService.set('sessionName', res.name);
      LocalStorageService.set('sessionValue', res.value);
      store.setState(true);
      Resource.setHeader(res.name, res.value);
      return true;
    } catch (err) {
      store.setState(false);
      throw err;
    }
  }

  static async valid() {
    try {
      if (!LocalStorageService.exist('sessionName') || !LocalStorageService.exist('sessionValue')) {
        throw {errorCode: 1, massage: 'noLocalSession'}
      }
      const sessName = LocalStorageService.get('sessionName');
      const sessValue = LocalStorageService.get('sessionValue');
      Resource.setHeader(sessName, sessValue);
      const res = await UserModel.valid();
      const json = await res.json();
      LocalStorageService.set('user', JSON.stringify(json.user));
      store.setState(true);
    } catch (err) {
      LocalStorageService.clear('user');
      LocalStorageService.clear('sessionName');
      LocalStorageService.clear('sessionValue');
      store.setState(false);
      throw err;
    }

  }

  static async logout() {
    await UserModel.logout();
    LocalStorageService.clear('user');
    LocalStorageService.clear('sessionName');
    LocalStorageService.clear('sessionValue');
    store.setState(false);
  }
}

import {UserModel} from '../resources/User';
import {LocalStorageService} from '../services/LocalStorageService';
import {Resource} from '../services/Resource';


export class AuthService {
  static auth = false;

  static async login(userName, password) {
    try {
      const resRaw = await UserModel.login({}, {email: userName, password: password});
      if (resRaw.status === 204) {
        AuthService.auth = true;
        return true;
      }
      const res = await resRaw.json();
      LocalStorageService.set('user', res.user);
      LocalStorageService.set('sessionName', res.name);
      LocalStorageService.set('sessionValue', res.value);
      AuthService.auth = true;
      Resource.setHeader(res.name, res.value);
      return true;
    } catch (err) {
      AuthService.auth = false;
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
      LocalStorageService.set('user', res.user);
      AuthService.auth = true;
    } catch (err) {
      AuthService.auth = false;
      throw err;
    }

  }

  static async logout() {
    await UserModel.logout();
    LocalStorageService.clear('user');
    LocalStorageService.clear('sessionName');
    LocalStorageService.clear('sessionValue');
    AuthService.auth = false;
  }

  static isAuth() {
    return AuthService.auth;
  }

}

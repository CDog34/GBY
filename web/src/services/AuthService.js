import {UserModel} from '../resources/User';
import {LocalStorageService} from '../services/LocalStorageService';
import {Resource} from '../services/Resource';


export class AuthService {
  static async login(userName, password) {
    try {
      const resRaw = await UserModel.login({}, {email: userName, password: password});
      if (resRaw.status === 204) return true;
      const res = await resRaw.json();
      LocalStorageService.set('user', res.user);
      LocalStorageService.set('sessionName', res.name);
      LocalStorageService.set('sessionValue', res.value);
      Resource.setHeader(res.name, res.value);
      return true;
    } catch (err) {
      throw Error(`${err.errorCode}/${err.message}`)
    }
  }

  static async valid() {
    if (!LocalStorageService.exist('sessionName') || !LocalStorageService.exist('sessionValue')) {
      throw {errorCode: 1, massage: 'noLocalSession'}
    }
    const sessName = LocalStorageService.get('sessionName');
    const sessValue = LocalStorageService.get('sessionValue');
    Resource.setHeader(sessName, sessValue);
    const res = await UserModel.valid();
    LocalStorageService.set('User', res.user);
  }

}

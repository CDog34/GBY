import {UserModel} from '../resources/User';
import {LocalStorageService} from '../services/LocalStorageService';
import {Resource} from '../services/Resource';


export class AuthService {
  static async login(userName, password) {
    try {
      const resRaw = await UserModel.login({}, {email: userName, password: password});
      console.log('[Dbg.jq:resRaw]:', resRaw);
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

}

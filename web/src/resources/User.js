import {Resource} from '../services/Resource';

export const UserModel = new Resource('', {
  login: {
    method: Resource.method.post,
    uri: 'login'
  }
});

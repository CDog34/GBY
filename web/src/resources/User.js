import {Resource} from '../services/Resource';

export const UserModel = new Resource('auth', {
  login: {
    method: Resource.method.post,
    uri: '/login'
  }
});

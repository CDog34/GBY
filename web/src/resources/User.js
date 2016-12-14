import {Resource} from '../services/Resource';

export const UserModel = new Resource('auth', {
  login: {
    method: Resource.method.post,
    uri: '/login'
  },
  valid: {
    method: Resource.method.get,
    uri: '/valid'
  },
  logout: {
    method: Resource.method.get,
    uri: '/logout'
  }
});

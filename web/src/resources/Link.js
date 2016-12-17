import {Resource} from '../services/Resource';

export const LinkModel = new Resource('link', {
  list: {
    method: Resource.method.get,
    uri: ''
  }
});

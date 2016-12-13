import {Resource} from '../services/Resource';

export const ArticleModel = new Resource('article', {
  list: {
    method: Resource.method.get,
    uri: ''
  },
  get: {
    method: Resource.method.get,
    uri: '/:articleId'
  }
});

import {Resource} from '../../services/Resource';

export const ArticleAdminModel = new Resource('admin/article', {
  list: {
    method: Resource.method.get,
    uri: ''
  },
  delete: {
    method: Resource.method.del,
    uri: '/:articleId'
  },
  recover: {
    method: Resource.method.get,
    uri: '/:articleId/recover'
  },
  modify: {
    method: Resource.method.put,
    uri: '/:articleId'
  },
  create: {
    method: Resource.method.post,
    uri: ''
  }
});

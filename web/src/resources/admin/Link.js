import {Resource} from '../../services/Resource';

export const LinkAdminModel = new Resource('admin/link', {
  list: {
    method: Resource.method.get,
    uri: ''
  },
  delete: {
    method: Resource.method.del,
    uri: '/:linkId'
  },
  recover: {
    method: Resource.method.get,
    uri: '/:linkId/recover'
  },
  modify: {
    method: Resource.method.put,
    uri: '/:linkId'
  },
  create: {
    method: Resource.method.post,
    uri: ''
  }
});

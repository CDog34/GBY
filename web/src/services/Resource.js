import config from 'config';
import _ from 'lodash';
import {NetworkStore} from '../stores/NetworkStore';

export class Resource {
  static header = new Headers({
    'X-App-Id': 'GBY-WEB'
  });

  static setHeader(key, value) {
    Resource.header.set(key, value);
  }

  static getHeader() {
    return Resource.header;
  }

  static method = {
    post: 'POST',
    get: 'GET',
    put: 'PUT',
    del: 'DELETE'
  };

  static  jsonToQueryString(json) {
    if (!json) return null;
    let isFirst = true;
    return _.map(Object.keys(json), function (key) {
      const prefix = isFirst ? '?' : '';
      isFirst = false;
      return prefix + encodeURIComponent(key) + '=' +
        encodeURIComponent(json[key]);
    }).join('&');
  }

  static parseUri(uri, params) {
    if (!params) return uri;
    let uriArr = uri.split('/');
    uriArr = _.map(uriArr, (path) => {
      if (path[0] === ':') {
        const res = params[path.slice(1)];
        delete params[path.slice(1)];
        return res;
      }
      return path;
    });
    return uriArr.join('/') + Resource.jsonToQueryString(params);
  }

  static generateBodyForJson(json) {
    if (!json) return null;
    return JSON.stringify(json);
  }

  constructor(baseUri, methods, constants) {
    const networkStore = NetworkStore.getInstance();
    this.baseUrl = config.apiBaseUrl + baseUri;
    _.forEach(methods, (value, key) => {
      this[key] = async(uriParams, bodyPayload) => {
        const parsedUri = Resource.parseUri(value.uri, uriParams);
        const option = {
          method: value.method,
          headers: Resource.getHeader()
        };
        if (option.method === Resource.method.post || option.method === Resource.method.put) option['body'] = Resource.generateBodyForJson(bodyPayload);
        networkStore.requestStart(parsedUri);
        try {
          const res = await fetch(this.baseUrl + parsedUri, option);
          if (!res.ok) throw await res.json();
          networkStore.requestComplete(parsedUri);
          return res;
        } catch (err) {
          networkStore.requestError(parsedUri, err);
          throw err;
        }

      }
    });
    _.forEach(constants, (value, key) => this[key] = value)
  }
}

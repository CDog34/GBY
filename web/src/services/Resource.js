import config from 'config';
import _ from 'lodash';

export class Resource {
  static header = new Headers({
    'X-App-Id': 'GBY-WEB'
  });

  static method = {
    post: 'POST',
    get: 'GET'
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
    let data = new FormData;
    data.append('json', JSON.stringify(json));
    return data;
  }

  constructor(baseUri, methods, constants) {
    this.baseUrl = config.apiBaseUrl + baseUri;
    _.forEach(methods, (value, key) => {
      this[key] = async(uriParams, bodyPayload) => {
        const parsedUri = Resource.parseUri(value.uri, uriParams);
        return await fetch(this.baseUrl + parsedUri, {
          method: value.method,
          headers: Resource.header,
          body: Resource.generateBodyForJson(bodyPayload)
        });
      }
    });
    _.forEach(constants, (value, key) => this[key] = value)
  }
}

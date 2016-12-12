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

  constructor(baseUri, methods, constants) {
    this.baseUrl = config.apiBaseUrl + baseUri;
    _.forEach(methods, (value, key) => {
      this[key] = async() => {
        return await fetch(this.baseUrl + value.uri, {
          method: value.method,
          headers: Resource.header
        });
      }
    });
    _.forEach(constants, (value, key) => this[key] = value)
  }
}

export class LocalStorageService {
  static set(key, data) {
    window.localStorage.setItem(key, data);
  }

  static get(key) {
    return window.localStorage.getItem(key);
  }

  static exist(key) {
    return !!window.localStorage.getItem(key)
  }

  static clear(key) {
    window.localStorage.setItem(key, '');
  }
}

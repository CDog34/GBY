export class LocalStorageService {
  static set(key, data) {
    window.localStorage.setItem(key, JSON.stringify(data));
  }

  static get(key) {
    const data = window.localStorage.getItem(key);
    try {
      return JSON.parse(data);
    } catch (err) {
      return data;
    }
  }

  static exist(key) {
    return !!window.localStorage.getItem(key)
  }

  static clear(key) {
    window.localStorage.setItem(key, '');
  }
}

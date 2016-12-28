import uuid from 'uuid';
import {LocalStorageService} from './LocalStorageService';
import config from 'config';
export class VisitorService {
  static generateUuid() {
    return uuid.v4();
  }

  static getVisitor() {
    const instance = JSON.parse(LocalStorageService.get(config.visitorName));
    if (instance && instance.uuid) return instance;
    const newVisitor = {
      uuid: VisitorService.generateUuid()
    };
    LocalStorageService.set(config.visitorName, JSON.stringify(newVisitor));
    return newVisitor;

  }
}

import {LinkAdminModel} from '../resources/admin/Link';
import {LinkModel} from '../resources/Link';

export class LinkService {

  static async adminList() {
    const res = await LinkAdminModel.list();
    return await res.json();
  }

  static async list() {
    const res = await LinkModel.list();
    return await res.json();
  }

  static async adminDelete(linkId, isHard) {
    return await LinkAdminModel.delete({linkId: linkId, hard: isHard ? 1 : 0});
  }

  static async adminRecover(linkId) {
    return await LinkAdminModel.recover({linkId: linkId});
  }

  static async save(link) {
    if (link.id && link.id !== 'new') {
      const res = await LinkAdminModel.modify({linkId: link.id}, link);
      return await res.json();
    } else {
      const res = await LinkAdminModel.create({}, link);
      return await res.json();
    }
  }
}

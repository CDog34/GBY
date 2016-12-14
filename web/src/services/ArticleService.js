import {ArticleModel} from '../resources/Article';
import {ArticleAdminModel} from '../resources/admin/Article';
export class ArticleService {
  static async list() {
    const res = await ArticleModel.list();
    return await res.json();
  }

  static async adminList() {
    const res = await ArticleAdminModel.list();
    return await res.json();
  }

  static async adminDelete(articleId, isHard) {
    return await ArticleAdminModel.delete({articleId: articleId, hard: isHard ? 1 : 0});
  }

  static async adminRecover(articleId) {
    return await ArticleAdminModel.recover({articleId: articleId});
  }

  static async get(articleId) {
    const res = await ArticleModel.get({articleId: articleId});
    return await res.json();
  }

  static async save(article) {
    if (article.id) {
      const res = await ArticleAdminModel.modify({articleId: article.id}, article);
      return await res.json();
    } else {
      const res = await ArticleAdminModel.create({}, article);
      return await res.json();
    }
  }
}

import {ArticleModel} from '../resources/Article';
export class ArticleService {
  static async list() {
    const res = await ArticleModel.list();
    return await res.json();
  }

  static async get(articleId) {
    const res = await ArticleModel.get({articleId: articleId});
    return await res.json();
  }
}

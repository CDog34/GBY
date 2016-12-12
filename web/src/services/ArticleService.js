import {ArticleModel} from '../resources/Article';
export class ArticleService {
  static async list() {
    const res = await ArticleModel.list();
    return res.json();
  }
}

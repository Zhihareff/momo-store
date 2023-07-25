import ApiService from "../api.service";

import BaseResourceService from "./base.service";

abstract class NonPaginatedResourceService<T> extends BaseResourceService<T> {
  query(params?: any): Promise<T[]> {
    return ApiService.get(`${this.path}/`, params).then(data => {
      return data.map((item: any) => new this.resource(item));
    });
  }
}

export default NonPaginatedResourceService;
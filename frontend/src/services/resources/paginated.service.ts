import ApiService from "../api.service";

import BaseResourceService from "./base.service";

import Pagination from "@/models/pagination.model";

abstract class PaginatedResourceService<T> extends BaseResourceService<T> {
  query(params?: any): Promise<Pagination<T>> {
    return ApiService.get(`${this.path}/`, params).then(data => {
      return new Pagination(data, this.resource);
    });
  }
}

export default PaginatedResourceService;
import ApiService from "../api.service";

abstract class BaseResourceService<T>{

  abstract resource: { new(...args: any): T };
  abstract path: string;

  get(slug: string, params?: any): Promise<T> {
    return ApiService.get(`${this.path}/${slug}`, params).then(data => {
      return new this.resource(data);
    });
  }

  post(data?: any) {
    return ApiService.post(`${this.path}/`, data);
  }

  put(slug: string, data: any) {
    return ApiService.put(`${this.path}/${slug}`, data);
  }

  delete(slug: string) {
    return ApiService.delete(`${this.path}/${slug}`);
  }

}

export default BaseResourceService;
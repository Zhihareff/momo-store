import axios, { AxiosRequestHeaders } from 'axios';
import { parse } from 'cookie';

axios.defaults.baseURL = process.env.VUE_APP_API_URL;
class ApiService {

  static csrf_token: string;

  static async getCsrfToken(fetch = true): Promise<string> {
    const cookies = parse(document.cookie);

    if(cookies['mock_store_csrftoken']) {
      return cookies['mock_store_csrftoken'];
    }

    if(!fetch) return Promise.reject('No CSRF token found');

    await this.get('/csrf/');

    return this.getCsrfToken(false);
  }

  static async getHeaders(extra: {[name: string]: string} = {}): Promise<AxiosRequestHeaders> {
    return Object.assign({
      'X-CSRFTOKEN': await this.getCsrfToken(),
    }, extra);
  }

  static get(url: string, params?: any) {
    return axios.get(url, { params }).then(r => r.data);
  }

  static async post(url: string, data?: any) {
    return axios.post(url, data, { headers: await this.getHeaders() }).then(r => r.data);
  }

  static async put(url: string, data?: any) {
    return axios.put(url, data, { headers: await this.getHeaders() }).then(r => r.data);
  }

  static async delete(url: string, data?: any) {
    return axios.delete(url, { headers: await this.getHeaders(), data }).then(r => r.data);
  }
}

export default ApiService;
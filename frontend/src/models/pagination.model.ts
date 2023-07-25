
class Pagination<T> {
  count: number;
  total_pages: number;
  current_page: number;
  next: string | null;
  previous: string | null;
  results: T[];

  constructor(data: any, resource: new(data: any) => T) {
    this.count = data.count;
    this.total_pages = data.total_pages;
    this.current_page = data.current_page;
    this.next = data.next;
    this.previous = data.previous;
    this.results = data.results.map((item: any) => new resource(item));
  }

  static empty<T>(resource: new(data: any) => T) {
    return new Pagination({
      count: 0,
      total_pages: 0,
      current_page: 1,
      next: null,
      previous: null,
      results: []
    }, resource);
  }
}

export default Pagination;
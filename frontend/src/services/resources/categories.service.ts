import NonPaginatedResourceService from "./non_paginated.service";
import CategoryModel from "@/models/category.model";

class CategoriesService extends NonPaginatedResourceService<CategoryModel> {
  resource = CategoryModel;
  path = "categories";
}

export default CategoriesService;
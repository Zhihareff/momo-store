import PaginatedResourceService from "./paginated.service";
import ProductModel from "@/models/product.model";

class ProductsService extends PaginatedResourceService<ProductModel> {
  resource = ProductModel;
  path = "products";
}

export default ProductsService;
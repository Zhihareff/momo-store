import ProductModel from "@/models/product.model";

export type Cart = {
  [key: string]: CartItem
}

export type CartItem = {
  product: ProductModel;
  quantity: number;
}
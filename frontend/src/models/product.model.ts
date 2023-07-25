
export default class ProductModel {

  id: number;
  name: string;
  price: number;
  description: string;
  image: string;

  constructor(data: any) {
    this.id = data.id;
    this.name = data.name;
    this.price = data.price;
    this.description = data.description;
    this.image = data.image;
  };
}
// Inheritance is make child with adding new variabel base from parent element with using extends, and can replace the getInfo

class Clothes {
  brand = "";
  size = "";
  price = 0;

  constructor(brand, size, price) {
    this.brand = brand;
    this.size = size;
    this.price = price;
  }

  getInfo() {
    return `Iam buying ${this.brand} clothes with size ${this.size}, and I paid for it with ${this.price}`;
  }
}

//   Inheritance
class detailClothes extends Clothes {
  constructor(brand, size, price, mallName) {
    super(brand, size, price); //must call the super for calling the parent element
    this.mallName = mallName;
  }

  // override / menimpa -> polymorphism
  getInfo() {
    return `Iam buying ${this.brand} clothes with size ${this.size}, and I paid for it with ${this.price} from ${this.mallName} Mall `;
  }
}

const detailClotes1 = new detailClothes("Uniqlo", "M", 350000, "Cassablanka");
const detailClotes2 = new detailClothes("H&M", "S", 250000, "Margocity");
console.log(detailClotes1.getInfo());
console.log(detailClotes2.getInfo());

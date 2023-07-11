// Abstraction is showing just the important or element that want to show
class Car {
  #brand = ""; //# or hastag is for making private value that cant access out of class Car
  #model = "";
  #price = 0;

  constructor(brand, model, price) {
    this.#brand = brand;
    this.#model = model;
    this.#price = price;
  }

  // getter
  get brand() {
    return this.#brand;
  }

  get model() {
    return this.#model;
  }

  get price() {
    return this.price;
  }
}

const carModel = new Car("Lexus", "Electric Car", 2500);
console.log(carModel.model); // so will show just electric car or way to get datas what we want to choose

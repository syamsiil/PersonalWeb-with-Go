// Encapsulation is for set or change datas from abstraction getter, simply converts the data obtained from getter
class Car {
  #brand = "";
  #model = "";
  #price = 0;

  constructor(brand, model, price) {
    this.#brand = brand;
    this.#model = model;
    this.#price = price;
  }

  get brand() {
    return this.#brand;
  }

  get model() {
    return this.#model;
  }

  get price() {
    return this.#price;
  }

  //   Setter
  set price(value) {
    if (value < 1000) {
      return console.log("Sorry, Not enough money");
    }
    this.#price = this.#price + value;
  }
}

// example the car price is 10000
let myCar = new Car("Lexus", "Electric Car", 9000);
myCar.price = 1200;

console.log(myCar.price);

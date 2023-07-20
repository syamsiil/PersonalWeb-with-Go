class Car {
  color = "";
  price = 0;

  constructor(color, price) {
    this.color = color;
    this.price = price;
  }

  getInfo() {
    return `I have a car with color ${this.color}, and i buy it in ${this.price} `;
  }
}

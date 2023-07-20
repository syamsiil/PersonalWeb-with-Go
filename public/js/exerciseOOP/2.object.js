// class Car {
//   color = "";
//   price = 0;

//   constructor(color, price) {
//     this.color = color;
//     this.price = price;
//   }

//   getInfo() {
//     return `I have a car with color ${this.color}, and i buy it in ${this.price} `;
//   }
// }

// //   Show Value
// const mobil1 = new Car("Red", "$150.000"); // use new for every make the object
// const mobil2 = new Car("Blue", "$250.000");
// console.log(mobil1.getInfo());
// console.log(mobil2.getInfo());

// =================================================
// Make class name
// Define parameters
// Make the constructor
// Make getInfo or anything else u need
// then show the value with make a new variabel
// console the new variabel

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

let clothes1 = new Clothes("The Executive", "M", "$150");
let clothes2 = new Clothes("Uniqlo", "S", "$110");
console.log(clothes1.getInfo());
console.log(clothes2.getInfo());

// =============================================================
// class Mahasiswa {
//   who = "";
//   age = 0;
//   university = "";

//   constructor(who, age, university) {
//     this.who = who;
//     this.age = age;
//     this.university = university;
//   }

//   getInfo() {
//     return `Heloo, Its me ${this.who}, my age ${this.age} years old, and I studied at ${this.university} `;
//   }
// }

// const mahasiswa1 = new Mahasiswa("Syamsi", 25, "Universitas Indonesia");
// const mahasiswa2 = new Mahasiswa("Ali", 25, "Politeknik Negeri Jakarta");
// console.log(mahasiswa1.getInfo());
// console.log(mahasiswa2.getInfo());

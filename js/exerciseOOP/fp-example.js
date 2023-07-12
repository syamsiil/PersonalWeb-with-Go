// Functional Programming (FP)
let x = car1("Lexus", "Electric Car", 2500);
let y = car2("Marcedes Benz", "Electric Car", 5000);

function car1(brand, model, price) {
  return getInfo(brand, model, price);
}

function car2(merk, jenis, harga) {
  return getInfo(merk, jenis, harga);
}

function getInfo(a, b, c) {
  return `I have a car with Brand ${a}, that is ${b} , I paid it with ${c} `;
}

console.log(x);
console.log(y);

// ==================================================
let clothesName1 = clothes1("Uniqlo", "M");
let clothesName2 = clothes2("The Executive", "S");

function clothes1(brand, size) {
  return getInformasi(brand, size);
}

function clothes2(merknya, ukurannya) {
  return getInformasi(merknya, ukurannya);
}

function getInformasi(parameterKe1, parameterKe2) {
  return `Aku baru saja membeli baju ${parameterKe1} dengan ukuran ${parameterKe2}  `;
}

console.log(clothesName1);
console.log(clothesName2);

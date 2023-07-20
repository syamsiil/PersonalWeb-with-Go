// Define a function that take a function as an argument
function getCapture(test) {
  // Invoke the passed function
  test("Seeus");
}

// Invoke the function by passing a function as an argument
getCapture((nama) => {
  console.log("Helloo", nama);
});

// =========================================
function returnFunc() {
  return function (nama) {
    return "Hello" + nama;
  };
}

console.log(returnFunc()("Syams"));

let a = function () {
  return "Hello";
};

console.log(a());

// ==================================
function getBrand(brand) {
  brand("Uniqlo");
}

getBrand((brandName) => {
  console.log("I have an", brandName, "clothes");
});

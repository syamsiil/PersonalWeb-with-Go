// Status promise: pending, fullfill (resolve), rejected
let condition = true;

let janji = new Promise((resolve, reject) => {
  if (condition) {
    resolve("Good promise");
    resolve("Good promise again");
  }
  if (!condition) {
    reject("Promise is canceled");
  }
});

console.log(janji);
janji.then((value) => console.log(value)).catch((err) => console.log(err));

// or can use this
janji
  .then((value) => {
    console.log(value);
  })
  .catch((err) => {
    console.log(err);
  });

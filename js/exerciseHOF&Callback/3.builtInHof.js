// FOREACH -> Looping without return data
const data1 = [5, 4, 3, 2, 1];

data1.forEach((item, index) => {
  //   console.log("item: ", item);
  //   console.log("index: ", index);
});

// MAP -> looping with return data
const data2 = [1, 2, 3, 4, 5];

const double = data2.map((value, index) => {
  return value * 2;
});

console.log(double);

// FILTER -> for create a rating system
const data3 = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

const filterData = data3.filter((value, index) => value % 2 == 1);

console.log(filterData);

// REDUCE -> previous value, current value
const data4 = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

const sumData = data4.reduce((prev, value) => {
  return prev + value;
});
console.log(sumData);

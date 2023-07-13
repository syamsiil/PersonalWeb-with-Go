console.log("Hello World 1");
setTimeout(() => console.log("Hello World Delay 3s"), 3000);
setTimeout(() => {
  return console.log("Hello World Delay 2s ");
}, 2000);
console.log("Hello World 2");

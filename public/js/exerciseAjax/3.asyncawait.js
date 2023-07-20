let condition = true;

let janji = new Promise((resolve, reject) => {
  if (condition) {
    setTimeout(() => {
      resolve("Promise is success");
    }, 3000);
  } else {
    reject("Promise is canceled");
  }
});

async function getData() {
  try {
    const response = await janji;
    console.log(response);
    // alert succes
    Swal.fire("Good job!", response, "success");
  } catch (err) {
    //   alert -> registrasi user gagal
    Swal.fire({
      icon: "error",
      title: "Oops...",
      text: "Registrasi gagal!",
    });
    console.log(err);
  }
}

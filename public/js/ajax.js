// Step
// create const promise
// in new Promise create const xhr for holding datas from json
// get datas with use xhr.open("GET","Your URL npoint json", true) //true its mean for get data with asynchronous way
// validation data with use xhr.onload with create asynchronous func
// create xhr.error for error function
// send datas with use xhr.send
const promise = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest();

  xhr.open("GET", "https://api.npoint.io/585a1820e56fa40db30d", true);
  xhr.onload = function () {
    // http code: 200 -> OK
    if (xhr.status === 200) {
      resolve(JSON.parse(xhr.response));
    } else if (xhr.status >= 400) {
      reject("Error loading data");
    }
  };
  xhr.onerror = function () {
    reject("Network error");
  };
  xhr.send();
});

// Way to show the data have 2 ways
// promise-chaining Way
// promise.then((value) => {
//     console.log(value)
// }).catch((reason) => {
//     console.log(reason)
// })

// Async-await Way

let testimonialData = [];

async function getData(rating) {
  try {
    const response = await promise;
    console.log(response);
    testimonialData = response;
    allTestimonial();
  } catch (err) {
    console(err);
  }
}

getData();

function allTestimonial() {
  let testimonialHTML = "";

  // or can use testimonialData.forEach((card)=>{}) make funtion with arrow function
  testimonialData.forEach(function (card) {
    testimonialHTML += `
      <div class="project-card">
      <img
        src="${card.image}"
        alt="image"
        class="project-card-img"
      />
  
      <p class="project-card-description">
        ${card.quote}
      </p>
  
      <p style="font-weight: 500; margin-top: 15px; text-align: end">
        ${card.user}
      </p>
      <p style="font-weight: 500; margin-top: 15px; text-align: end">${card.rating}<i class="fa-solid fa-star"></i></p>
    </div>`;
  });

  document.getElementById("testimonials").innerHTML = testimonialHTML;
}

// Filter by star
function filteredTestimonial(rating) {
  let filteredTestimonialHTML = "";

  const testimonialFiltered = testimonialData.filter((card) => {
    return card.rating === rating;
  });

  // console.log(testimonialFiltered);
  if (testimonialFiltered.length === 0) {
    filteredTestimonialHTML += `<h1 style="text-align:center" >Data not found! </h1>`;
  } else {
    testimonialFiltered.forEach((card) => {
      filteredTestimonialHTML += `
       <div class="project-card">
       <img
         src="${card.image}"
         alt=""
         class="project-card-img"
       />
   
       <p class="project-card-description">
         ${card.quote}
       </p>
   
       <p style="font-weight: 500; margin-top: 15px; text-align: end">
         ${card.user}
       </p>
       <p style="font-weight: 500; margin-top: 15px; text-align: end">${card.rating}<i class="fa-solid fa-star"></i></p>
     </div>`;
    });
  }

  document.getElementById("testimonials").innerHTML = filteredTestimonialHTML;
}

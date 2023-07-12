// List Task
// Class
// Object
// Inheritance
// Polymorphism
// Abstraction
// Encapsulation

// Step
// Make a Testimonial Class
// Add variabel quote and img
// Add constructor
// Add get
// Add get for testimonial html
// Make Testimonial data container / array of object type

// THIS IS OOP
// class Testimonial {
//   #quote = "";
//   #image = "";

//   constructor(quote, image) {
//     this.#quote = quote;
//     this.#image = image;
//   }

//   get quote() {
//     return this.#quote;
//   }

//   get image() {
//     return this.#image;
//   }

//   get user() {
//     return new Error("there is must be user to make testimonials");
//   }

//   get testimonialHTML() {
//     return `
//     <div class="project-card">
//     <img
//       src="${this.image}"
//       alt=""
//       class="project-card-img"
//     />

//     <p class="project-card-description">
//       ${this.quote}
//     </p>

//     <p style="font-weight: 500; margin-top: 15px; text-align: end">
//       ${this.user}
//     </p>
//   </div>`;
//   }
// }

// // Inheritance
// class UserTestimonial extends Testimonial {
//   #user = "";

//   constructor(user, quote, image) {
//     super(quote, image);
//     this.#user = user;
//   }

//   get user() {
//     return "User: " + this.#user;
//   }
// }

// class CompanyTestimonial extends Testimonial {
//   #company = "";

//   constructor(company, quote, image) {
//     super(quote, image);
//     this.#company = company;
//   }

//   get user() {
//     return "Company: " + this.#company;
//   }
// }

// const testimonials1 = new CompanyTestimonial(
//   "Enough",
//   "If children don't learn religion, what will they pray for their parents with?",
//   "https://media.istockphoto.com/id/1282695693/photo/little-boy-sitting-on-chair-at-the-table.webp?b=1&s=170667a&w=0&k=20&c=SN7888BLbvEl8Yh3tTwjDYA6YHcvDEPZ60eAwcLMUfo="
// );

// const testimonials2 = new UserTestimonial(
//   "Seeme",
//   "If the wife doesn't study religion, what will she raise her children with?",
//   "https://media.istockphoto.com/id/1349298009/photo/muslim-college-students-talking-to-a-group-in-counseling.webp?b=1&s=170667a&w=0&k=20&c=9_lqf5Fn2qL8aLtJaWd9rqQvqrUwY2lK650-mTE-If4="
// );

// const testimonials3 = new UserTestimonial(
//   "Seeus",
//   "If the husband doesn't study religion, what will he teach his wife?",
//   "https://media.istockphoto.com/id/1368004438/photo/shot-of-a-couple-enjoying-a-day-at-the-beach.webp?b=1&s=170667a&w=0&k=20&c=g2OhqvTR0RgGPqDRdilhLZmHLTjbk6Kc9n-haJo1dZQ="
// );

// // console.log(testimonials1, testimonials2, testimonials3);
// let testimonialData = [testimonials1, testimonials2, testimonials3];

// let testimonialHTML = "";

// for (let i = 0; i < testimonialData.length; i++) {
//   testimonialHTML += testimonialData[i].testimonialHTML;
// }

// document.getElementById("testimonials").innerHTML = testimonialHTML;

// =====================================================================
// THIS IS HOF RATING
// Step to make rating system
// Make Variable for Testimonial Container Data Array of object and fill dummy datas
// Make allTestimonials function with use foreach
// Make filteredTestimonial function, in the function make a varible testimonialfiltered for hold the data after use hof filter
// Make a conditional function for filter rating by star with use foreach
// Dont forget to get element with use document.getElementbyId(   )

const testimonialData = [
  {
    user: "Seeus",
    quote:
      "If the husband doesn't study religion, what will he teach his wife?",
    image:
      "https://media.istockphoto.com/id/1368004438/photo/shot-of-a-couple-enjoying-a-day-at-the-beach.webp?b=1&s=170667a&w=0&k=20&c=g2OhqvTR0RgGPqDRdilhLZmHLTjbk6Kc9n-haJo1dZQ=",
    rating: 5,
  },
  {
    user: "Seeme",
    quote:
      "If the wife doesn't study religion, what will she raise her children with?",
    image:
      "https://media.istockphoto.com/id/1349298009/photo/muslim-college-students-talking-to-a-group-in-counseling.webp?b=1&s=170667a&w=0&k=20&c=9_lqf5Fn2qL8aLtJaWd9rqQvqrUwY2lK650-mTE-If4=",
    rating: 4,
  },
  {
    user: "Enough",
    quote:
      "If children don't learn religion, what will they pray for their parents with?",
    image:
      "https://media.istockphoto.com/id/1282695693/photo/little-boy-sitting-on-chair-at-the-table.webp?b=1&s=170667a&w=0&k=20&c=SN7888BLbvEl8Yh3tTwjDYA6YHcvDEPZ60eAwcLMUfo=",
    rating: 4,
  },
];

function allTestimonial() {
  let testimonialHTML = "";

  // or can use testimonialData.forEach((card)=>{}) make funtion with arrow function
  testimonialData.forEach(function (card) {
    testimonialHTML += `
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

  document.getElementById("testimonials").innerHTML = testimonialHTML;
}

allTestimonial();

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

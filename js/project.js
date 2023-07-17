let projectCardData = [];

function projectCard(event) {
  event.preventDefault();

  let projectName = document.getElementById("input-project-name").value;
  let startDate = new Date(document.getElementById("input-start-date").value);
  let endDate = new Date(document.getElementById("input-end-date").value);
  let description = document.getElementById("input-description").value;
  let image = document.getElementById("input-image").files;

  //Alert if endDate > Today
  let today = new Date().toISOString().split("T")[0];
  if (endDate > today) {
    return alert("No time travel here!");
  }

  const nodeJsIcon = '<i class="fa-brands fa-node-js"></i>';
  const reactJsIcon = '<i class="fa-brands fa-react"></i>';
  const javascriptIcon = '<i class="fa-brands fa-square-js"></i>';
  const html5Icon = '<i class="fa-brands fa-html5"></i>';

  // Alert if no one choose the technologies
  let multiInput = document.querySelectorAll(".multi-input:checked");
  if (multiInput.length === 0) {
    return alert("Select at least one technology use!");
  }

  let nodeJsIconDecide = document.getElementById("input-nodejs").checked
    ? nodeJsIcon
    : "";
  let reactJsIconDecide = document.getElementById("input-reactjs").checked
    ? reactJsIcon
    : "";
  let javascriptIconDecide = document.getElementById("input-javascript").checked
    ? javascriptIcon
    : "";
  let html5IconDecide = document.getElementById("input-html5").checked
    ? html5Icon
    : "";

  //change image to URL
  image = URL.createObjectURL(image[0]);
  console.log(image);

  // Alert if the value startDate>endDate
  const sDvalidation = new Date(startDate);
  const eDvalidation = new Date(endDate);
  if (sDvalidation > eDvalidation) {
    return alert("Input your dates correctly!");
  }

  let previewCard = {
    projectName,
    startDate,
    endDate,
    description,
    image,
    nodeJsIconDecide,
    reactJsIconDecide,
    javascriptIconDecide,
    html5IconDecide,
    postAt: new Date(),
    author: "Seeus",
  };

  projectCardData.push(previewCard);
  console.log(projectCardData);

  renderCard();

  // reset/clear the input field after submit
  document.getElementById("input-project-name").value = "";
  document.getElementById("input-start-date").value = "";
  document.getElementById("input-end-date").value = "";
  document.getElementById("input-description").value = "";
  document.getElementById("input-nodejs").checked = false;
  document.getElementById("input-reactjs").checked = false;
  document.getElementById("input-javascript").checked = false;
  document.getElementById("input-html5").checked = false;
  document.getElementById("input-image").value = "";

  // document.getElementById("project-form").reset();
}

function renderCard() {
  document.getElementById("contents").innerHTML = "";

  for (let index = 0; index < projectCardData.length; index++) {
    const sDate = new Date(projectCardData[index].startDate);
    // console.log(sDate);
    const eDate = new Date(projectCardData[index].endDate);
    // console.log(endDate);
    const differenceTime = eDate - sDate;
    const timeUnits = [
      { value: 365.25 * 24 * 60 * 60 * 1000, label: "years" },
      { value: 30 * 24 * 60 * 60 * 1000, label: "months" },
      { value: 7 * 24 * 60 * 60 * 1000, label: "weeks" },
      { value: 24 * 60 * 60 * 1000, label: "days" },
    ];

    let distanceTime = "";
    for (let calculation = 0; calculation < timeUnits.length; calculation++) {
      const { value, label } = timeUnits[calculation];
      const calculate = Math.floor(differenceTime / value);
      if (calculate > 0) {
        distanceTime = `${calculate} ${label}`;
        break;
      }
    }

    if (distanceTime === "") {
      distanceTime = "today";
    }

    document.getElementById("contents").innerHTML += `
    <div class="project-card">
          <img src="${
            projectCardData[index].image
          }" alt="" class="project-card-img" />
          <a href="detail-project.html" class="project-card-title">${
            projectCardData[index].projectName
          }</a>
          <p class="project-card-duration">${convertDate(
            projectCardData[index].postAt
          )} | ${distanceTime} </p>
          <p class="project-card-description">
          ${projectCardData[index].description}
          </p>

          <div class="project-card-app">
          ${projectCardData[index].nodeJsIconDecide}
          ${projectCardData[index].reactJsIconDecide}
          ${projectCardData[index].javascriptIconDecide}
          ${projectCardData[index].html5IconDecide}
         
          </div>

          <div class="project-card-btn">
            <button class="btn-project">Edit</button>
            <button class="btn-project">Delete</button>
          </div>
        </div>`;
  }
}

function convertDate() {
  // surya : new Date()
  let date = new Date();

  // TANGGAL
  const tanggal = date.getDate();

  // BULAN
  const listBulan = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];

  // bulan agustus
  // console.log(listBulan[7])

  // bulan sesuai dengan bulan ini
  // console.log(listBulan[date.getMonth()])
  const bulan = listBulan[date.getMonth()];

  const year = date.getFullYear();

  let hours = date.getHours();

  let minutes = date.getMinutes();

  if (hours < 10) {
    hours = "0" + hours; // 0-9 -> 00, 01, 02, .. 09 -> 10, 11, 12
  }

  if (minutes < 10) {
    minutes = "0" + minutes; // 0-9 00, 01, 02, .. 09
  }

  // 5 Jul 2023 09:34 WIB
  return `${tanggal} ${bulan} ${year} ${hours}:${minutes} WIB`;
}

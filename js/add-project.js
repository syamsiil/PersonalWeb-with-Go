// Time check
// const date = new Date();
// console.log(date);
// console.log("Date check: ", date.getDate());
// console.log("Hour check: ", date.getHours());
// console.log("UTC hour check : ", date.getUTCHours()); // international time
// console.log("Day check  : ", date.getDay());
// console.log("Year check : ", date.getFullYear());
// console.log("Month check : ", date.getMonth());
// console.log("Seconds check : ", date.getSeconds());
// console.log("Minute check : ", date.getMinutes());

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

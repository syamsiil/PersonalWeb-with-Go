function submitData() {
  let name = document.getElementById("input-name").value;
  let email = document.getElementById("input-email").value;
  let phoneNumber = document.getElementById("input-number").value;
  let subject = document.getElementById("input-subject").value;
  let message = document.getElementById("input-message").value;

  if (name == "") {
    return alert("Name Required!");
  } else if (email == "") {
    return alert("Email Required!");
  } else if (phoneNumber == "") {
    return alert("Phone Number Required");
  } else if (subject == "") {
    return alert("Subject Required!");
  } else if (message == "") {
    return alert("Message is Required!");
  }

  let emailReceiver = "";

  //<a href =`mailto:muhammadalisyamsi@gmail.com?subject?=frontend&body=Hello, Its me Seeus, what can i do for u?. Lets connect with me to 08211112233, thank you`></a>

  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Hello, Its me ${name}, ${message}. What can i do for you?. Lets connect with me to ${phoneNumber}. Thankyou.`;

  a.click();

  console.log(name);
  console.log(email);
  console.log(phoneNumber);
  console.log(subject);
  console.log(message);

  let emailer = {
    name,
    email,
    phoneNumber,
    subject,
    message,
  };

  console.log(emailer);
}

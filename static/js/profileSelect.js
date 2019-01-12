let submit = document.getElementById("submit");
let form = document.querySelector("form");
var submission;
let reader = new FileReader();
let propic = document.getElementById("propic");
let newPic;
let picSubmit = document.getElementById("picsubmit");
//when clicking the submit button all the values from the inputs
//should be collected and packaged so that they can be sent to the server.
submit.onclick = function() {
  submission = {
    "type": "command",
    "commandType": "profile",
    "pform": {
      "fname": form.elements[0].value,
      "lname": form.elements[1].value,
      "uname": form.elements[2].value,
      "email": form.elements[3].value,
      "gender": form.elements[4].selectedOptions[0].value,
      "orientation": form.elements[5].selectedOptions[0].value,
      "interests": form.elements[6].value.split(" ")
    }
  }
  let subs = JSON.stringify(submission);
  ws.send(subs);
}

propic.onchange = function() {
  reader.readAsDataURL(propic.files[0]);
}

picsubmit.onclick = function() {
  submission = {
    "type": "command",
    "commandType": "propic",
    "command": reader.result
  }
  let subs = JSON.stringify(submission);
  ws.send(subs);
}

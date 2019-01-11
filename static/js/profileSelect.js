var gender;
var gender_choice;
var orientation;
var orientation_choice;
let submit = document.getElementById("submit");
let form = document.querySelector("form");
gender = document.getElementById("gender");
orientation = document.getElementById("orientation");
//when clicking the submit button all the values from the inputs
//should be collected and packaged so that they can be sent to the server.
submit.onclick = function() {
  let submission = {
    "type": "command",
    "commandType": "submission"
  }
  gender_choice = gender.selectedOptions[0].text;
  orientation_choice = orientation.selectedOptions[0].text;
}

window.onload = function() {
  let login = document.getElementById("login");
  let register = document.getElementById("register");
  let close_login = document.querySelector("#login_modal .delete");
  let close_register = document.querySelector("#register_modal .delete");
  let modal = document.querySelectorAll(".modal");
  let reg_email = document.getElementById("reg_email");
  let forgotPassword = document.getElementById("forgotPassword");

  forgotPassword.onclick = function() {
    console.log("request password reset");
  }
  function varifyEmail(email) {
    //regular expression that will match most emails and tell the user if that account is valid
    var regex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    var found = email.match(regex);
    return (found);
  }

  function varifyUsername(username) {
    //regular expression to varify if username is legit
    var regex = /^[A-Za-z0-9]+(?:[_-][A-Za-z0-9]+)*$/;
    var found = username.match(regex);
    return (found);
  }

  function varifyPassword(password) {
    //check the length of the password
    if (password.length < 8)
      return (null);
    var numReg = /\d/;
    if (password.match(numReg) == null)
      return (null);
    var upCase = /[A-Z]/;
    if (password.match(upCase) == null)
      return (null);
    var loCase = /[a-z]/;
    if (password.match(loCase) == null)
      return (null);
    return (true);
  }
  
  login.onclick = function(){
    modal[0].classList.add("is-active");
  }

  register.onclick = function(){
    modal[1].classList.add("is-active");
  }
  
  close_login.onclick = function(){
    modal[0].classList.remove("is-active");
  }

  close_register.onclick = function(){
    modal[1].classList.remove("is-active");
  }

  function closeReg() {
    modal[1].classList.remove("is-active");
  }
  const xhr = new XMLHttpRequest();

  xhr.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
      let username = document.getElementById("reg_username");
      let email = document.getElementById("reg_email");
      let response = this.responseText.split(" ");
      if("Fail" == response[0]) {
        if (response[1] == "email")
          email.classList.add("is-danger");
        else
          username.classList.add("is-danger");
      } else {
        email.classList.remove("is-danger");
        email.value = "";
        username.classList.remove("is-danger");
        username.value = "";
        closeReg()
      }
    }
  }

  let register_button = document.getElementById("register_button");
  register_button.onclick = function() {
    let username = document.getElementById("reg_username").value;
    let email = document.getElementById("reg_email").value;
    let password = document.getElementById("reg_password").value;
    if (varifyEmail(email) && varifyUsername(username) && varifyPassword(password)) {
      const url = location.protocol + "//" + location.host + "/users";
      var user = JSON.stringify({
        "username": username,
        "email": email,
        "password": password
      });
      xhr.open("POST", url);
      xhr.setRequestHeader("Content-Type", "application/json");
      xhr.send(user);
      let pw = document.getElementById("reg_password");
      pw.classList.remove("is-danger");
      pw.value = "";
      let mail = document.getElementById("reg_email");
      mail.classList.remove("is-danger");
      let uname = document.getElementById("reg_username");
      uname.classList.remove("is-danger");
    }
    
    if (varifyPassword(password) == null) {
      let pw = document.getElementById("reg_password");
      pw.classList.add("is-danger")
    } else if (varifyPassword(password) != null) {
      let pw = document.getElementById("reg_password");
      pw.classList.remove("is-danger")
    }
    
    if (varifyEmail(email) == null) {
      let mail = document.getElementById("reg_email");
      mail.classList.add("is-danger");
    } else if (varifyEmail(email) == null) {
      let mail = document.getElementById("reg_email");
      mail.classList.remove("is-danger");
    }
    
    if (varifyUsername(username) == null) {
      let uname = document.getElementById("reg_username");
      uname.classList.add("is-danger");
    } else if (varifyUsername(username) == null) {
      let uname = document.getElementById("reg_username");
      uname.classList.remove("is-danger");
    }
  }
  const request = new XMLHttpRequest();
  let loginBtn = document.getElementById("login_button");
  request.onreadystatechange = function() {
    /*if (this.readyState == 4 && this.status == 200) {
      var obj = JSON.parse(this.responseText);
      if (obj["success"] == true) {
        console.log("signed in");
      } else {
        console.log("login failed");
      }
    }*/
    console.log(this.responseText);
  }
  loginBtn.onclick = function() {
    const url = location.protocol + "//" + location.host + "/users/login";
    let login_password = document.getElementById("login_password");
    let login_username = document.getElementById("login_username");


    let user = JSON.stringify({
      "username": login_username.value,
      "password": login_password.value
    });
    request.open("POST", url);
    request.setRequestHeader("Content-Type", "application/json");
    request.send(user);
    console.log(user);
  }
}

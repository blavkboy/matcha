window.onload = function() {
  let login = document.getElementById("login");
  let register = document.getElementById("register");
  let close_login = document.querySelector("#login_modal .delete");
  let close_register = document.querySelector("#register_modal .delete");
  let modal = document.querySelectorAll(".modal");
  let reg_email = document.getElementById("reg_email");
  
  reg_email.onchange = function() {
    //regular expression that will match most emails and tell the user if that account is valid
    var regex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    var found = reg_email.value.match(regex);
    
    if (reg_email.classList.contains("is-success") == true && found == null)
      reg_email.classList.remove("is-success");
    else if (reg_email.classList.contains("is-success") == false && found != null)
      reg_email.classList.add("is-success");
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
}

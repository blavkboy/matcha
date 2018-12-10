window.onload = function() {
  let login = document.getElementById("login");
  let register = document.getElementById("register");
  let close = document.querySelector(".modal .delete");
  let modal = document.querySelectorAll(".modal");
  let log_username = document.getElementById("log_username");
  
  log_username.onchange = function() {
    //regular expression that will match most emails and tell the user if that account is valid
    var regex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    var found = log_username.value.match(regex);
    
    if (log_username.classList.contains("is-success") == true && found == null)
      log_username.classList.remove("is-success");
    else if (log_username.classList.contains("is-success") == false && found != null)
      log_username.classList.add("is-success");
  }
  
  login.onclick = function(){
    modal[0].classList.add("is-active");
  }
  
  close.onclick = function(){
    modal[0].classList.remove("is-active");
  }
}

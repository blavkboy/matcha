window.onload = function() {
  let login = document.getElementById("login");
  let register = document.getElementById("register");
  let close = document.querySelector(".modal .delete");
  let  modal = document.querySelectorAll(".modal");
  login.onclick = function(){
    modal[0].classList.add("is-active");
  }
  close.onclick = function(){
    modal[0].classList.remove("is-active");
  }
}

window.onload = function() {
  //update geolocation
  console.log("Loaded page");
  let navbar = document.querySelector(".navbar-burger");
  let navmenu = document.querySelector(".navbar-menu");

  navbar.onclick = function() {
    if (navbar.classList.contains("is-active")) {
      navbar.classList.remove("is-active");
      navmenu.classList.remove("is-active");
    } else {
      navbar.classList.add("is-active");
      navmenu.classList.add("is-active");
    }
  }
}

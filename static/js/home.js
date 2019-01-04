window.onload = function() {
  //update geolocation
  function showPosition(position) {
    let local = document.getElementById("location");
    local.innerHTML = "latitude: " + position.coords.latitude + " longitude: "
      + position.coords.longitude;
  }
  function getLocation() {
    let local = document.getElementById("location");
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(showPosition);
    } else {
      local.innerHTML = "Geolocation is not supported on this browser";
    }
  }
  getLocation();
}

document.addEventListener("DOMContentLoaded", function(event) {
  $('.btn-logout').click(function(e) {
    Cookies.remove('auth-session');
  });
})
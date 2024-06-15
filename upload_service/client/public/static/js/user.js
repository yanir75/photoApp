// document.addEventListener("DOMContentLoaded", function(event) {
  // doc = document.getElementById('logout')
  // console.log(doc)
//   // doc.addEventListener(function(e) {
//   //   console.log("here")
//   //   Cookies.remove('auth-session');
//   // });
// })

if (document.readyState !== 'loading') {
  // console.log('document is already ready, just execute code here');
  myInitCode();
} else {
  document.addEventListener('DOMContentLoaded', function () {
      // console.log('document was not ready, place code here');
      myInitCode();
  });
}

function myInitCode() {
  doc = document.getElementById('logout')
  doc.onclick = () =>{
    Cookies.remove('auth-session');
  }
}


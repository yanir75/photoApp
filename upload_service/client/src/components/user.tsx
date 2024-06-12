import useScript from "./useScript";

// Usage inside a component
function User(){
		useScript("http://code.jquery.com/jquery-3.1.0.min.js")
	  useScript('/public/js/js.cookie.js');
	  useScript("/public/js/user.js")
	  const doc = document.getElementById("content");
	  let parsed = null;
	  if (doc) {
		let val = doc.getAttribute("vals");
		// console.log(val)
		if (val) {
		  parsed = JSON.parse(val);
		  // console.log(parsed)
		}
	  }  
	 var permstring = "["
	 parsed["permissions"].map((permissions: string) => permstring += permissions+ ", ")
	 permstring = permstring.substring(0,permstring.length-2)
	 permstring+="]"
    return (
<>
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  {/* font awesome from BootstrapCDN */}
  <link
    href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"
    rel="stylesheet"
  />
  <link
    href="//maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"
    rel="stylesheet"
  />
  <link href="/public/css/app.css" rel="stylesheet" />
  <div className="container">
    <div className="login-page clearfix">
      <div className="logged-in-box auth0-box logged-in">
        <h1 id="logo">
          <img src="/public/media/auth0_logo_final_blue_RGB.png" />
        </h1>
        <img className="avatar" src={parsed["picture"]} />
        <h2 style={{color: "black"}}>
          Welcome {parsed["nickname"]}
        </h2>
        <h2 style={{color: "black"}}>
          Your permissions are {permstring}
        </h2>
        <a
          id="qsLogoutBtn"
          className="btn btn-primary btn-lg btn-logout btn-block"
          href="/logout"
        >
          Logout
        </a>
      </div>
    </div>
  </div>
</>

    )
}


export default User;
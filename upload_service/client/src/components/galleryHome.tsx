import ResponsiveAppBar from "./menu";

function galleryHome() {
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
  
    
      return (
        <>
        <ResponsiveAppBar/>
        <div className="wrapper">
          {parsed ? Object.keys(parsed).map((item: any)=> 
          <div className="container">
          <img src={parsed[item]} title={item}/>
          <div className="centered">{item}</div>      
        </div>
        ): null}
          
          </div>
          </>
      )
  }

  export default galleryHome;

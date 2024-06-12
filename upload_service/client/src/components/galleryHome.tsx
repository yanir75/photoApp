import ResponsiveAppBar from "./menu";
import { useNavigate } from "react-router-dom";
function GalleryHome() {
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
    const navigate = useNavigate()
    function goPage(page:string,keys:string[]){
        navigate("/country/".concat(page),{state: {keys: keys}})
    }
    const parsedKeys = Object.keys(parsed)

      return (
        <>
        <ResponsiveAppBar countries={parsedKeys}/>
        <div className="wrapper">
          {parsed ? parsedKeys.map((item: string)=> 
          <div className="container" onClick={()=> goPage(item,parsedKeys)}>
          <img src={parsed[item]} title={item}/>
          <div className="centered">{item}</div>      
        </div>
        ): null}
          
          </div>
          </>
      )
  }

  export default GalleryHome;

import ResponsiveAppBar from "./menu";
import { useCountries } from "./useScript";

function GalleryHome() {
    // const doc = document.getElementById("content");
    // let parsed = null;
    // if (doc) {
    //   let val = doc.getAttribute("vals");
    //   // console.log(val)
    //   if (val) {
    //     parsed = JSON.parse(val);
    //     // console.log(parsed)
    //   }
    // }     
    // const navigate = useNavigate()
    // function goPage(page:string,keys:string[]){
    //     navigate("/country/".concat(page),{state: {keys: keys}})
    // }
    // const parsedKeys = Object.keys(parsed)
    const {parsed,parsedKeys,goPage} = useCountries()
    var newKeys: string[] = parsedKeys
    var exists = newKeys.indexOf("aud") > -1
    if(exists){
    
      newKeys = []

      Object.keys(parsed).map((key:string)=>{
       if (key.startsWith("country:")){
        var newKey =key.split("country:")[1]
        newKeys.push(newKey)
        parsed[newKey] = parsed[key]
      }
     })
    }
    

      return (
        <>
        <ResponsiveAppBar countries={newKeys}/>
        <div className="wrapper">
          {parsed ? newKeys.map((item: string)=> 
          <div className="upcont" onClick={()=> goPage(item,newKeys)}>
          <img src={parsed[item]} title={item}/>
          <div className="centered">{item}</div>      
        </div>
        ): null}
          
          </div>
          </>
      )
  }

  export default GalleryHome;

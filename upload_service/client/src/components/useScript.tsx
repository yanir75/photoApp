import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const useScript = (url: string,defer:boolean=false) => {
  useEffect(() => {
    const script = document.createElement('script');

    script.src = url;
    script.defer = defer
    // script.async = true;

    document.body.appendChild(script);

    return () => {
      document.body.removeChild(script);
    }
  }, [url]);
};

function useCountries(){
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
  return {parsed,parsedKeys,goPage}
}

export { useScript,useCountries};
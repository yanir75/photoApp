import "./App.css";


function App() {
  const doc = document.getElementById("content");
  let parsed = null;
  if (doc) {
    let val = doc.getAttribute("vals");
    if (val) {
      parsed = JSON.parse(val);
      console.log(parsed)
    }
  }
  let folders = parsed["folders"]

  return (
    <>
      {folders ? folders.map((item: {"Prefix" : string})=> <h1>{item["Prefix"].substring(0,item["Prefix"].length-1)}</h1>): null}
    </>
  );
}

export default App;

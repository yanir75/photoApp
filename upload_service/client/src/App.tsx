import { Route, Routes } from "react-router-dom";
import "./App.css";
import Gallery from "./components/gallery";
import GalleryHome from "./components/galleryHome";
import User from "./components/user";
// import lightGallery from "lightgallery";

function App() {
  return (
    <>
       <Routes>
          <Route path="/gallery" element={<GalleryHome />} />
          <Route path="/country/:country" element={<Gallery />} />
          <Route path="/user" element={<User/>} />
       </Routes>
    </>
 );

}



export default App;

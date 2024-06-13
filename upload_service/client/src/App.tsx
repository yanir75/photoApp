import { Route, Routes } from "react-router-dom";
import Gallery from "./components/gallery";
import GalleryHome from "./components/galleryHome";
import User from "./components/user";
import Upload from "./components/upload";
import "./App.css"
// import lightGallery from "lightgallery";

function App() {
  return (
    <>
       <Routes>
          <Route path="/gallery" element={<GalleryHome />} />
          <Route path="/country/:country" element={<Gallery />} />
          <Route path="/user" element={<User/>} />
          <Route path="/upload" element={<Upload/>} />
       </Routes>
    </>
 );

}



export default App;

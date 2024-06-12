import './gallery.css'
import { useLocation, useParams } from "react-router-dom";
import ResponsiveAppBar from './menu';
import LightGallery from 'lightgallery/react';
import lgVideo from 'lightgallery/plugins/video';

// import styles
import 'lightgallery/css/lightgallery.css';
import 'lightgallery/css/lg-zoom.css';
import 'lightgallery/css/lg-thumbnail.css';
import 'lightgallery/css/lg-video.css';
// If you want you can use SCSS instead of css
import lgThumbnail from 'lightgallery/plugins/thumbnail';

// import plugins if you need
import lgZoom from 'lightgallery/plugins/zoom';
import  { useCallback, useEffect, useRef, useState } from 'react';


function Gallery() {
    const [images, setImages] = useState([]); // Array instead of object

async function initPhotos(country:string) {
    await fetch('/api/'.concat(country))
        .then(response => response.json())
        .then(data => {
            var parsed = JSON.parse(data);
            setImages(parsed[country]);
        })
}

const { country } = useParams();
// let ok = false
const {state} = useLocation();
const { keys } = state; // Read values passed on state
if(country){
    useEffect(()=>{initPhotos(country)},[])
    
}

  const lightGallery = useRef<any>(null);

  const items = images.map(({Description,Type,Name,ThumbnailUrl},index)=> {
    if(Type === "image"){
        return (    {
            id: index,
            size: '1400-933',
            src: Name,
            thumb:
              Name,
            subHtml: `<div class="lightGallery-captions">
                      <h4>`+Description+`</h4>
                  </div>`,
          })
    }
    else{
        return(    {
            video: {
              source: [
                {
                  src: Name,
                  type: 'video/mp4',
                },
              ],
              attributes: { preload: false, controls: true },
            },
            thumb: ThumbnailUrl,
            //   'https://www.lightgalleryjs.com//images/demo/html5-video-poster.jpg',
            subHtml: `<div class="lightGallery-captions">
                      <h4>`+Description+`</h4>
                  </div>`,
          } as any)
    }
  })
  const openGallery = useCallback(() => {
    lightGallery.current.openGallery();
  }, []);

  const onInit = useCallback((detail: any) => {
    if (detail) {
      lightGallery.current = detail.instance;
    }
  }, []);

  // Add new slides

  return (
    
    <div className="App">
                <ResponsiveAppBar
        countries={keys}/>
        {images.map(({ThumbnailUrl})=>  <img src={ThumbnailUrl} title={ThumbnailUrl} onClick={openGallery}/>) }
      {/* <button onClick={openGallery}>Open Gallery</button> */}
      <LightGallery
        elementClassNames="custom-classname"
        dynamic
        speed={500}
        animateThumb
        dynamicEl={items}
        onInit={onInit}
        plugins={[lgZoom,lgThumbnail, lgVideo]}
      ></LightGallery>
    </div>
  );
}

export default Gallery;
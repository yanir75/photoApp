const dropContainer = document.getElementById("dropcontainer")
const fileInput = document.getElementById("images")

dropContainer?.addEventListener("dragover", (e) => {
  // prevent default to allow drop
  e.preventDefault()
}, false)

dropContainer?.addEventListener("dragenter", () => {
  dropContainer.classList.add("drag-active")
})

dropContainer?.addEventListener("dragleave", () => {
  dropContainer.classList.remove("drag-active")
})

dropContainer?.addEventListener("drop", (e) => {
  e.preventDefault()
  dropContainer.classList.remove("drag-active")
  fileInput.files = e.dataTransfer.files
})

const urlParams = new URLSearchParams(window.location.search);
myParam = urlParams.get('msg');
if (myParam !== null){
   alert(myParam.replace(/\\n/g,"\n"));
}

if (document.readyState !== 'loading') {
    // console.log('document is already ready, just execute code here');
    myInitCode();
  } else {
    document.addEventListener('DOMContentLoaded', function () {
        // console.log('document was not ready, place code here');
        myInitCode();
    });
  }
  

  function myInitCode() {
    var doc = document.getElementById('upload_form')
    doc.addEventListener('submit', async function(event) {
        event.preventDefault();
        var files = document.getElementById('images').files;
        var fileName = document.getElementsByName('fileName')[0].value
        var description = document.getElementsByName('description')[0].value
        var country = document.getElementsByName('country')[0].value

        for (var i = 0; i < files.length; i++) {
            await uploadFile(files[i],fileName+i.toString(),description,country);
        }
        
    });
    
  }
  
  function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
  

async function uploadFile(file,fileName,description,country) {
    var xhr = new XMLHttpRequest();
    xhr.open('POST', 'upload', true);

    xhr.upload.onprogress = function(event) {
        if (event.lengthComputable) {
            var percentComplete = (event.loaded / event.total) * 100;
            document.getElementById('progress').innerText = 'File ' + file.name + ': ' + percentComplete.toFixed(2) + '% uploaded';
        }
    };

    xhr.onload = function() {
        if (xhr.status === 200) {
            // Upload successful
            console.log('File uploaded: ' + file.name);
        } else {
            // Error handling
            console.error('Error uploading file: ' + file.name);
        }
    };

    xhr.onerror = function() {
        console.error('Error uploading file: ' + file.name);
    };

    var formData = new FormData();
    formData.append('fileName',fileName)
    formData.append('description',description)
    formData.append('country',country)
    formData.append('file', file);
    xhr.send(formData);
    perc = "0"
    filearr = file.name.split(".")
    fileending = filearr[filearr.length -1 ]
    while (perc != "100") {
        fetch("/api/percentage/".concat(fileName).concat('.').concat(fileending))
        .then((response) => response.text())
        .then((text) => {
            perc = text
            document.getElementById('progressUpload').innerText = 'File ' + fileName +'.'+ fileending + ': ' + perc + '% uploaded to S3';
        }
    );
    await sleep(5000)
    }
    
}

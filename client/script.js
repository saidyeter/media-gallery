const contentBox = document.querySelector("div.content");

function getBaseUrl() {
  return window.location.origin + window.location.pathname;
}

function params() {
  const items = new Proxy(new URLSearchParams(window.location.search), {
    get: (searchParams, prop) => searchParams.get(prop),
  });
  return items;
}

async function fillContent(address) {
  const res = await fetch(address);
  const data = await res.json();
  //  console.log(data);
  if (data.Contents && data.Contents.length > 0) {
    data.Contents.forEach((content) => { 
      createContent(content);
    });

    if (data.Next) {
      createButton(data.Next);
    }

    return data.Contents.length;
  }
  return 0;
}

(async function () {
  let address = "/content";

  const folder = params().directory;

  if (folder) {
    address += "/" + folder;
  }

  const inserted = await fillContent(address);
  if (inserted == 0) {
    fillContent("/content");
  }
})();

function handleNext(folder, button) {
  let address = "/content";
  address += "/" + folder;

  fillContent(address);
  button.parentNode.removeChild(button);
}

function createButton(url) {
  var button = document.createElement("button");
  button.type = "button";
  button.innerText = "more";
  button.onclick = () => {
    handleNext(url, button);
  };
  button.classList.add("grid-item");
  contentBox.appendChild(button); // add the button to the context
}

function createContent({ Name, ThumbPath, ActualPath, IsDir }) {
  // console.log("createContent", { Name, ThumbPath, ActualPath, IsDir });
  const content = document.createElement("img");
  content.onerror = () => {
    content.src = "image-not-found.png";
  } 
   
  content.setAttribute("alt", Name);
  if (IsDir) {
    content.setAttribute("src", 'folder.png');  
  } else {
    content.setAttribute("src", "/file/" + ThumbPath);
    content.setAttribute("target", "_blank"); 
  }
  const contentWrapper = document.createElement('div')  
  contentWrapper.classList.add('content-wrapper')

  content.classList.add("photo");
  content.classList.add("grid-item");
  let href;
  if (IsDir) {
    href = getBaseUrl() + "?directory=" + ActualPath; 
  } else {
    href = "/file/" + ActualPath; 
  }

  contentWrapper.addEventListener("click", () => (window.location = href));
  contentWrapper.style.cursor = "pointer";
  contentWrapper.appendChild(content)

  const contentName = document.createElement("div");
  contentName.innerText = Name;
  contentName.classList.add("content-name");
  contentWrapper.appendChild(contentName);
  contentBox.appendChild(contentWrapper);
}

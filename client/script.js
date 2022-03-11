const apiAddress = "http://localhost:8080/";

const contentBox = document.querySelector("div.content");

function getBaseUrl() {
  return window.location.origin + window.location.pathname
}

function createLink({ Name, ThumbPath, ActualPath, IsDir }) {
  const link = document.createElement("a");
  link.setAttribute("alt", Name);
  if (IsDir) {
    const href = getBaseUrl() + "?directory=" + ActualPath
    link.setAttribute("href", href);
    link.style.backgroundImage = `url('folder.png')`;
  } else {
    link.setAttribute("href", apiAddress + "/file/" + ActualPath);
    link.setAttribute("target", "_blank");
    link.style.backgroundImage = `url('${apiAddress + "/file/" + ThumbPath}')`;
  }
  link.classList.add("photo");
  if (getContentState() == 'desc') {
    link.style.display = 'none'
  }
  contentBox.appendChild(link);
}

function createDesc({ Name, ThumbPath, ActualPath, IsDir }) {
  const div = document.createElement("div");
  div.innerHTML = Name
  let href
  if (IsDir) {
    href = getBaseUrl() + "?directory=" + ActualPath
    div.style.backgroundImage = `url('folder.png')`;
  } else {
    href = apiAddress + "/file/" + ActualPath
    div.style.backgroundImage = `url('${apiAddress + "/file/" + ThumbPath}')`;
  }

  div.addEventListener("click", () => window.location = href);
  div.classList.add("photo");
  div.classList.add("desc");
  if (getContentState() == 'content') {
    div.style.display = 'none'
  }
  contentBox.appendChild(div);
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
      createLink(content);
      createDesc(content);
    });

    if (data.Next) {
      createButton(data.Next);
    }

    return data.Contents.length;
  }
  return 0;
}

(async function () {
  const address = window.location.origin + window.location.pathname;
  const home = document.getElementById("home");
  home.href = address;
})();

(async function () {
  let address = apiAddress + "content";

  const folder = params().directory;

  if (folder) {
    address += "/" + folder;
  }

  const inserted = await fillContent(address);
  if (inserted == 0) {
    fillContent(apiAddress + "content");
  }
})();

function handleNext(folder, button) {


  let address = apiAddress + "content";
  address += "/" + folder;

  fillContent(address);
  button.parentNode.removeChild(button)
}
function createButton(url) {
  var button = document.createElement("button");
  button.type = "button";
  button.innerText = "more"
  button.onclick = () => {
    handleNext(url, button);
  };
  contentBox.appendChild(button); // add the button to the context
}

function getContentState() {
  const a = document.querySelector('a.photo')
  if (a) {
    return a.style.display == "none" ? 'desc' : 'content'
  }
  return  'desc'
}
function toggleDescAndLink() {
  if (getContentState() == 'content') {
    document.querySelectorAll('a.photo').forEach(link => link.style.display = 'none')
    document.querySelectorAll('div.photo').forEach(link => link.style.display = 'flex')

  }
  else {
    document.querySelectorAll('a.photo').forEach(link => link.style.display = 'block')
    document.querySelectorAll('div.photo').forEach(link => link.style.display = 'none')

  }
}


const toggle = document.querySelector("div.toggle");
toggle.addEventListener('click', toggleDescAndLink)
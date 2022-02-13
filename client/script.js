const apiAddress = "http://192.168.1.105:8080/";

const contentBox = document.querySelector("div.content");

function createLink({ Name, ThumbPath, ActualPath, IsDir }) {
  const link = document.createElement("a");
  link.setAttribute("alt", Name);
  if (IsDir) {
    link.setAttribute(
      "href",
      window.location.origin +
        window.location.pathname +
        "?directory=" +
        ActualPath
    );
    link.style.backgroundImage = `url('folder.png')`;
  } else {
    // console.log(ActualPath,ThumbPath);
    link.setAttribute("href", apiAddress + "/file/" + ActualPath);
    link.setAttribute("target", "_blank");
    link.style.backgroundImage = `url('${apiAddress + "/file/" + ThumbPath}')`;
    // console.log(link);
  }
  link.classList.add("photo");
  contentBox.appendChild(link);
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
  if (data.Files && data.Files.length > 0) {
    data.Files.forEach((content) => {
      createLink(content);
    });

    if (data.Next) {
      createButton(data.Next);
    }

    return data.Files.length;
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

function handleNext(folder,button) {


  let address = apiAddress + "content";
  address += "/" + folder;

  fillContent(address);
  button.parentNode.removeChild(button)
}
function createButton(url) {
  var button = document.createElement("button");
  button.type = "button";
  button.innerText= "more"
  button.onclick = () => {
    handleNext(url,button);
  };
  contentBox.appendChild(button); // add the button to the context
}

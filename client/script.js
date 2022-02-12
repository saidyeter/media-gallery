const apiAddress = "http://192.168.1.104:8080/";

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
    link.setAttribute("href", apiAddress + "/file/" + ActualPath);
    link.setAttribute("target", "_blank");
    link.style.backgroundImage = `url('${apiAddress + "/file/" + ThumbPath}')`;
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

(async function () {
  const address = window.location.origin + window.location.pathname;
  const home = document.getElementById("home");
  home.href = address;
})();

(async function () {
  let address = apiAddress + "content";

  let folder = params().directory;

  if (folder) {
    folder = folder.replace(new RegExp(":", "g"), "%3A");
    folder = folder.replace(new RegExp("/", "g"), "%5C");
    folder = folder.replace(new RegExp("\\\\", "g"), "%5C");
    // console.log(folder);
    address += "/" + folder;
  }
  const res = await fetch(address);
  const data = await res.json();
  // console.log(data);
  if (data.Files) {
    data.Files.forEach((content) => {
      createLink(content);
    });
  } else {
    address = apiAddress + "content";
    const res = await fetch(address);
    const data = await res.json();
    data.Files.forEach((content) => {
      createLink(content);
    });
  }
})();

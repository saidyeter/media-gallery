const contentBox = document.querySelector("div.content");

function onImgClick(e) {
  console.log(this);
}
function createLink(name, thumbPath, actualPath, isFolder) {
  /*
    <a href="https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg"
        title="Photo 1" class="photosGrid__Photo" target="_blank"
        style="background-image: url('https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg')"></a>
    */
  const link = document.createElement("a");
  link.setAttribute("href", actualPath);
  link.setAttribute("alt", name);
  link.setAttribute("target", "_blank");
  link.style.backgroundImage = `url('${thumbPath}')`;
  if (isFolder) {
      console.log(isFolder);
    link.style.backgroundImage = `url('folder.png')`;
    // link.innerText = name;
  }
  link.classList.add("photo");
  contentBox.appendChild(link);
}

const imgs = [
    {
      name: "tumblr_mrraat0H431st5lhmo1_1280.jpg",
      thumbPath:
        "https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
      actualPath:
        "https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
    },
    {
      name: "tumblr_mrraat0H431st5lhmo1_1280.jpg",
      thumbPath:
        "https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
      actualPath:
        "https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
        isFolder:true,
    },
  {
    name: "tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
    thumbPath:
      "https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
    actualPath:
      "https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  },
  {
    name: "tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
    thumbPath:
      "https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
    actualPath:
      "https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  },
  {
    name: "blake-verdoorn-gM-RfQsZK98-unsplash.jpg",
    thumbPath:
      "http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
    actualPath:
      "http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  },
  {
    name: "bnsplash.jpg",
    thumbPath:
      "https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
    actualPath:
      "https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  },
  // {
  //     name : "tumblr_mrraat0H431st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "blake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  //     thumbPath:"http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  //     actualPath: "http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  // },
  // {
  //     name : "bnsplash.jpg",
  //     thumbPath:"https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  //     actualPath: "https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  // },
  // {
  //     name : "tumblr_mrraat0H431st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "blake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  //     thumbPath:"http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  //     actualPath: "http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  // },
  // {
  //     name : "bnsplash.jpg",
  //     thumbPath:"https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  //     actualPath: "https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  // },
  // {
  //     name : "tumblr_mrraat0H431st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "blake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  //     thumbPath:"http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  //     actualPath: "http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  // },
  // {
  //     name : "bnsplash.jpg",
  //     thumbPath:"https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  //     actualPath: "https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  // },
  // {
  //     name : "tumblr_mrraat0H431st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "blake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  //     thumbPath:"http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  //     actualPath: "http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  // },
  // {
  //     name : "bnsplash.jpg",
  //     thumbPath:"https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  //     actualPath: "https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  // },
  // {
  //     name : "tumblr_mrraat0H431st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  //     thumbPath:"https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  //     actualPath: "https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg",
  // },
  // {
  //     name : "blake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  //     thumbPath:"http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  //     actualPath: "http://192.168.1.104:8080/file/C%3A%5C%5CUsers%5C%5Csaid.yeter%5C%5CDesktop%5C%5Cfotolar%5C%5Cblake-verdoorn-gM-RfQsZK98-unsplash.jpg",
  // },
  // {
  //     name : "bnsplash.jpg",
  //     thumbPath:"https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  //     actualPath: "https://images.unsplash.com/photo-1519120944692-1a8d8cfc107f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2hpdGUlMjB3YWxscGFwZXJ8ZW58MHx8MHx8&w=1000&q=80",
  // }
];

function getContent() {
  // fetch from the server
  imgs.forEach((img) => {
    createLink(img.name, img.thumbPath, img.actualPath,img.isFolder);
  });
}

getContent();

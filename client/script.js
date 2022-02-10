
const sideBar = document.querySelector('div.sidebar')
const contentBox = document.querySelector('div.content')

function onImgClick(e) {
    console.log(this);
}
function createImgPreview(name , path, thumb) {
    const imgPreview = document.createElement("img-preview")
    imgPreview.setAttribute("thumb", thumb); 
    imgPreview.setAttribute("name", name);  
    imgPreview.setAttribute("path", path);  
    imgPreview.addEventListener('imgClick', onImgClick)
    contentBox.appendChild(imgPreview)
}



createImgPreview("by Seval","path1","https://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg")
createImgPreview("by Said","path2","https://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg")
createImgPreview("by John","path3","https://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg")


function createSidebarItem(name) {
    const item = document.createElement("sidebar-item")
    item.innerHTML = name
    item.addEventListener('click', onImgClick)
    sideBar.appendChild(item)
}

createSidebarItem('üni')
createSidebarItem('ankara gezisi')
createSidebarItem('ayşe kına')
createSidebarItem('ahmet düğün')
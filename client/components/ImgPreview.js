const imgPreviewTemplate = document.createElement("template");
imgPreviewTemplate.innerHTML = `
<style> 
  div > img {
    vertical-align: middle;
    max-width : 20rem;
    max-height: 20rem;
  }

  @media screen and (max-width: 30rem) {
    div {
      width : 100%;
      height: 100%;
    }
  }

  div > aside {
     transform: translate(0, -2rem);
    /**/
  }

  div > aside > span {
    color: white;
    background-color: rgba(0, 0, 0, .5);
  }



</style>
<div>
  <img />
  <aside>
    <span />
  </aside>
</div>
`;
/*
  div > img {
    object-fit: contain;
    max-height: calc(100vh - 1rem);
    max-width : 20rem;
  }

  @media screen and (max-width: 30rem) {
    div {
      max-width: calc(100vw - 1rem);
      max-height: 20rem;
    }
  }

 */
class ImgPreview extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.appendChild(imgPreviewTemplate.content.cloneNode(true));

    //"data:image/png;base64,"
    setTimeout(() => {
      this.shadowRoot.querySelector("img").src = this.getAttribute("thumb");
      this.shadowRoot.querySelector("img").alt = this.getAttribute("path");
      this.shadowRoot.querySelector("span").innerHTML =
        this.getAttribute("name");
      this.shadowRoot.querySelector("img").addEventListener("click", () => {
        this.onImgClick();
      });

      this.imgClickEvent = new Event("imgClick");
    }, 100);
  }
  onImgClick() {
    this.dispatchEvent(this.imgClickEvent);
    // console.log("dispathced", this.imgClickEvent, this.parentElement);
  }

  // connectedCallback() {}

  // disconnectedCallback() {}
}

window.customElements.define("img-preview", ImgPreview);

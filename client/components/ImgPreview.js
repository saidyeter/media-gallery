const imgPreviewTemplate = document.createElement("template");
imgPreviewTemplate.innerHTML = `
<style>
div{

}
div > img {
  max-width : 20rem;
  max-height: 20rem;
}
div > aside {
  /* transform: translateY() */
  transform: translate(1rem, -2rem);
  /* text-align: right; */
}
</style>
<div>
  <img />
  <aside>
    <span/>
  </aside>
</div>
`;

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

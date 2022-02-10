const sidebarTemplate = document.createElement("template");
sidebarTemplate.innerHTML = `
<style>
div{

}
</style>
<div>
    <slot />
</div>
`;

class SidebarItem extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.appendChild(sidebarTemplate.content.cloneNode(true)); 
  }
}

window.customElements.define("sidebar-item", SidebarItem);

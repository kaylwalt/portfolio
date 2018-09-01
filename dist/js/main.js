// Select DOM Items
(() => {
  const menuBtn = document.querySelector(".menu-btn");
  const menu = document.querySelector(".menu");
  const menuNav = document.querySelector(".menu-nav");
  const menuBranding = document.querySelector(".menu-branding");
  const navItems = document.querySelectorAll(".nav-item");
  const webDev = document.querySelector("#web-dev");
  const dataAnalyst = document.querySelector("#data-analyst");
  const softwareEngineer = document.querySelector("#software-engineer");
  const containers = document.querySelectorAll(".container");
  const portrait = document.querySelector(".portrait");
  const resumeDownloadButton = document.querySelector("#resumeDownload");
  const menu_links = document.querySelectorAll(".nav-item");
  const contactInfo = document.querySelector("#contact");

  menu_links[0].addEventListener("click", () => {
    portrait.classList.add("show");
    resumeDownloadButton.classList.remove("show");
    contactInfo.classList.remove("show");
  });
  menu_links[1].addEventListener("click", () => {
    portrait.classList.remove("show");
    resumeDownloadButton.classList.add("show");
    contactInfo.classList.remove("show");
  });
  menu_links[2].addEventListener("click", () => {
    portrait.classList.remove("show");
    resumeDownloadButton.classList.remove("show");
    contactInfo.classList.add("show");
  });
  let cartIndex = 0;
  let portfolioIndex = 1;
  let researchIndex = 2;
  let autonomousIndex = 3;

  webDev.addEventListener("mouseenter", e => {
    containers[cartIndex].classList.add("highlighted");
    containers[portfolioIndex].classList.add("highlighted");
  });
  webDev.addEventListener("mouseleave", e => {
    containers[cartIndex].classList.remove("highlighted");
    containers[portfolioIndex].classList.remove("highlighted");
  });
  dataAnalyst.addEventListener("mouseenter", e => {
    containers[researchIndex].classList.add("highlighted");
  });
  dataAnalyst.addEventListener("mouseleave", e => {
    containers[researchIndex].classList.remove("highlighted");
  });
  softwareEngineer.addEventListener("mouseenter", e => {
    containers[autonomousIndex].classList.add("highlighted");
  });
  softwareEngineer.addEventListener("mouseleave", e => {
    containers[autonomousIndex].classList.remove("highlighted");
  });

  var resized = false;

  window.onresize = () => {
    if (!resized) {
      resized = true;
      const images = document.querySelectorAll("img");
      images.forEach(image => (image.style.height = "auto"));
      images.forEach(image => (image.style.width = "100%"));
    }
  };
  equalizeImages = () => {
    const images = document.querySelectorAll("img");
    const maxHeight = Math.max(...Array.from(images).map(i => i.height));
    const maxWidth = Math.max(...Array.from(images).map(i => i.width));
    images.forEach(image => (image.style.height = maxHeight.toString() + "px"));
    images.forEach(image => (image.style.width = maxWidth.toString() + "px"));
  };
  equalizeImages();
  //Set Initial State Of Menu
  let showMenu = false;

  menuBtn.addEventListener("click", toggleMenu);

  function toggleMenu() {
    if (!showMenu) {
      menuBtn.classList.add("close");
      menu.classList.add("show");
      menuNav.classList.add("show");
      menuBranding.classList.add("show");
      navItems.forEach(item => item.classList.add("show"));

      // Set Menu State
      showMenu = true;
    } else {
      menuBtn.classList.remove("close");
      menu.classList.remove("show");
      menuNav.classList.remove("show");
      menuBranding.classList.remove("show");
      navItems.forEach(item => item.classList.remove("show"));

      // Set Menu State
      showMenu = false;
    }
  }
})();

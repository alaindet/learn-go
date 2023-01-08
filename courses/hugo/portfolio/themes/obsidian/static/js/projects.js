document.addEventListener('DOMContentLoaded', () => {
  resizePlaceholderImages();
});

function resizePlaceholderImages() {
  const w = window.innerWidth;
  for (const img of document.querySelectorAll('img')) {
    if (!img.src.includes("placeholder")) return;
    img.src = `http://via.placeholder.com/${w}x150`;
  }
}

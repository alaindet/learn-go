document.addEventListener('DOMContentLoaded', () => {
  parseContentAsToHTML(
    document.querySelector('.snippet__content')
  );
});

function parseContentAsToHTML(element) {
  element.innerHTML = element.innerHTML
    .replace(/\\r\\n/g, '<br>')
    .replace(/\\n/g, '<br>');
}

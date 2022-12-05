for (const navLink of document.querySelectorAll('nav a')) {
	if (navLink.getAttribute('href') == window.location.pathname) {
		link.classList.add('live');
		break;
	}
}

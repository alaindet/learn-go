document.addEventListener('DOMContentLoaded', init);
window.onbeforeunload = exit;

const APP = {
  cleanup: [],
};

function init() {
  console.log('Welcome to Snippetbox v1.0.0');
  initFlashMessages();
}

function exit() {
  console.log('Leaving Snippetbox v1.0.0');
  APP.cleanup.forEach(fn => fn());
}

function initFlashMessages() {
  APP.cleanup.push(
    addBoundaryEventListener(
      document,
      'click',
      '.flash__dismiss',
      onFlashDismissClick,
    )
  );
}

function onFlashDismissClick(element) {
  element.parentElement.remove();
}

/**
 *
 * TODO: Move to global
 *
 * This adds an event listener on a "boundary" element so that it listens
 * to all events of a given type matching a given CSS selector inside the boundary
 *
 * It is useful for listening to clicks of dynamic elements inside a "boundary"
 * parent element which always exists in the DOM
 *
 * Ex.:
 * addBoundaryEventListener(document, 'click', '.flash__dismiss', target => {
 *   console.log('Clicked on flash dismiss');
 * });
 *
 * @param Element boundaryElement
 * @param string eventType (Ex.: 'click')
 * @param string selector
 * @param Function handler
 * @returns
 */
function addBoundaryEventListener(boundaryElement, eventType, selector, handler) {

  const listener = event => {
    if (!event.target.matches(selector)) return;
    handler(event.target);
  };

  boundaryElement.addEventListener(eventType, listener);
  return () => boundaryElement.removeEventListener(eventType, listener);
}

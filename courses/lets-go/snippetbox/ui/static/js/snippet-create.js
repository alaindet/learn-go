// TODO: This should be handled by the server
document.addEventListener('DOMContentLoaded', () => {
  addErrorModifierClassToInvalidFormControls();
});

function addErrorModifierClassToInvalidFormControls() {
  document.querySelectorAll('.form-control').forEach(formControl => {
    if (!!formControl.querySelector('.error-message')) {
      formControl.classList.add('--error');
    }
  });
}

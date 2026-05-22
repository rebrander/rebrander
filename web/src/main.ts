type RebrandForm = HTMLFormElement & {
  elements: HTMLFormControlsCollection & {
    source: HTMLInputElement;
    target: HTMLInputElement;
  };
};

const form = document.querySelector<RebrandForm>("#rebrand-form");
const preview = document.querySelector<HTMLPreElement>("#preview");

form?.addEventListener("submit", (event) => {
  event.preventDefault();
  const source = form.elements.source.value.trim();
  const target = form.elements.target.value.trim();
  const command = `rebrander ${source} ${target} --plan-only`;
  if (preview) preview.textContent = command;
});


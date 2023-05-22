// <button id="fontPlus">Font +</button>
// <button id="fontMinus">Font -</button>
const fontPlus = document.getElementById("fontPlus");
const fontMinus = document.getElementById("fontMinus");
const body = document.body;

function bodyFonstSize() {
    const size = getComputedStyle(body).fontSize;
    return parseFloat(size.substring(0, size.length - 2));
}


fontPlus.addEventListener("click", () => {
    const size = `${bodyFonstSize() + 1}px`
    body.style.fontSize = size
    fontPlus.innerText = `font ${size} +`
});


fontMinus.addEventListener("click", () => {
    const size = `${bodyFonstSize() - 1}px`
    body.style.fontSize = size
    fontPlus.innerText = `font ${size} -`
});

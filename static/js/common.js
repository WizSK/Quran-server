// {{ define "theme-js" }}
const Classes = document.body.classList;
const themeTgl = document.getElementById("theme-tgl")
Classes.add("dark");
themeTgl.addEventListener("click", () => {
    if (Classes.contains("dark")) {
        Classes.remove("dark");
        Classes.add("light");
        themeTgl.innerText = "Go Dark"
    } else {
        Classes.add("dark");
        Classes.remove("light");
        themeTgl.innerText = "Go Light"
    }
})
// {{ end }}


// {{ define "font-size-js" }}

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
    fontPlus.innerText = `Font ${size} +`
});


fontMinus.addEventListener("click", () => {
    const size = `${bodyFonstSize() - 1}px`
    body.style.fontSize = size
    fontMinus.innerText = `Font ${size} -`
});

// {{ end }}

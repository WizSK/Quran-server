// {{ define "theme-js" }}

const body = document.body;
const theme_btn = document.getElementById("theme-tgl");

// On page load or when changing themes, best to add inline in to avoid FOUC
if (
    localStorage.theme === "dark" ||
    (!("theme" in localStorage) &&
        window.matchMedia("(prefers-color-scheme: dark)").matches)
) {
    body.classList.remove("light");
    body.classList.add("dark");
    theme_btn.innerText = "dark";
} else {
    body.classList.remove("dark");
    body.classList.add("light");
    theme_btn.innerText = "light";
}

function themeTgl() {
    if (body.classList.contains("dark")) {
        body.classList.remove("dark");
        body.classList.add("light");
        theme_btn.innerText = "light";
        localStorage.theme = "light";
    } else {
        body.classList.remove("light");
        body.classList.add("dark");
        theme_btn.innerText = "dark";
        localStorage.theme = "dark";
    }
}


theme_btn.addEventListener("click", themeTgl)


// Returns true if its on any index page; js to tai agla lektasi :)
function onIndexPage() {
    const wl = window.location.pathname;
    switch (wl) {
        case "/", "/t", "/t/", "/w", "/w/":
            return true;
        default:
            return false;
    }


}


const scrollPercentage = 0.8; // Scroll 60% of the screen height
const windowHeight = window.innerHeight;
const scrollDuration = 400; // Animation duration in milliseconds

// Easing function to create a smooth start and end with a sudden stop
function easeInOutCubic(t) {
    return t < 0.5
        ? 4 * t * t * t
        : 1 - Math.pow(-2 * t + 2, 3) / 2;
}

function smoothScroll(targetY) {
    const startY = window.scrollY;
    const distance = targetY - startY;
    const startTime = performance.now();

    function scrollStep(timestamp) {
        const currentTime = timestamp - startTime;
        const timeFraction = Math.min(currentTime / scrollDuration, 1);
        const animationProgress = easeInOutCubic(timeFraction);
        const scrollPosition = startY + distance * animationProgress;
        window.scrollTo(0, scrollPosition);

        if (currentTime < scrollDuration) {
            requestAnimationFrame(scrollStep);
        }
    }

    requestAnimationFrame(scrollStep);
}


// for keyboard shortcuts
addEventListener("keydown", (e) => {
    console.log("keyDown:", e.key.toLowerCase());

    switch (e.key.toLowerCase()) {
        case "k":
            smoothScroll(window.scrollY - windowHeight * scrollPercentage);
            break;

        case "j":
            smoothScroll(window.scrollY + windowHeight * scrollPercentage);
            break;

        case "t":
            themeTgl();
            break;

        // this shoud not run on index pages
        case "+", "=":
            if (onIndexPage()) return;
            font.increase();
            break;
        case "-":
            if (onIndexPage()) return;
            font.decrease();
            break;
        case "0":
            if (onIndexPage()) return;
            font.reset();
            break;

    }
})

// {{ end }}

// {{ define "font-size-js" }}

if (localStorage.fontSize) {
    body.style.fontSize = localStorage.fontSize;
}

const fontPlus = document.getElementById("fontPlus");
const fontMinus = document.getElementById("fontMinus");
// const body = document.body;

function bodyFonstSize() {
    const size = getComputedStyle(body).fontSize;
    return parseFloat(size.substring(0, size.length - 2));
}

// some methods for the font Minupulation
const font = {
    reset: () => {
        body.style.fontSize = "";
        localStorage.fontSize = "";
    },

    increase: () => {
        const size = `${bodyFonstSize() + 1}px`
        localStorage.fontSize = size;
        body.style.fontSize = size
        fontPlus.innerText = `Font ${size} +`
    },

    decrease: () => {
        const size = `${bodyFonstSize() - 1}px`
        localStorage.fontSize = size;
        body.style.fontSize = size
        fontPlus.innerText = `Font ${size} +`
    }
}

fontPlus.addEventListener("click", font.increase);
fontMinus.addEventListener("click", font.decrease);


// {{ end }}

# Quran is the central religious text of Islam

This is a Small server written entirely in go. And it's fully server side rendered. Fast and Works fine.

## Pros
- Fast, no Client side JavaScript
- It' made with [Gin Web Framework](https://gin-gonic.com/)
- All fonts and files are served form this server
- Rendered using go templates

## Todo
- [x] Cash the main page with maps
- [x] Jump to any line with "http://localhost:8000/2#255" to go-to line 255 of chapter 2
- [ ] Add functionality to add any translations.
- [ ] Work on styling
    - [ ] Day and night mode

## To run Run
```bash
git clone https://github.com/WizSK/Quran-server.git
cd Quran-server
go run . # It will run at port 8000 if no argument is provided
# go run . 8888  to specify the port
```

## Read Quran form [Quran.com](https://quran.com/)

# Quran is the central religious text of Islam

This is a Small server written entirely in go. No 3rd party library is used. And it's fully server side rendered. Fast and Works fine.

## Pros
- Fast, no Client side JavaScript. A little bit for light/dark mode.
- All fonts and files are served form this server
- Rendered using go templates

## Todo
- [x] Docker image.
- [x] Cache
- [x] Jump to any line with "http://localhost:8000/2#255" to go-to line 255 of chapter 2
- [x] Add functionality to add any translations. (with array)
- [x] Add Word by word. "/w/" path.
- [x] Add Word by word with translation. "/t/" path.
- [ ] Work on styling
    - [ ] Some more styling needed.
    - [x] Day and night mode
    - [x] Font resize.
- [ ] Add documentation to add translation.

## Run
```bash
git clone https://github.com/WizSK/Quran-server.git
cd Quran-server
go run . # It will run at port 8001 if no argument is provided
# go run . 8888  to specify the port
```

## Docker
[https://hub.docker.com/r/wizsk/quran-server](https://hub.docker.com/r/wizsk/quran-server)

```
docker pull wizsk/quran-server:0.1
docker run -p 8001:8001 wizsk/quran-server:0.1
```

## Read Quran form [Quran.com](https://quran.com/)

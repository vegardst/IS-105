# Oppgave 2

Se youtube video
* https://www.youtube.com/watch?v=ELZRtzlJTLA

###Bygge til server:
* GOOS=linux GOARCH=amd64 go build -v hello.go
* Sende til server: pscp  ~/OneDrive/Skole/GoLang hello root@wiklem.sexy:
* Endre rettigheter på fil: chmod 744 hello
* Kjøre filen ./hello

Nyttige linker:
* https://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5
* https://golang.org/doc/install/source#environment
# ICA 06

## Eksperiment 1
/// Figur 1

Her i figure 1. ser vi en frekvensspektrum på setningen “gruppe 1” som ble sagt av en kvinne(venstre) og en mann (høyre).

Den øverste halvdelen ser vi at begge har delt uttalelsen av setningen i tre deler og det er en viss likhet mellom de, både den første og andre delen avsluttes ganske skarpt mens den siste daler ned og avslutter med en liten top. forskjellen ligger mest i kraften på uttalelsen enn formen på ordene(dialekt).

I den nederste delen (spektrogram) ser vi litt mer forskjell fra person til person,
gul viser til intensiteten til lyden, mens blå viser til hvor lyst eller mørk tonen er.


/// Figur 2
Dette bilde viser en annen mann som også sier samme ordet, litt annerledes i uttalelsen enn i figure 1. men fortsatt ser vi en likhet mellom figurene. I den øverste halvdelen er de to første delene mer bredere, mens siste delen er ganske lik til figure 1.

Den nederste halvdelen ser vi mer oppbyggingen av setningen.

##### Spektrogram - grå/sorte bakgrunnen
spektrogram er et visuelt representasjon av lyd




##### Tonehøyde - blå linje
Tonehøyde er hvor mørk eller lys en stemme er, menn har som regel lavere tonehøyde enn kvinner, noe som ser ut til å stemme med figure 1, der den blå linjen ligger høyere i Hz på
bildet til venstre, og lavere i Hz i bildet til høyre.


##### Formant  - røde prikker
Formanter viser hvor vi legger mest trykk på ordet(tonen) en tone blir som regel delt opp i to deler den laveste er grunntone (lavest Hz) og resten er overtoner. Det er i overtonene at vi skiller mellom person til person, et problem vil være skille mellom dialekter, og uttalelse av ord.
 Vi ser en likhet i grunntonen på bildene (nederste røde prikker) og mer variert i resten av de røde prikkene.

##### Intensitet - gul linje
Intensitet er hvor høyt ordene blir uttalt.

##### Hvordan man kan gjenkjenne språk
Vi kan bruke formanter til å gjenkjenne vokaler og hvor personen legger trykk på ordene og sammenligne det med en database for så se om det passer til et innlagt språket i databasen.


## Eksperiment 2
Start server.go, gå til localhost:8080 og skriv inn tekst som bli konvertert til tale ved hjelp av javascript funksjon som sender teksten til en docker client som kjører https://github.com/parente/espeakbox

## Eksperiment 3
Her brukes Google Speech API, med eksempler fra https://cloud.google.com/speech/docs / for å gjøre om tale til tekst.
Her finnes det både streaming og ikke streaming eksempler til forskjellige programeringsspråk.

## Ekstra

Laget en Android applikasjon som bruker go koden ved hjelp av GoMobile sin gobind or presanterer eksperiment 2 og 3 samt ICA05 i en mobil applikasjon:

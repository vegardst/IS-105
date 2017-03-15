# ICA 04
#### Oppgave 1:
a) Når man gjør utskrift av byteslicen fra tekstfilen, ser man at linjeskift i text1 er gjort med Carriage Return "\r", i tilegg til "\n" LineFeed som også brukes for linjeskifte.<br>
Dette er grunnen til at den første filen er litt større en den andre, som kun bruker Line Feed.<br>
Denne fenomenet kommer av at forskjellige OS tolkninger av "ny linje".<br>
Hvis vi itererer gjennom byteslicen vil vi se:<br>
```
D '\r' 1101 
A '\n' 1010 
```
som vi finner igjen i ASCII kode tabellen:

```
10	012	0A	00001010	LF	&#010;	 	Line Feed 
13	015	0D	00001101	CR	&#013;	 	Carriage Return
```

Grunnen til dette, er som beskrevet i linken at forskjellige operativsystem har brukt forskjellig kode for å indikere linjeskifte.
<br>Kilde: http://stackoverflow.com/questions/1761051/difference-between-n-and-r

b) Her har jeg laget tre metoder i filehandling.go som sjekker om text1 og text2 har linjeskift, og gir deg muligheten til å skrive inn egen filbane ved hjelp av scanner.

#### Oppgave 2
a) Her "os" pakken og koden ligger i fileinfo.go<br>
Se veldegsmappen for eksempler. Jeg slet dog med å velge snarveier som input.
<br>SELFNOTE: HUSK Å LEGG TIL: Unix block device<br>
<br>
b) 
<br>
<br>
c)
#### Oppgave 3
a) Ved bruk av os og ioutil pakkene kan man behandle filer.
I disse pakken finner vi metoder som kan gjøre det meste med filer. Eksempler:<br>
```
os.Create
os.Remove
os.Rename
os.OpenFile
os.Open
os.Chmod
```
b)

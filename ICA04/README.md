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
<br>

Resultat oppgave 1a:
```
Oppgave 1a:
Utskrift %%X, base15:
ByteSlice Text1: 546573746572206C696E6A65736B6966742E0D0A4F6720656E2074696C202E2E2E0D0A4F6720656E2074696C202E2E2E0D0A4F6720656E2074696C202E2E2E0D0A
ByteSlice Text2: 546573746572206C696E6A65736B6966742E0D0A4F6720656E2074696C202E2E2E0D0A4F6720656E2074696C202E2E2E0D0A4F6720656E2074696C202E2E2E0D0A
Utskrivt %%q , characters
ByteSlice Text1: "Tester linjeskift.\r\nOg en til ...\r\nOg en til ...\r\nOg en til ...\r\n"
ByteSlice Text2: "Tester linjeskift.\r\nOg en til ...\r\nOg en til ...\r\nOg en til ...\r\n"
```
Grunnen til dette, er som beskrevet i linken at forskjellige operativsystem har brukt forskjellig kode for å indikere linjeskifte.
<br>Kilde: http://stackoverflow.com/questions/1761051/difference-between-n-and-r

b) Her har jeg laget tre metoder til oppgaven som sjekker om text1 og text2 har linjeskift, og gir deg muligheten til å skrive inn egen filbane ved hjelp av scanner.
<br> Resultat oppgave 1b:

```
Oppgave 1b:
Text 1:  Teksten inneholder LineBreak og CarriageReturn
Text 2:  Teksten inneholder LineBreak og CarriageReturn
Skriv filbane på filen du vil scanne: (Root soruce definert i var)
files/text1.txt
Din fil: files/text1.txt
Scanning :  Teksten inneholder LineBreak og CarriageReturn
```

#### Oppgave 2
a) Skriv inn filen du ønsker å scanne.
```
Oppgave 2a:
Skriv filbane på filen du vil scanne: (Root soruce definert i var)
files/pg100.txt
Informasjon om filen: files/pg100.txt
Bytes:  5.589889e+06
Kibibytes:  5589.889
Mibibytes:  5.589889
Mibibytes:  0.005589889
Dette er en fil
Filrettigheter:  -rw-rw-rw-
Filen er ikke 'append only'
Filen er ikke 'name pipe'
Filen er ikke  'device oppgaver'
Filen er ikke  ' Unix character device'
Filen er ikke  'symbolic link'
```

b) På linux instansen:
```
Oppgave 2b:
Oppgaven må gjøres på linux
/dev/stdin​ :
Informasjon om filen: /dev/stdin
Bytes:  0
Kibibytes:  0
Mibibytes:  0
Mibibytes:  0
Filrettigheter:  Dcrw--w----
Filen er ikke 'append only'
Filen er ikke 'name pipe'
Filen er 'device oppgaver'
Filen er ' Unix character device'
Filen er ikke  'symbolic link'
/dev/ram0 :
Informasjon om filen: /dev/ram0
Bytes:  0
Kibibytes:  0
Mibibytes:  0
Mibibytes:  0
Filrettigheter:  Drw-rw----
Filen er ikke 'append only'
Filen er ikke 'name pipe'
Filen er 'device oppgaver'
Filen er ikke  ' Unix character device'
Filen er ikke  'symbolic link'
```

<br>
c) Filene er bygd for hvert system og ligger under src\oppgaver\fileinfo\builds
Test av pg100.txt på windows:

```
Skriv filbane på filen du vil scanne:
C:\OneDrive\Skole\ICA\ICA04\src\files\pg100.txt
Scanner fil......
Bytes:  5.589889e+06
Kibibytes:  5589.889
Mibibytes:  5.589889
Mibibytes:  0.005589889
Dette er en fil
Filrettigheter:  -rw-rw-rw-
Filen er ikke 'append only'
Filen er ikke 'name pipe'
Filen er ikke  'device oppgaver'
Filen er ikke  ' Unix character device'
Filen er ikke  'symbolic link'
```

<br>
Test av pg100.txt på linux:

```
Skriv filbane på filen du vil scanne:
pg100.txt
Bytes:  5.589889e+06
Kibibytes:  5589.889
Mibibytes:  5.589889
Mibibytes:  0.005589889
Dette er en fil
Filrettigheter:  -rw-rw-r--
Filen er ikke 'append only'
Filen er ikke 'name pipe'
Filen er ikke  'device oppgaver'
Filen er ikke  ' Unix character device'
Filen er ikke  'symbolic link'
```


<br>
Test av pg100.txt på mac:

```
Skriv filbane på filen du vil scanne:
gitGroup/IS-105-Gruppe1/ICA04/src/files/pg100.txt
Scanner fil......
Bytes:  5.4651e+06
Kibibytes:  5465.1
Mibibytes:  5.4651
Mibibytes:  0.0054651
Dette er en fil
Filrettigheter:  -rw-r--r--
Filen er ikke 'append only'
Filen er ikke 'name pipe'
Filen er ikke  'device oppgaver'
Filen er ikke  ' Unix character device'
Filen er ikke  'symbolic link'
logout
Saving session...completed.

```

<br>
Ved print av filen pg100.txt fant vi ingen forskjeller ved å kjøre samme go fil på alle tre platformer.

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
b) Total antall runer og antall for fem runes som forekommer mest i fil:
```
Oppgave 3b:
Skriv inn filen du ønsker å scanne:
files/pg100.txt
Antall linjeskift i filen: 800

5 mest brukte runes:
Antall: 6679 Rune:
Antall: 2718 Rune: e
Antall: 1962 Rune: t
Antall: 1615 Rune: o
Antall: 1372 Rune: s
```
<br>
c) Kjør benchmark "go test -bench=." i benchmark mappen

```
    3000            463270 ns/op
PASS
ok      _/C_/OneDrive/Skole/ICA/ICA04/src/oppgaver/benchmark    3.315s
```

#### Oppgave 4
a) Sannsynligheten regner vi ved å ta Muligheter / total * 100 %
```
                   Fakulitet |    Sannsynlighet 2015 |
                             |                       |
         Helse og Idrettsdag |  17.346358118361152 % |
     Humaniora og pedagogikk |  14.463201820940819 % |
                    Kunstfag |   3.983308042488619 % |
        Teknologi og realfag |   20.54248861911988 % |
              Lærerutdanning |  14.283004552352049 % |
Økonomi og samfunnsvitenskap |   29.38163884673748 % |
```
<br>

b)  Fakulitetet Kunstfag får minst informasjon som sett over.
```
                   Fakulitet |    Informasjon |
                             |                |
         Helse og Idrettsdag |   11100100101  |
     Humaniora og pedagogikk |   10111110101  |
                    Kunstfag |     110100100  |
        Teknologi og realfag |  100001110110  |
              Lærerutdanning |   10111100010  |
Økonomi og samfunnsvitenskap |  110000011010  |
```
<br>
c) Huffmancode for alle fakulitet <br>

![alt text](https://github.com/Zwirc/IS-105/blob/master/ICA04/vedlegg/huffman.PNG "Wapp")

```
                   Fakulitet | Informasjon |
                             |             |
         Helse og Idrettsdag |        111  |
     Humaniora og pedagogikk |        110  |
                    Kunstfag |        000  |
        Teknologi og realfag |         01  |
              Lærerutdanning |        001  |
Økonomi og samfunnsvitenskap |         10  |
```

<br>
d) Antall * (Lengden på melding * sansynlighet) + (Lengden på ......

```
100*(3*0.1734) + (3 * 0.1446) + (3 * 0.0398) + (2 * 0.2054) + (3 * 0.1428) + (2 * 0.2938)<br>
```
   Lengde på melding til 100 stk fra huffman bit lengde: 54
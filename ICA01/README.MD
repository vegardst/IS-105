# Oppgave 1.2.1

#####Konvertering av desimaltall til 2-tallsystemet (binært tallsystem):

Her skal vi regne ut desimaltall til binære tall, for å forstå hvordan datamaskinen leser disse tallene.
Grunnen til at datamaskiner bruker binære tall henger sammen med hvordan digitale kretser er bygd opp (f.eks OG, NOG, NELLER-kretser) med logiske porter. Binære tall blir derfor tolket som 1-Sann , 0-Usann.

Man må kjenne det binære tallsystemet for å forstå hvordan data lagres og aritmetiske operasjoner blir utført i en datamaskin.

Tallene i en mikroprosessor er representert med et spenningsnivå hvor en transistor enten blir slått av eller på, for eksempel:

0 volt = 0 binært, 5 volt = 1 binært, kilde: http://www.it.hiof.no/grit/forelesning/datamaskin/digital.html

For å telle binære tall leses disse ,"Arabisk" , fra høyre til venstre. Og imotsetning til titallssystemet, blir binære tall hverdier to ganger så høye for hver siffer plassering til venstre.

Eksempel: 

<table style="height: 59px;" border="1" width="311">
<tbody>
<tr>
<td style="text-align: center;">2^7</td>
<td style="text-align: center;">2^6</td>
<td style="text-align: center;">2^5</td>
<td style="text-align: center;">2^4</td>
<td style="text-align: center;">2^3</td>
<td style="text-align: center;">2^2</td>
<td style="text-align: center;">2^1</td>
<td style="text-align: center;">2^0</td>
</tr>
<tr>
<td style="text-align: center;">128</td>
<td style="text-align: center;">64</td>
<td style="text-align: center;">32</td>
<td style="text-align: center;">16</td>
<td style="text-align: center;">8</td>
<td style="text-align: center;">4</td>
<td style="text-align: center;">2</td>
<td style="text-align: center;">1</td>
</tr>
<tr>
<td style="text-align: center;"><strong>1</strong></td>
<td style="text-align: center;"><strong>0</strong></td>
<td style="text-align: center;"><strong>1</strong></td>
<td style="text-align: center;"><strong>0</strong></td>
<td style="text-align: center;"><strong>1</strong></td>
<td style="text-align: center;"><strong>0</strong></td>
<td style="text-align: center;"><strong>1</strong></td>
<td style="text-align: center;"><strong>1</strong></td>
</tr>
</tbody>
</table>

Så for å regne ut dette legger vi sammen verdien for 1 bitene.
128 + 32 + 8 + 2 + 1 = 171


Nå vi skal regne fra desimal til binært tall som i oppgaven, går vi motsatt vei.Dvs at vi begynner med det høyest mulig tallet i rekken 

Så ser vi på tabellen over, blir 1 = 1

Tallet 2 får verdien 10

Tallet 5, har både verdien 4 og 1, og blir derfor 101 (Vi teller fra det høyeste tallet)

* (1) 1 = 1 (1bit)
* (2) 2 = 10 (2bit)
* (3) 5 = 101 (3bit) 
* (4) 8 = 1000 (4bit) 
* (5) 16 = 10000 (5bit) 
* (6) 256 = 100000000 (9bit)

<table style="height: 62px;" border="1" width="417">
<tbody>
<tr>
<td style="text-align: center;">Desima ltall</td>
<td style="text-align: center;">&nbsp;1</td>
<td style="text-align: center;">&nbsp;2</td>
<td style="text-align: center;">&nbsp;5</td>
<td style="text-align: center;">&nbsp;8</td>
<td style="text-align: center;">&nbsp;16</td>
<td style="text-align: center;">&nbsp;256</td>
</tr>
<tr>
<td style="text-align: center;">&nbsp;Bin&aelig;r tall</td>
<td style="text-align: center;">&nbsp;1</td>
<td style="text-align: center;">&nbsp;10</td>
<td style="text-align: center;">&nbsp;101</td>
<td style="text-align: center;">&nbsp;1000</td>
<td style="text-align: center;">10000&nbsp;</td>
<td style="text-align: center;">100000000</td>
</tr>
</tbody>
</table>
For å regne ut tall der du ikke finner verdien direkte i tallsystemet, slik som tallet 5 tar du binær tallet som er nermest, f.eks med tallet 22:

- 22 = 16 + 6 (Men siden 6 ikke finnes i totalssystemet, må vi gå ned en verdi)
- 22 = 16 + 4 + 2 (Da har vi funnet tallene i totalssystemet, og setter de sammen til binær kode)
- 22 = 2^4 + 2^2 + 2^1 
- 22 = 10110

Når vi jobber med både 2talls og 10talls systemet er det vanlig å skrive bak tallene hvilket system de tilhører: 10110 (to)

##### 1) Konverter følgende binære tall til desimaltall (mest signifikante bit-en er til venstre):
Hvis vi vil konvertere binære tall til tall i titallssystemet kan vi legge sammen plassverdien til alle 1 sifrene.

- (1) 100 = 2^2 = 4
- (2) 1001 = 2^3 + 2^0 = 9
- (3) 1100110011 = 819
<table style="height: 40px;" border="1" width="436">
<tbody>
<tr>
<td style="text-align: center;">2^9</td>
<td style="text-align: center;">2^8</td>
<td style="text-align: center;">2^7</td>
<td style="text-align: center;">2^6</td>
<td style="text-align: center;">2^5</td>
<td style="text-align: center;">2^4</td>
<td style="text-align: center;">2^3</td>
<td style="text-align: center;">2^2</td>
<td style="text-align: center;">2^1</td>
<td style="text-align: center;">2^0</td>
</tr>
<tr>
<td style="text-align: center;">1</td>
<td style="text-align: center;">1</td>
<td style="text-align: center;">0</td>
<td style="text-align: center;">0</td>
<td style="text-align: center;">1</td>
<td style="text-align: center;">1</td>
<td style="text-align: center;">0</td>
<td style="text-align: center;">0</td>
<td style="text-align: center;">1</td>
<td style="text-align: center;">1</td>
</tr>
</tbody>
</table>
= 2^9 + 2^8 + 2^5 + 2^4 + 2^1 + 2^0 
= 819


Resultat:
<table style="height: 40px;" border="1" width="291">
<tbody>
<tr>
<td style="text-align: center;">Bin&aelig;rtall</td>
<td style="text-align: center;">100</td>
<td style="text-align: center;">1001</td>
<td style="text-align: center;">1100110011</td>
</tr>
<tr>
<td style="text-align: center;">&nbsp;Desimaltall</td>
<td style="text-align: center;">&nbsp;4</td>
<td style="text-align: center;">9&nbsp;</td>
<td style="text-align: center;">&nbsp;819</td>
</tr>
</tbody>
</table>
Informasjonskilde: Tek.no

# Oppgave 1.2.2

Informasjonsmengde Flere personer prøver å gjette et tresifret (3-bit) binært tall.

- (1) Lise har fått vite / lærer at tallet er et oddetall.
- (2) Per har fått vite at tallet er IKKE et multiplum av 3 (dvs. ikke 0, 3, 6). 
- (3) Oskar har fått vite at tallet inneholder nøyaktig 2 enere. 
- (4) Louise har fått vite alt det Lise, Per og Oskar vet.

3 bit binære tall gir mulighetene:
<table style="height: 154px;" border="1" width="251">
<tbody>
<tr>
<td style="text-align: center;">000</td>
<td style="text-align: center;">0</td>
</tr>
<tr>
<td style="text-align: center;">001</td>
<td style="text-align: center;">1</td>
</tr>
<tr>
<td style="text-align: center;">010</td>
<td style="text-align: center;">2</td>
</tr>
<tr>
<td style="text-align: center;">011</td>
<td style="text-align: center;">3</td>
</tr>
<tr>
<td style="text-align: center;">100</td>
<td style="text-align: center;">4</td>
</tr>
<tr>
<td style="text-align: center;">101</td>
<td style="text-align: center;">5</td>
</tr>
<tr>
<td style="text-align: center;">110</td>
<td style="text-align: center;">6</td>
</tr>
<tr>
<td style="text-align: center;">111</td>
<td style="text-align: center;">7</td>
</tr>
</tbody>
</table>
1. Lise vet at det er oddetall, og står igjen med tallene: 1, 3, 5, 7
<table style="height: 154px;" border="1" width="251">
<tbody>
<tr>
<td style="text-align: center;">001</td>
<td style="text-align: center;">1</td>
</tr>
<tr>
<td style="text-align: center;">011</td>
<td style="text-align: center;">3</td>
</tr>
<tr>
<td style="text-align: center;">101</td>
<td style="text-align: center;">5</td>
</tr>
<tr>
<td style="text-align: center;">111</td>
<td style="text-align: center;">7</td>
</tr>
</tbody>
</table>
2. Per vet at tallet IKKE er multiplum av 3 (multiplum er et tall som forekommer når tallet multipliseres med et helt tall.)

- 3*0 = 0
- 3*1 = 3
- 3*2 = 6

Og Per står igjen med tallene: 1, 2, 4, 5, 7
<table style="height: 154px;" border="1" width="251">
<tbody>
<tr>
<td style="text-align: center;">001</td>
<td style="text-align: center;">1</td>
</tr>
<tr>
<td style="text-align: center;">010</td>
<td style="text-align: center;">2</td>
</tr>
<tr>
<td style="text-align: center;">100</td>
<td style="text-align: center;">4</td>
</tr>
<tr>
<td style="text-align: center;">101</td>
<td style="text-align: center;">5</td>
</tr>
<tr>
<td style="text-align: center;">111</td>
<td style="text-align: center;">7</td>
</tr>
</tbody>
</table>

3.  Oskar har fått vite at tallet inneholder nøyaktig 2 enere som er tallene 3,5,6

<table style="height: 154px;" border="1" width="251">
<tbody>
<tr>
<td style="text-align: center;">011</td>
<td style="text-align: center;">3</td>
</tr>
<tr>
<td style="text-align: center;">101</td>
<td style="text-align: center;">5</td>
</tr>
<tr>
<td style="text-align: center;">110</td>
<td style="text-align: center;">6</td>
</tr>
</tbody>
</table>

(4) Louise har fått vite alt det Lise, Per og Oskar vet. 

Så Louise vet at tallet er 5 , binærtall 101
<table style="height: 173px;" border="1" width="303">
<tbody>
<tr>
<td>Lise</td>
<td>Per</td>
<td>Oscar</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;1</td>
<td>1</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>2</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;3</td>
<td>&nbsp;</td>
<td>3</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>4</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;5</td>
<td>5</td>
<td>5</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>6</td>
</tr>
<tr>
<td>&nbsp;7</td>
<td>7</td>
<td>&nbsp;</td>
</tr>
</tbody>
</table>

##### Hvor mye informasjon (i bits) har hver spiller fått?

Det er totalt 8 muligheter med 3 sifret binære tall: N=8 som gir Log2(8) = 3 bit
Formel brukt i oppgave: log2(8/mulige valg)

1) Lise vet at tallet er odetall, N=8/4
<br>Lise sin informasjon i bits Log2(8/4) = 1 bit. Så Lise har 2 bit igjen å gjette.

2) Per vet at tallet ikke er multiplum av 3, og står igjen med 5 tall. N=8/5
<br>Per sin informasjon i bits Log2(8/5) = 0 bit. Per får under 1 bit informasjon å må gjette 3 bits ut av 5 mulige valg.

3) Oskar vet at tallet inneholder nøyaktig 2enere som gir 3 muligheter. N=8/3
<br>Oskar sin informasjon i bits Log2(8/3) = 1 bit. Per har like mye informasjon som Lise og har 2 bit igjen å gjette.

4) Louise har fått vite alt Lise,Per og Oskar vet = log2(8/1)
<br>Louse sin informasjon i bits 3 . Louise har 3Bits informasjon og 1 mulig valg

# Oppgave 1.2.3

Alle har tilgang til repository is-105-uke04 og har laget egen branch og lastet opp eksempel filer (Se branch liste: https://github.com/IS-105-GitGroup/is105-uke04 )

Vi brukte følgene mal:

Gå inn på mappen dere vil "lagre" Repoen på. (f.eks mkdir GitGroup -> cd GitGroup) 
1. Last ned med git clone kommandoen 
git clone https://github.com/IS-105-GitGroup/is105-uke04.git 

2. Lag en ny branch med deres navn. 
git checkout -b Navn 
Du ser nå at du er i branchen ved at branchnavn står etter mappestrukturen: GitGroup/is105-uke04 (Navn) 

3. Opprett en ny .go Fil i Go mappa enten med kommandolinje eller OS brukergrensesnitt. 
Når du har lagret filen kan du skrive: git status og vil se at filen er rød der. 

4. Legg til filen, commit og last opp. 
git add . : "." gjør at du legger til alle filer, du kan også skrive filnavn. 
git commit -m "kommentar" : Commit filen med kommentar 
git push origin Navn : Laster opp branchen Navn til Repoen vår. 

For å oppdatere lokale filer, kan git pull origin BRANCH-NAVN brukes.

# Oppgave 1.2.4

##### 1) Hvilken fordeler og ulemper har en git-flow-modell med en hovedrepository?

Fordelen med en hovedrepo er at man kan kopiere denne og jobbe separat på forskjellige elementer av koden for så å sammenligne / legge til endringer i hovedrepo når endringene er klare. Ved å bruke en "public" repo er det også enkelt å dele koden, få tilbakemeldinger og se endringer.

 

Ulempen kan være at andre kan slette nødvendige filer / endre hovedrepo istedenfor test brancher, man kan være uheldig å dele passord og essensiel data mm.
Andre ulemper kan være at man lager mange repo, har lite dokumentasjon og ender opp med bare rot.

 

##### 2) Finn ut hva heter objektfiler for de mest brukte platformer (Unix/Linux, MS Windows, Mac OS X)! Hvorfor, etter deres mening, har disse platformene så forskjellige objektfil-formater

ELF brukes for det meste på platformer, ellers har vi:

a.out AIF COFF CMD COM ECOFF GOFF Hunk Mach-O MZ NE OMF OS/360 PE PEF XCOFF

Disee pakkes i typiske kjørbare filer for operativsystemet, f.eks .exe for Windows.
Andre typiske kjørbare filer er :
.acm, .ax, .cpl, .dll, .drv, .efi, .mui, .ocx, .scr, .sys, .tsp

Operativsystemene kjører på forskjellig "Architectures":
Linux (x86, PowerPC, ARM, and others)
Microsoft Windows (x86, ARM)
macOS (x86, PowerPC (on 10.5 and below))

 

##### 3) Hvilke forskjeller er det fra Java og GoLang?

Go - Mindre kode og i stor vekst.
Go - Mer rettet mot HTTP
Java - Har fortsatt større community
Java - Mange frameworks og IDE support

 

##### 4) Hvilke viktige poeng illustrerer denne øvelsen når det gjelder bruk av et programmeringsmiljø på en platform?

Poenget denne øvelsen illustrerer er siden jeg bygger filen på Windows, så får jeg en kjørbar exe fil.

Med GoLang kan man også "cross compile" ved å sette "GOOS" og "GOARCH" til verdiene for operativsystems arkitekturen programmet skal kjøre på. Dermed "go build -v pakke" 

I motsetning til Java, må vi i dette tilfellet kjøre begge filene hvis vi skal ha de i samme mappe: "go run main.go  log.go" , så velger den main metoden fra main.go for å kjøre, se eksempler under GoLang.
For mer info se: https://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5

Skal man kun kjøre bare "main.go" (go run main.go) må filen Log importeres ved å legges i en subfolder dersom vi ikke endrer "%GOPATH%"


NB! Om man prøver å importere en mappe der main filen allerede er får man error, import "../GoLang":
"main.go:5:2: found packages log (log.go) and main (main.go) "

For mer utfyllende info se: http://gowithconfidence-blog.tumblr.com/post/118493925546/getting-started-with-go-workspaces

 

##### 5) Bruk eksemplet fra https://gobyexample.com/command-line-arguments og lag et utførbart program logcli?, som kan ta et argument (tallet som man skal beregne logaritmen av base 2 for);  Er det hensiktsmessig å legge inn denne filen i git repository? Begrunn svaret!

Filen logcli.go i Oppgaver/GoLang på master branch kjøres ved "go run logcli.go 10 11" der første parameter er LogBase-2 og andre parameter er LogBase-10. Det er også brukt import av "strconv" som konverterer string til Float64 for å bruke denne i math.Log metodene.
Hvis man har bruk for denne utregningen, er det hensiktsmessig å ha filen i sitt repository slik at man enkelt kan gjøre utregninger senere ved å kalle på metoden.

##### 6) Hva skille pakke log, fra andre pakker i go?

Hver pakke har forskjellige metoder som utfører forskjellige funksjoner.

I dette tilfellet importeres log lokalt, mens fmt er en del av standardbibloteket, og metoder finnes i oversikten
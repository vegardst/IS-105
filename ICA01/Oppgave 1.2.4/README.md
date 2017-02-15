<h1># Oppgave 1.2.4</h1><p>
<p>&nbsp;<strong>1) Hvilken fordeler og ulemper har en git-flow-modell med en hovedrepository?<br /><br /></strong>Fordelen med en hovedrepo er at man kan kopiere denne og jobbe separat p&aring; forskjellige elementer av koden for s&aring; &aring; sammenligne / legge til endringer i hovedrepo n&aring;r endringene er klare. Ved &aring; bruke en "public" repo er det ogs&aring; enkelt &aring; dele koden, f&aring; tilbakemeldinger og se endringer.</p>
<p>&nbsp;</p>
<p>Ulempen kan v&aelig;re at andre kan slette n&oslash;dvendige filer / endre hovedrepo istedenfor test brancher, man kan v&aelig;re uheldig &aring; dele passord og essensiel data mm.<br />Andre ulemper kan v&aelig;re at man lager mange repo, har lite dokumentasjon og ender opp med bare rot.</p>
<p>&nbsp;</p>
<p><strong><strong>2) Finn ut hva heter objektfiler for de mest brukte platformer&nbsp;</strong></strong><strong>(Unix/Linux, MS Windows, Mac OS X)! Hvorfor, etter deres mening, har disse&nbsp;</strong><strong>platformene s&aring; forskjellige objektfil-formater<br /><br /></strong>ELF brukes for det meste p&aring; platformer, ellers har vi:</p>
<p>a.out AIF COFF CMD COM ECOFF GOFF Hunk Mach-O MZ NE OMF OS/360 PE PEF XCOFF</p>
<p>Disee pakkes i typiske kj&oslash;rbare filer for operativsystemet, f.eks .exe for Windows.<br />Andre typiske kj&oslash;rbare filer er :<br />.acm, .ax, .cpl, .dll, .drv, .efi, .mui, .ocx, .scr, .sys, .tsp</p>
<p>Operativsystemene kj&oslash;rer p&aring; forskjellig "Architectures":<br />Linux (x86, PowerPC, ARM, and others)<br />Microsoft Windows (x86, ARM)<br />macOS (x86, PowerPC (on 10.5 and below))</p>
<p>&nbsp;</p>
<p><strong style="font-weight: bold;">3) Hvilke forskjeller er det fra Java og GoLang?<br /><br /></strong>Go - Mindre kode og i stor vekst.<br />Go - Mer rettet mot HTTP<br />Java - Har fortsatt st&oslash;rre community<br />Java - Mange frameworks og IDE support</p>
<p>&nbsp;</p>
<p><strong style="font-weight: bold;">4) Hvilke viktige poeng illustrerer denne &oslash;velsen n&aring;r det gjelder bruk av et programmeringsmilj&oslash; p&aring; en platform?<br /><br /></strong>Poenget denne &oslash;velsen illustrerer er siden jeg bygger filen p&aring; Windows, s&aring; f&aring;r jeg en kj&oslash;rbar exe fil.</p>
<p>Med GoLang kan man ogs&aring; "cross compile" ved &aring; sette "GOOS" og "GOARCH" til verdiene for operativsystems arkitekturen programmet skal kj&oslash;re p&aring;. Dermed "go build -v pakke"&nbsp;</p>
<p>I motsetning til Java, m&aring; vi i dette tilfellet kj&oslash;re begge filene hvis vi skal ha de i samme mappe: "go run main.go&nbsp; log.go" , s&aring; velger den main metoden fra main.go for &aring; kj&oslash;re, se eksempler under GoLang.<br />For mer info se: <a href="https://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5">https://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5</a></p>
<p>Skal man kun kj&oslash;re bare "main.go" (go run main.go) m&aring; filen Log importeres ved &aring; legges i en subfolder dersom vi ikke endrer "%GOPATH%"</p>
<p><br />NB! Om man pr&oslash;ver &aring; importere en mappe der main filen allerede er f&aring;r man error, import "../GoLang":<br />"main.go:5:2: found packages log (log.go) and main (main.go) "</p>
<p>For mer utfyllende info se: <a href="http://gowithconfidence-blog.tumblr.com/post/118493925546/getting-started-with-go-workspaces">http://gowithconfidence-blog.tumblr.com/post/118493925546/getting-started-with-go-workspaces</a></p>
<p>&nbsp;</p>
<p><strong><strong>5) Bruk eksemplet fra </strong><a href="https://gobyexample.com/command-line-arguments"><strong>https://gobyexample.com/command-line-arguments</strong></a><strong> og&nbsp;</strong></strong><strong>lag et utf&oslash;rbart program logcli?, som kan ta et argument (tallet som man&nbsp;</strong><strong>skal beregne logaritmen av base 2 for); &nbsp;</strong><strong>Er det hensiktsmessig &aring; legge inn&nbsp;</strong><strong>denne filen i git repository? Begrunn svaret!<br /><br /></strong>Filen logcli.go i Oppgaver/GoLang på master branch kjøres ved "go run logcli.go 10 11" der første parameter er LogBase-2 og andre parameter er LogBase-10. Det er også brukt import av "strconv" som konverterer string til Float64 for å bruke denne i math.Log metodene.<br>
   Hvis man har bruk for denne utregningen, er det hensiktsmessig &aring; ha filen i sitt repository slik at man enkelt kan gj&oslash;re utregninger senere ved &aring; kalle p&aring; metoden.</p>
<p><strong>6) Hva skille pakke log, fra andre pakker i go?<br /><br /></strong>Hver pakke har forskjellige metoder som utf&oslash;rer forskjellige funksjoner.</p>
<p>I dette tilfellet importeres log lokalt, mens fmt er en del av standardbibloteket, og metoder finnes i oversikten</p>
<p>&nbsp;</p> 

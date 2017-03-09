# ICA 04
Denne ICA er også satt opp med scanner, så du kjører main.go og går igjennom oppgavene i terminal.
<br>
Oppgave 4b vil be om tilgang til å starte serveren på localhost:3000
<br>
<br>
#### Oppgave 1:
a) Her kan man se i kildekoden hvordan utskriften gjøres.<br>
I mitt tilfelle vil ikke 0-1F og 7F symbolene vises med kodeopsettet jeg har i terminal.
<br>
<br>
b) Resultatet ligger i vedlegg mappa :)
<br>
<br>
c) Kjør oppgave 1a i main.go.<br>
<br>
For oppgave 1 ser vi at man ikke nødvendigvis får vist symbolet da oppsettet i terminalen bestemmer om du bruker et kodesett som kjenner symbolet.
<br>
<br>
#### Oppgave 2:
a) I denne oppgaven ser vi igjen hvilken utfordring det som dukker opp ved å prøve å vise "ukjente" symboler.
<br>
<br>
b) Se vedlegg for utførelse på andre platformer.
<br>
<br>
c) Testene ligger i src/test mappen
<br>
<br>
#### Oppgave 3:
a) Se kode for oppgave 3a. Som beskrevet i fmt pakka har hver prefix forskjellig resultat på utskriften.
<br>
<br>
b) Denne oppgaven var veldig vanskelig å løse, prøvde å implementere mange pakker som ikke lå i standard instalasjonen uten hell og resultatet ble at jeg måtte definere resultatet i HEX som en egen verdi for så å skrive ut denne.<br>
Vi kom fram til en løsning i timen som jeg har jobbet med og implementert i koden for oppgave 3b.
<br>
<br>
c) Kjør oppgave 3b i main filen og få Henrik Arnold Wergelands utskrift :)
<br>
<br>
#### Oppgave 4:
a) Strengen representeres i HEX og vises ved å kjøre 4a.
<br>
<br>
b) Siden jeg satt "text/html;charset=ISO-8859-1" i headeren, bruker jeg dette charset for å vise klokke symbolet og bruker time pakken for å gjøre en utskrift av dato, tid.

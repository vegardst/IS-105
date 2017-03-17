#Oppgaver for IS-105
#### Hvor starter jeg?
Alle prosjekt er laget under ICA som root source.
<br>
 Last ned GoLang på: https://golang.org/dl/
Last ned prosjektet lokalt: git clone LINK-TIL-REPO
 Anbefaler å åpne prosjektet i Go IDE fra JetBrains: https://www.jetbrains.com/go/ 
Når du har åpnet prosjektet i Gogland, ser du følgende oversikt over ICA moduler:
<br>
![alt text](https://github.com/Zwirc/IS-105/blob/master/HowTo/Gogland.png "Wapp")
<br> Nyttige funksjoner:
Gogland hjelper deg automatisk å finne Go, og du kan enkelt kjøre go filer fra Gogland.
<br>
Det er også mulighet for å åpne terminal direkte i Gogland, eller automatisk pushe til Git repoen din.<br>
![alt text](https://github.com/Zwirc/IS-105/blob/master/HowTo/Gogland2.png "Wapp")

<br>
Du lager ny fil ved å høyre-klikke på mappen du ønsker filen i, og velg Go File:
![alt text](https://github.com/Zwirc/IS-105/blob/master/HowTo/NewFile.png "Wapp")
<br>
<br>
Gogland har også autofullfør  funksjon:<br>

![alt text](https://github.com/Zwirc/IS-105/blob/master/HowTo/Autofullfor.png "Wapp")
<br>
Når du kjører en main.go fil, har du mulighet til å gi input og teste om programmet fungerer:
![alt text](https://github.com/Zwirc/IS-105/blob/master/HowTo/Kjoring-Av-MainGO.png "Wapp")

##How to git?
<br> Step by step:
<br> Gå inn på mappen dere vil "lagre" Repoen på. (f.eks mkdir GitGroup  -> cd GitGroup)
<br> 1. Last ned med git clone kommandoen
<br> <i><b> git clone https://github.com/IS-105-GitGroup/is105-uke04.git</i></b>
<br>
<br> 2. Lag en ny branch med deres navn.
<br> <i><b> git checkout -b Erlend </i></b>
<br> Du ser nå at du er i branchen ved at branchnavn står etter mappestrukturen: GitGroup/is105-uke04 (Erlend)
<br>
<br> 3. Opprett en ny .go Fil i Go mappa enten med kommandolinje eller OS brukergrensesnitt.
<br> Når du har lagret filen kan du skrive: <i><b> git status </i></b> og vil se at filen er rød der.
<br>
<br> 4. Legg til filen, commit og last opp.
<br> <i><b> git add . </i></b> : "." gjør at du legger til alle filer, du kan også skrive filnavn.
<br> <i><b> git commit -m "kommentar"</i></b> : Commit filen med kommentar
<br> <i><b> git push origin Erlend </i></b> : Laster opp branchen Erlend til Repoen vår.
<br>
<br> For å oppdatere lokale filer, kan <i><b> git pull origin BRANCH-NAVN </i></b> brukes.
<br>
<br> <b> Nyttige kommandoer: </b>
<br><b>git clone "url til repository"     </b>// clone et repository fra github f.eks.
<br><b>git status</b>				  // vil gi en oversikt over endrede filer
<br><b>git add "filename" 		  </b>// vil legge til filen til stage for commit
<br><b>git add .		  </b>// vil legge til alle filer som er endret til stage for commit
<br><b>git reset "filename"		</b>   // fjerne staged file for commit.
<br><b>git commit	-m “melding”	</b>  // pakke inn endringene klar til push
<br><b>git push origin "branchnavn" </b>// Legg ut commit’en til github.
<br><b>git checkout -b "branchnavn" </b>// opprette en ny branch
<br><b>git checkout "branchnavn" 	  </b>// vil gå til branch
<br><b>git merge "branchnavn"	</b>// merger branchen man er i til spesifisert branch. 
#ICA-05 

###Systemspesifikasjon
Systemet vi har utviklet skal kunne gi brukeren informasjon om lokasjon, vær på lokasjon stedet, kart over lokasjonen, lokasjonsinformasjon og informasjon om brukeren som kun brukeren kan se, type IP-adresse, land, by, internettleverandør. Systemet skal inneholde en søkefunksjon slik at brukeren kan søke seg fram til lokasjoner som brukeren ønsker informasjon om. Informasjonen som skal vises er stedsnavn (land og by/sted), landskode, nåværende temperatur, vindstyrke og hvilket vær det er. I tillegg vise relevant informasjon om lokasjonen. Brukeren skal også få opp et kart over valgt lokasjon hvor han også har muligheten til å “dra” seg rundt og zoome i kartet for evt å finne nye lokasjoner hvis han er usikker på steder han ønsker å søke opp. 

###Brukerscenario 
 Brukeren besøker siden for å ha mulighet til å søke opp informasjon om en lokasjon. 

Siden laster dermed ned informasjon fra forskjellige sider avhengig av søk 
 
 #####Brukerscenario 1 
 Knut og Harald reiser rundt i Frankrike, mens de reiser bruker de nettsiden til å finne ut hvordan været blir i de byen Marseille som de har tenkt å besøke og litt informasjon om byen. De finner ut at det er vindstille og høy temperatur, og bestemmer seg for å reise til stranden. 

 #####Brukerscenario 2 
 Maia skal seile fra Oslo til Bergen, hun bruker nettsiden til å holde seg oppdatert på været mens hun reiser. Før hun kommer til Flekkefjord, ser hun på nettsiden at vinden har økt betraktelig og at fuktigheten har økt, og at det er overskyet og regner. Hun bestemmer seg for søke opp Lyngdal som er på veien og ser at det kan være en spennende plass for å vente på bedre vær. 

 ###Testing / Pålitelig tjeneste 
 Det er lagt inn feilhåndtering slik at siden ikke krasjer hvis det ikke finnes informasjon på søket i de forskjellige datasettene. Resultatet blir da enten tomt eller gir feilmelding avhengig av datakilde. 
 Dette er da testet ved å gi forskjellig input med ukjente tegn og verdier uten at serveren krasjer. 

 Eksempel: " Could not load Wiki for AS%C3%A5%5Cf2%C2%A8%60%22%60%C2%A8qw%27 ,There may be a error with the search! " 

Andre feil:
Feil ved kjøring på "skoleserver": 
net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers) fix: netstat -i og set mtu for Iface.   

###Serveren 
Serveren er satt opp på Ubuntu server og kjøres via go miljøet. 
UFW er brukt for å åpne porter der serveren kjører. 

Server.go starter en webserver som henter inn informasjon fra forskjellige json data (se kilder), og viser dette for brukeren som nevnt i brukerscenario. 
 
* Siden bruker besøkendes IP for å finne brukerens lokasjon 
* Siden bruker landskode (eks NO) fra weatherdata søket for å fortelle hvilken land lokasjonen er i. 
* Siden gir realtids-data , da data hentes når siden åpnes. 

For egenlæring er det lagt til funksjoner, logging av besøk og diverse effekter på siden.

####Applikasjonen , Go program (WebClient) med concurrency 
Serveren bruker allerede gorutiner for å hente inn flere datasett samtidig, men det er også vedlagt et enkelt go-program som viser hvordan man kan bruke data likt i en applikasjon. 

###Videreutvikling
Legge til relevante bilder / twitter meldinger eller lignende for lokasjonen
Legge til arrangementer i nærheten, bruke Facebook eller annet som inneholder informasjon om arrangementer i nærheten
Legge til hvor mye klokka er på en angitt lokasjon, eks: Hvis du søker Kristiansand får du opp klokka lokal tid for Kristiansand, når du søker eks New York får du opp lokaltiden for New York som eksempelvis vil være 6 timer bak oss. Og evt hvilken tidssone de ligger i forhold til hvor du søker fra.
Legge til valuta til angitt lokasjon


###OS Modellen



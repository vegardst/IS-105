# Oppgave 1

Formål: bli kjent med systemet sitt  :+1:

#### Hvor mange prosesser som kjører på din datamaskin?
Vi kan se hvor mange prosesser som kjører på Windows ved å bruke taskmgr, enten ved å kjøre programmet ved å skrive "taskmgr" i terminal, eller skrive "tasklist" for å få liste over alle prosesser i terminal.

På min maskin kjører 100 prosesser for øyeblikket.


#### Hvor mange prosesser som kjører på din virtuelle server i nettskyen?
For ubuntu skriver jeg "top" i terminal for å få opp tilsvarende oversikt som viser 147 prosesser.



#### Kan man gi et nøyaktig antall? Begrunn.
Da jeg tolker prosesser som "Tasks" (set av program instruksjoner som er lastet i minnet) så kjører min server 147 prosesser, som vises i oversikten.


#### Hvor mange av prosessene som “kjører”?
Av prosessene på min server er det bare 3 som kjører.


#### Hvis de ikke kjører, hvilke tilstander befinner de seg da?
Prosessene som ikke kjører er i tilstanden idle/sleep.


#### Hva er maskinvarespesifikasjon til din datamaskin (noter prosessortype, prosessorarkitektur, klokkefrekvens, informasjon om primært minne, størrelse på cache (både L1, L2 og L3 er ønskelig))?
Har laget en "ComputerInfo.go" fil for å hente ut denne informasjonen for windows / linux. 

- Lokal maskinvarespesifikasjoner:Prosessortype: Intel(R) Core(TM) i7-4770K , 4-core 8-threads
- Prosessorarkitektur: Intel`s Haswell, 64bit
- Klokkefrekvens:  4.1GHz 
- Informasjon om primært minne: 16GB DIMM DDR3 @ 1600 MHz
- Størrelse på cache ( 256 , 1024 , 8192)


#### Hvor mange CPU-“cores” har du tilgjengelig på din maskin? Noter.
Jeg har 4-Cores på min maskin som kjører HyperThreading teknologi


#### Hvor mange CPU-”cores” har du tilgjengelig på din virtuelle server? Noter.
Ved å bruke kommandoen "lscpu" ser jeg at jeg har 1 CPU med 1 Core på min virtuelle server.


#### Finn ut hvilken prosess i ditt system bruker mest minne. Beskrive denne prosessen kort.
Min datamaskin bruker mest minne på IntelliJ da jeg programmerer og kjører virtuelle prosesser med programmet.


#### Teamarbeid: Oppsummer alle data i en tabell i deres team-besvarelsen.
- Sammenlign deres platformer og diskuter forkjeller

<table style="height: 154px;" width="710">
<tbody>
<tr>
<td><strong>Navn</strong></td>
<td><strong>OS</strong></td>
<td><strong>CPU</strong></td>
<td><strong>RAM</strong></td>
<td><strong>HK</strong></td>
<td><strong>GPU</strong></td>
</tr>
<tr>
<td>Erlend W</td>
<td>Windows 10 pro</td>
<td>Intel i7 4770K, 3.5GHz</td>
<td>16GB 1333 MHz</td>
<td>Asus Z87 PRO Socket 1150</td>
<td>Nvidia 1070 MSI</td>
</tr>
<tr>
<td>Eirik S</td>
<td>Windows 7 home premium</td>
<td>Intel core i7-3610QM, 2.30 GHz</td>
<td>8 GB 798 MHz</td>
<td>ASUSTeK K55VM (Socket 0)</td>
<td>NVIDIA GeForce GT 630M</td>
</tr>
<tr>
<td>Helle T</td>
<td>Windows 10 Home 64-bit</td>
<td>Inter Core i5 7500U, 2,50GHz Kaby Lake</td>
<td>8 GB 798 MHz</td>
<td>Acer Minicooper_SK (U3E1) </td>
<td>Intel HD Graphics 620</td>
</tr>
<tr>
<td>Brynjar B</td>
<td>OS X 10.11.6 El Capitan</td>
<td>Intel i7 2.2GHz I7-5650U</td>
<td>8GB 1600MHz DDR3</td>
<td>A1466</td>
<td>Intel HD Graphics 6000 1536 MB</td>
</tr>
<tr>
<td>Adrian J</td>
<td>Windows 10 pro</td>
<td>Intel i7 4790k, 4.6GHz</td>
<td>8GB 1600 MHz</td>
<td>Asus Z97A</td>
<td>Radeon R9 290x GIGABYTE</td>
</tr>
<tr>
<td>Daniel P</td>
<td>Mac OS X 10.12 Sierra</td>
<td>Intel Core i7-2640M 2,8 GHz Sandy Bridge</td>
<td>16GB RAM</td>
<td>A1278</td>
<td>Intel HD Graphics 3000</td>
</tr>
<tr>
<td>Mathias E</td>
<td>Windows 10 Home 64bit</td>
<td>Intel i7-3630QM, 2.4GHz</td>
<td>8GB 798MHz</td>
<td>Micro-Star International Co. Ltd. MS-16GA (SOCKET 0)</td>
<td>Nvidia GTX 660M</td>
</tr>
</tbody>
</table>

#### Hvilke komponenter (både fysiske og abstrakte) i deres datasystemer er involvert i oppstart, administrasjon og avslutning av prosesser? Definer komponentene du nevner
1. Power Supply : Når du trykker på POWER knappen, gir PSU strøm til hovedkortet og eksterne komponenter.
2. HK inneholder BIOS som sitter på grunnleggende instillinger for hvilke drivere som skal aktiveres og kjører en Power-On-Self-Test for så å starter bootloaderen.
3. BootLoaderen velger hvilken ressurs PCn skal starte fra (CDrom/HDD/USB)
4. Når bootloaderen finner OS, lastes dette inn på minnet og operativsystemet tar over oppstartsprosessen ved å laste inn sine drivere og instillinger.
5. Du kan nå logge inn på din PC.

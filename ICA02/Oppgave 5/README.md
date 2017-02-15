# Oppgave 5

WinTOP.bat kjører "TOP" for Windows.

* Programmet boring_main er bygget for Linux og kaller opp metoden i boring.go

* Jeg kjører programmet på min server, og kan få ut følgende info:
Ellers kan jeg bruke TOP/HTOP på serveren og kjøre den der:

#### Resultat Boring på Windows
![alt text](https://github.com/IS-105-GitGroup/IS-105-Gruppe1/blob/master/ICA02/Oppgave%205/Boring/BoringWin.PNG "Wapp")

#### Resultat Boring10 på Windows
![alt text](https://github.com/IS-105-GitGroup/IS-105-Gruppe1/blob/master/ICA02/Oppgave%205/Boring/Boring10Win.PNG "Wapp")


#####Hva er jeg?
```
get-help get-process -full


The default display of a process is a table that includes the following columns:

-- Handles: The number of handles that the process has opened.

-- NPM(K): The amount of non-paged memory that the process is using, in kilobytes.

-- PM(K): The amount of pageable memory that the process is using, in kilobytes.

-- WS(K): The size of the working set of the process, in kilobytes. The working set consists of the pages of me mory that were recently referenced by the process.

-- CPU(s): The amount of processor time that the process has used on all processors, in seconds.

-- ID: The process ID (PID) of the process.

-- ProcessName: The name of the process.
```


#### Resultat Boring på Linux server
![alt text](https://github.com/IS-105-GitGroup/IS-105-Gruppe1/blob/master/ICA02/Oppgave%205/Boring/Boring.PNG "Wapp")

#### Resultat Boring10 på Linux server
![alt text](https://github.com/IS-105-GitGroup/IS-105-Gruppe1/blob/master/ICA02/Oppgave%205/Boring/Boring10.PNG "Wapp")

* NLWP Viser antall "tråder" i prosessen.

* Jeg kan stoppe prosessen ved å trykke CTRL+C, ellers så har jeg satt opp at den skal kjøre 100 ganger før den stopper.

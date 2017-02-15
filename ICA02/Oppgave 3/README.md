# Oppgave 3

Programmeringsoppgave:

Kjør programmet:
```Go
go run main_sum.go sum.go
```

#### Se sum mappe for sum_tests.
Test som produserer feil:
* Se sum_tests_fail
```Go
var test_int32 = []struct {
	// Gir feil da det er int8
	n1       int32
	n2       int8
	expected int32
}{
	{1, 2, 3},
	{4, 5, 9},
	{123, 2, 125},
}

func TestSumInt32(t *testing.T) {
	for _, v := range test_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
```
Error melding:
```Error
.\sum_test_fail.go:43: cannot use v.n2 (type int8) as type int32 in argument to SumInt32
```
#### Main_sum.go tar inn parametre fra kommandolinjen
Kjør programmet og følg instruksene.
```
go run main_sum.go sum.go
```
#### Bruk av verdier
Man får feilmeldinger ved bruk av feil type, eller for høye verdier iforhold til type.
For å unngå dette, må man skrive kode som tar seg av feil input. 

#### Løsning som gjør bruken av pakken "sum" trygg for brukeren
Main_Sum.go bruker scanner input der du får velge hvilken type du vil regne ut.
Bruker du feil verdi iforhold til type, får du beskjed om dette. Har prøvd å gjøre den så trygg for bruker som mulig, men man erfarer fort at koden for å gjøre programmet bruker sikkert blir mye større enn selve koden for kalkulasjonen.


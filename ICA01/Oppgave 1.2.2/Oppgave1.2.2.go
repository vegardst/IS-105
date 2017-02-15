package main
// Oppgave 1.2.2
// Erlend.W

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Oppgave 1.2.2");
	fmt.Println("");
	fmt.Println("Det er totalt 8 muligheter med 3 sifret binære tall: N=8 som gir Log2(8) =" , math.Log2(8), "bit");
	fmt.Println("Formel brukt i oppgave: log2(8/mulige valg)");
	fmt.Println("");

	fmt.Println("1 ) Lise vet at tallet er odetall, N=8/4");
	var Lise = math.Log2(8/4);
	fmt.Println("Lise sin informasjon i bits Log2(8/4) = ", Lise, "bit. Så Lise har", 3-1, "bit igjen å gjette. ");
	fmt.Println("");

	fmt.Println("2) Per vet at tallet ikke er multiplum av 3, og står igjen med 5 tall. N=8/5");
	var Per = math.Log2(8/5);
	fmt.Println("Per sin informasjon i bits Log2(8/5) = ", Per, "bit. Per får under 1 bit informasjon å må gjette 3 bits ut av 5 mulige valg.");
	fmt.Println("");

	fmt.Println("3) Oskar vet at tallet inneholder nøyaktig 2enere som gir 3 muligheter. N=8/3");
	var Oskar = math.Log2(8/3);
	fmt.Println("Oskar sin informasjon i bits Log2(8/3) =", Oskar, "bit. Per har like mye informasjon som Lise og har 2 bit igjen å gjette.");
	fmt.Println("");

	fmt.Println("4) Louise har fått vite alt Lise,Per og Oskar vet = log2(8/1)");
	//var Louise = Lise + Per + Oskar;
	var Louise = math.Log2(8/1)
	fmt.Println("Louse sin informasjon i bits", Louise, ". Louise har 3Bits informasjon og 1 mulig valg" );
}
package oppgaver

import (
	"text/tabwriter"
	"os"
	"fmt"
	"strconv"
	"github.com/mozu0/huffman"
)

// Oppgave 4a
func UIAlist(){
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Fakulitet \t år 2014 \t år 2015 \t")
	fmt.Fprintln(w, "\t\t")
	fmt.Fprintln(w, "Helse og Idrettsdag \t 1420 \t 1829 \t")
	fmt.Fprintln(w, "Humanoira og pedagogikk \t 1182 \t 1525 \t")
	fmt.Fprintln(w, "Kunstfag \t 293 \t 420 \t")
	fmt.Fprintln(w, "Teknologi og realfag \t 1337 \t 2166 \t")
	fmt.Fprintln(w, "Lærerutdanning \t 1158 \t 1506 \t")
	fmt.Fprintln(w, "Økonomi og samfunnsvitenskap \t 2398 \t 3098 \t")
	w.Flush()
	var TOT float64 = 10544
	var HI float64 = 1829
	var HP float64 = 1525
	var K float64 = 420
	var TR float64 = 2166
	var L float64 = 1506
	var OS float64 = 3098
	var hundred float64 = 100

	// Sannsynlighet = Muligheter / total * 100 %
	fmt.Println()
	fmt.Fprintln(w, "Fakulitet \t Sannsynlighet 2015 \t")
	fmt.Fprintln(w, "\t\t")
	fmt.Fprintln(w, "Helse og Idrettsdag \t ", HI/TOT*hundred, "% \t")
	fmt.Fprintln(w, "Humaniora og pedagogikk \t ", HP/TOT*hundred, "% \t")
	fmt.Fprintln(w, "Kunstfag \t", K/TOT*hundred, "% \t")
	fmt.Fprintln(w, "Teknologi og realfag \t ", TR/TOT*hundred, "% \t")
	fmt.Fprintln(w, "Lærerutdanning \t ", L/TOT*hundred, "% \t")
	fmt.Fprintln(w, "Økonomi og samfunnsvitenskap \t ", OS/TOT*hundred, "% \t")
	w.Flush()
}

// Oppgave 4b
func UIAinfo(){
	var HI int64 = 1829
	var HP int64 = 1525
	var K int64 = 420
	var TR int64 = 2166
	var L int64 = 1506
	var OS int64 = 3098
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Fakulitet \t Informasjon \t")
	fmt.Fprintln(w, "\t\t")
	fmt.Fprintln(w, "Helse og Idrettsdag \t", strconv.FormatInt(HI, 2), " \t")
	fmt.Fprintln(w, "Humaniora og pedagogikk \t ", strconv.FormatInt(HP, 2), " \t")
	fmt.Fprintln(w, "Kunstfag \t ", strconv.FormatInt(K, 2), " \t")
	fmt.Fprintln(w, "Teknologi og realfag \t ", strconv.FormatInt(TR, 2), " \t")
	fmt.Fprintln(w, "Lærerutdanning \t ", strconv.FormatInt(L, 2), " \t")
	fmt.Fprintln(w, "Økonomi og samfunnsvitenskap \t", strconv.FormatInt(OS, 2), " \t")
	w.Flush()
}

// Oppgave 4c
// Importerer "github.com/mozu0/huffman"
// Skriver så ut resultatet i table
func UIAhuffman(){
	HI  := 1829
	HP := 1525
	K := 420
	TR := 2166
	L := 1506
	OS := 3098

	//HIb  := 11100100101
	//HPb := 10111110101
	//Kb := 110100100
	//TRb := 100001110110
	//Lb := 10111100010
	//OSb := 110000011010

	all := []int{HI, HP, K, TR, L, OS}
	huffstring := huffman.FromInts(all)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Fakulitet \t Informasjon \t")
	fmt.Fprintln(w, "\t\t")
	fmt.Fprintln(w, "Helse og Idrettsdag \t", huffstring[0], " \t")
	fmt.Fprintln(w, "Humaniora og pedagogikk \t ", huffstring[1], " \t")
	fmt.Fprintln(w, "Kunstfag \t ", huffstring[2], " \t")
	fmt.Fprintln(w, "Teknologi og realfag \t ", huffstring[3], " \t")
	fmt.Fprintln(w, "Lærerutdanning \t ", huffstring[4], " \t")
	fmt.Fprintln(w, "Økonomi og samfunnsvitenskap \t", huffstring[5], " \t")
	w.Flush()

}


func UIAaverage() {

	// Av bit lengden
	fmt.Println("Antall * (Lengden på melding * sansynlighet) + (Lengden på ......")
	fmt.Print("Lengde på melding til 100 stk fra huffman bit lengde: ")
	tott := 100*(3*0.1734) + (3 * 0.1446) + (3 * 0.0398) + (2 * 0.2054) + (3 * 0.1428) + (2 * 0.2938)
	fmt.Print(tott)
	fmt.Print("\nTil binary:")
	fmt.Print(strconv.FormatInt(int64(tott), 2))
	fmt.Println("")
}

//////////////////
// EXPERIMENTAL //
//////////////////

// Av "kodelengde"
//	A := 11100100101
//	B := 10111110101
//	C := 110100100
//	D := 100001110110
//	E := 10111100010
//	F := 110000011010
//	Al := len(string(A))
//	Bl := len(string(B))
//	Cl := len(string(C))
//	Dl := len(string(D))
//	El := len(string(E))
//	Fl := len(string(F))
//	al := float64(Al)
//	bl := float64(Bl)
//	cl := float64(Cl)
//	dl := float64(Dl)
//	el := float64(El)
//	fl := float64(Fl)
//	fmt.Print("Lengde på melding til 100 stk fra bit lengde av string: ")
//	tot := 100*(al*0.1734) + (bl * 0.1446) + (cl * 0.0398) + (dl * 0.2054) + (el * 0.1428) + (fl * 0.2938)
//	fmt.Print(tot)
//	fmt.Print("\nTil binary:")
//	fmt.Println(strconv.FormatInt(int64(tot), 2))
//
//
//	// Test for lengde av faktisk string
//	As := "Helse og Idrettsdag"
//	Bs := "Humanoira og pedagogikk"
//	Cs := "Kunstfag"
//	Ds := "Teknologi og realfag"
//	Es := "Lærerutdanning"
//	Fs := "Økonomi og samfunnsvitenskap"
//	as := float64(len(As))
//	bs := float64(len(Bs))
//	cs := float64(len(Cs))
//	ds := float64(len(Ds))
//	es := float64(len(Es))
//	fs := float64(len(Fs))
//	fmt.Print("Lengde på melding til 100 stk fra String lengde: ")
//	tots := 100*(as*0.1734) + (bs * 0.1446) + (cs * 0.0398) + (ds * 0.2054) + (es * 0.1428) + (fs * 0.2938)
//	fmt.Print(tot)
//	fmt.Print("\nTil binary:")
//	fmt.Println(strconv.FormatInt(int64(tots), 2))
//}
//////////////////////////////////
// EXPERIMENTAL /////////////////
//////////////////////////////////
//As := "Helse og Idrettsdag"
//Bs := "Humanoira og pedagogikk"
//Cs := "Kunstfag"
//Ds := "Teknologi og realfag"
//Es := "Lærerutdanning"
//Fs := "Økonomi og samfunnsvitenskap"
//a1 := stringToBin(As)
//b2 := stringToBin(Bs)
//c3 := stringToBin(Cs)
//d4 := stringToBin(Ds)
//e5 := stringToBin(Es)
//f6 := stringToBin(Fs)
//as, _ := strconv.ParseFloat(a1, 64)
//bs, _ := strconv.ParseFloat(b2, 64)
//cs, _ := strconv.ParseFloat(c3, 64)
//ds, _ := strconv.ParseFloat(d4, 64)
//es, _ := strconv.ParseFloat(e5, 64)
//fs, _ := strconv.ParseFloat(f6, 64)
//fmt.Print("Lengde på melding til 100 stk fra String lengde: ")
//tots := 100*(as*0.1734) + (bs * 0.1446) + (cs * 0.0398) + (ds * 0.2054) + (es * 0.1428) + (fs * 0.2938)
//fmt.Println(strconv.FormatInt(int64(tots), 2))
//
//}
//func stringToBin(s string) (binString string) {
//	for _, c := range s {
//		binString = fmt.Sprintf("%s%b",binString, c)
//	}
//	return
//}
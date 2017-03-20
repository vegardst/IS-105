package oppgaver

//////////////////////////////////
// 	ICA04, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////
import (
	"./huffman"
	"text/tabwriter"
	"os"
	"fmt"
	"strconv"
)

var huffstring []string

func Oppgave4a() {
	/*
	 Helse og Idrettsdag		1420	1829
	 Humanoira og pedagogikk	1182	1525
	 Kunstfag			 293	 420
	 Teknologi og realfag		1337	2166
	 Lærerutdanning			1158	1506
	 Økonomi og samfunnsvitenskap	2398	3098

	 Antall mulige utfall:		7788	10544
	 */
	UIAlist()

}
func Oppgave4b() {
	UIAinfo()
	fmt.Println(" Fakulitetet Kunstfag får minst informasjon som sett over.")

}
func Oppgave4c() {
	UIAhuffman()

}
func Oppgave4d() {
	UIAaverage()
}

func Oppgave4e() {
	EncodeDecore()
}

// Oppgave 4a
func UIAlist() {
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
func UIAinfo() {
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
// Skriver så ut resultatet i table
func UIAhuffman() {
	//HI := 1829
	//HP := 1525
	//K := 420
	//TR := 2166
	//L := 1506
	//OS := 3098

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Fakulitet \t Informasjon \t")
	fmt.Fprintln(w, "\t\t")
	fmt.Fprintln(w, "Helse og Idrettsdag \t", "111", " \t")
	fmt.Fprintln(w, "Humaniora og pedagogikk \t ", "110", " \t")
	fmt.Fprintln(w, "Kunstfag \t ", "000", " \t")
	fmt.Fprintln(w, "Teknologi og realfag \t ", "01", " \t")
	fmt.Fprintln(w, "Lærerutdanning \t ", "001", " \t")
	fmt.Fprintln(w, "Økonomi og samfunnsvitenskap \t", "10", " \t")
	w.Flush()

}

// Oppgave 4d
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

// Oppgave 4e
func EncodeDecore(){
	//	A	         Helse og Idrettsdag 	1829       |        111  |
	//	B	     Humaniora og pedagogikk	 1525      |        110  |
	//	C	                    Kunstfag 	293          |        000  |
	//	D	        Teknologi og realfag	 1337       |         01  |
	//	E	              Lærerutdanning 	1158        |        001  |
	//	F	Økonomi og samfunnsvitenskap 	2398         |         10  |

	huffmancode := make(map[string]int)
	huffmancode["Helse og Idrettsdag "] = 1829
	huffmancode["Humaniora og pedagogikk"] = 1525
	huffmancode["Kunstfag"] = 293
	huffmancode["Teknologi og realfag"] = 1337
	huffmancode["Lærerutdanning"] = 1158
	huffmancode["Økonomi og samfunnsvitenskap"] = 2398

	tree := huffman.BuildTree(huffmancode)

	// print out results
	w := tabwriter.NewWriter(os.Stdout, 0, 100, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Verdi \t Huffman \t Fakulitet \t")
	w.Flush()
	huffman.PrintCodes(tree, []byte{})


}


//////////////////
// EXPERIMENTAL //
//////////////////
//func stringToBin(s string) (binString string) {
//	for _, c := range s {
//		binString = fmt.Sprintf("%s%b",binString, c)
//	}
//	return
//}

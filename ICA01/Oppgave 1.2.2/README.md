<h1># Oppgave 1.2.2</h1><p>
<p><strong>Informasjonsmengde Flere personer pr&oslash;ver &aring; gjette et tresifret (3-bit) bin&aelig;rt tall. </strong></p>
<p><strong>(1) Lise har f&aring;tt vite / l&aelig;rer at tallet er et oddetall.</strong><br /><strong>(2) Per har f&aring;tt vite at tallet er IKKE et multiplum av 3 (dvs. ikke 0, 3, 6). </strong><br /><strong>(3) Oskar har f&aring;tt vite at tallet inneholder n&oslash;yaktig 2 enere. </strong><br /><strong>(4) Louise har f&aring;tt vite alt det Lise, Per og Oskar vet.</strong></p>
<p>3 bit bin&aelig;re tall gir mulighetene:</p>
<table style="height: 154px;" border="1" width="251">
<tbody>
<tr>
<td style="text-align: center;">000</td>
<td style="text-align: center;">0</td>
</tr>
<tr>
<td style="text-align: center;">001</td>
<td style="text-align: center;">1</td>
</tr>
<tr>
<td style="text-align: center;">010</td>
<td style="text-align: center;">2</td>
</tr>
<tr>
<td style="text-align: center;">011</td>
<td style="text-align: center;">3</td>
</tr>
<tr>
<td style="text-align: center;">100</td>
<td style="text-align: center;">4</td>
</tr>
<tr>
<td style="text-align: center;">101</td>
<td style="text-align: center;">5</td>
</tr>
<tr>
<td style="text-align: center;">110</td>
<td style="text-align: center;">6</td>
</tr>
<tr>
<td style="text-align: center;">111</td>
<td style="text-align: center;">7</td>
</tr>
</tbody>
</table>
<p><strong>1.&nbsp;</strong>Lise vet at det er oddetall, og st&aring;r igjen med tallene: 1, 3, 5, 7</p>
<table style="height: 154px;" border="1" width="251">
<tbody>
<tr>
<td style="text-align: center;">001</td>
<td style="text-align: center;">1</td>
</tr>
<tr>
<td style="text-align: center;">011</td>
<td style="text-align: center;">3</td>
</tr>
<tr>
<td style="text-align: center;">101</td>
<td style="text-align: center;">5</td>
</tr>
<tr>
<td style="text-align: center;">111</td>
<td style="text-align: center;">7</td>
</tr>
</tbody>
</table>
<p><strong>2.&nbsp;</strong>Per vet at tallet IKKE er multiplum av 3 (multiplum er et tall som forekommer n&aring;r tallet multipliseres med et helt tall.)<br />3*0 = 0<br />3*1 = 3<br />3*2 = 6<br />Og Per st&aring;r igjen med tallene: 1, 2, 4, 5, 7</p>
<table style="height: 154px;" border="1" width="251">
<tbody>
<tr>
<td style="text-align: center;">001</td>
<td style="text-align: center;">1</td>
</tr>
<tr>
<td style="text-align: center;">010</td>
<td style="text-align: center;">2</td>
</tr>
<tr>
<td style="text-align: center;">100</td>
<td style="text-align: center;">4</td>
</tr>
<tr>
<td style="text-align: center;">101</td>
<td style="text-align: center;">5</td>
</tr>
<tr>
<td style="text-align: center;">111</td>
<td style="text-align: center;">7</td>
</tr>
</tbody>
</table>
<p><strong>3.</strong>&nbsp; Oskar har f&aring;tt vite at tallet inneholder n&oslash;yaktig 2 enere som er tallene 3,5,6<strong><br /></strong></p>
<table style="height: 154px;" border="1" width="251">
<tbody>
<tr>
<td style="text-align: center;">011</td>
<td style="text-align: center;">3</td>
</tr>
<tr>
<td style="text-align: center;">101</td>
<td style="text-align: center;">5</td>
</tr>
<tr>
<td style="text-align: center;">110</td>
<td style="text-align: center;">6</td>
</tr>
</tbody>
</table>
<p>&nbsp;</p>
<p><strong>(4) </strong>Louise har f&aring;tt vite alt det Lise, Per og Oskar vet.&nbsp;</p>
<p>S&aring; Louise vet at tallet er 5 , bin&aelig;rtall 101 </p>
<table style="height: 173px;" border="1" width="303">
<tbody>
<tr>
<td>Lise</td>
<td>Per</td>
<td>Oscar</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;1</td>
<td>1</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>2</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;3</td>
<td>&nbsp;</td>
<td>3</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>4</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;5</td>
<td>5</td>
<td>5</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>6</td>
</tr>
<tr>
<td>&nbsp;7</td>
<td>7</td>
<td>&nbsp;</td>
</tr>
</tbody>
</table>
<p>&nbsp;</p>
<p><strong>Hvor mye informasjon (i bits) har hver spiller f&aring;tt? </strong></p>
<p>Det er totalt 8 muligheter med 3 sifret bin&aelig;re tall: N=8 som gir Log2(8) = 3 bit<br />Formel brukt i oppgave: log2(8/mulige valg)</p>
<p>1 ) Lise vet at tallet er odetall, N=8/4<br />Lise sin informasjon i bits Log2(8/4) = 1 bit. S&aring; Lise har 2 bit igjen &aring; gjette.</p>
<p>2) Per vet at tallet ikke er multiplum av 3, og st&aring;r igjen med 5 tall. N=8/5<br />Per sin informasjon i bits Log2(8/5) = 0 bit. Per f&aring;r under 1 bit informasjon &aring; m&aring; gjette 3 bits ut av 5 mulige valg.</p>
<p>3) Oskar vet at tallet inneholder n&oslash;yaktig 2enere som gir 3 muligheter. N=8/3<br />Oskar sin informasjon i bits Log2(8/3) = 1 bit. Per har like mye informasjon som Lise og har 2 bit igjen &aring; gjette.</p>
4) Louise har fått vite alt Lise,Per og Oskar vet = log2(8/1)<br>
Louse sin informasjon i bits 3 . Louise har 3Bits informasjon og 1 mulig valg
<h1># Oppgave 1.2.1</h1><p>
<p><strong>Konvertering av desimaltall til 2-tallsystemet (bin&aelig;rt tallsystem):</strong></p>
<p>Her skal vi regne ut desimaltall til bin&aelig;re tall, for &aring; forst&aring; hvordan datamaskinen leser disse tallene.<br />Grunnen til at datamaskiner bruker bin&aelig;re tall henger sammen med hvordan digitale kretser er bygd opp (f.eks OG, NOG, NELLER-kretser) med logiske porter. Bin&aelig;re tall blir derfor tolket som 1-Sann , 0-Usann.<br /><br />For &aring; telle bin&aelig;re tall leses disse ,"Arabisk" , fra h&oslash;yre til venstre. Og imotsetning til titallssystemet, blir bin&aelig;re tall hverdier to ganger s&aring; h&oslash;ye for hver siffer plassering til venstre.</p>
<p>Eksempel:&nbsp;</p>
<table style="height: 59px;" border="1" width="311">
<tbody>
<tr>
<td style="text-align: center;">2^7</td>
<td style="text-align: center;">2^6</td>
<td style="text-align: center;">2^5</td>
<td style="text-align: center;">2^4</td>
<td style="text-align: center;">2^3</td>
<td style="text-align: center;">2^2</td>
<td style="text-align: center;">2^1</td>
<td style="text-align: center;">2^0</td>
</tr>
<tr>
<td style="text-align: center;">128</td>
<td style="text-align: center;">64</td>
<td style="text-align: center;">32</td>
<td style="text-align: center;">16</td>
<td style="text-align: center;">8</td>
<td style="text-align: center;">4</td>
<td style="text-align: center;">2</td>
<td style="text-align: center;">1</td>
</tr>
<tr>
<td style="text-align: center;"><strong>1</strong></td>
<td style="text-align: center;"><strong>0</strong></td>
<td style="text-align: center;"><strong>1</strong></td>
<td style="text-align: center;"><strong>0</strong></td>
<td style="text-align: center;"><strong>1</strong></td>
<td style="text-align: center;"><strong>0</strong></td>
<td style="text-align: center;"><strong>1</strong></td>
<td style="text-align: center;"><strong>1</strong></td>
</tr>
</tbody>
</table>
<p>S&aring; for &aring; regne ut dette legger vi sammen verdien for 1 bitene.<br />128 + 32 + 8 + 2 + 1 = 171</p>
<p><br />N&aring; vi skal regne fra desimal til bin&aelig;rt tall som i oppgaven, g&aring;r vi motsatt vei.<br />Dvs at vi begynner med det h&oslash;yest mulig tallet i rekken&nbsp;<br />S&aring; ser vi p&aring; tabellen over, blir 1 = 1<br />Tallet 2 f&aring;r verdien 10<br />Tallet 5, har b&aring;de verdien 4 og 1, og blir derfor 101 (Vi teller fra det h&oslash;yeste tallet)<br /><strong>(1)&nbsp;</strong>1 = 1 (1bit)<br /><strong>(2)</strong>&nbsp;2 = 10 (2bit)<strong><br />(3)</strong>&nbsp;5 = 101 (3bit)<strong>&nbsp;<br />(4)</strong>&nbsp;8 = 1000 (4bit)<strong>&nbsp;<br />(5)</strong>&nbsp;16 = 10000 (5bit)<strong>&nbsp;<br />(6)&nbsp;</strong>256 = 100000000 (9bit)</p>
<table style="height: 62px;" border="1" width="417">
<tbody>
<tr>
<td style="text-align: center;">Desima ltall</td>
<td style="text-align: center;">&nbsp;1</td>
<td style="text-align: center;">&nbsp;2</td>
<td style="text-align: center;">&nbsp;5</td>
<td style="text-align: center;">&nbsp;8</td>
<td style="text-align: center;">&nbsp;16</td>
<td style="text-align: center;">&nbsp;256</td>
</tr>
<tr>
<td style="text-align: center;">&nbsp;Bin&aelig;r tall</td>
<td style="text-align: center;">&nbsp;1</td>
<td style="text-align: center;">&nbsp;10</td>
<td style="text-align: center;">&nbsp;101</td>
<td style="text-align: center;">&nbsp;1000</td>
<td style="text-align: center;">10000&nbsp;</td>
<td style="text-align: center;">100000000</td>
</tr>
</tbody>
</table>
<p>For &aring; regne ut tall der du ikke finner verdien direkte i tallsystemet, slik som tallet 5&nbsp;tar du bin&aelig;r tallet som er nermest, f.eks med tallet 22:<br />22 = 16 + 6 (Men siden 6 ikke finnes i totalssystemet, m&aring; vi g&aring; ned en verdi)<br />22 = 16 + 4 + 2 (Da har vi funnet tallene i totalssystemet, og setter de sammen til bin&aelig;r kode)<br />22 = 2^4 + 2^2 + 2^1&nbsp;<br />22 = 10110<br /><br /></p>
<p>N&aring;r vi jobber med b&aring;de 2talls og 10talls systemet er det vanlig &aring; skrive bak tallene hvilket system de tilh&oslash;rer: 10110 (to)</p>
<p>&nbsp;</p>
<p><strong>1) Konverter f&oslash;lgende bin&aelig;re tall til desimaltall (mest signifikante bit-en er til venstre):</strong></p>
<p>Hvis vi vil konvertere bin&aelig;re tall til tall i titallssystemet kan vi legge sammen&nbsp;plassverdien til alle&nbsp;<strong>1</strong> sifrene.</p>
<p><strong>(1)&nbsp;</strong>100 = 2^2 = 4<br /><strong>(2)&nbsp;</strong>1001 = 2^3 + 2^0 = 9<br /><strong>(3)&nbsp;</strong>1100110011 = 819</p>
<table style="height: 40px;" border="1" width="436">
<tbody>
<tr>
<td style="text-align: center;">2^9</td>
<td style="text-align: center;">2^8</td>
<td style="text-align: center;">2^7</td>
<td style="text-align: center;">2^6</td>
<td style="text-align: center;">2^5</td>
<td style="text-align: center;">2^4</td>
<td style="text-align: center;">2^3</td>
<td style="text-align: center;">2^2</td>
<td style="text-align: center;">2^1</td>
<td style="text-align: center;">2^0</td>
</tr>
<tr>
<td style="text-align: center;">1</td>
<td style="text-align: center;">1</td>
<td style="text-align: center;">0</td>
<td style="text-align: center;">0</td>
<td style="text-align: center;">1</td>
<td style="text-align: center;">1</td>
<td style="text-align: center;">0</td>
<td style="text-align: center;">0</td>
<td style="text-align: center;">1</td>
<td style="text-align: center;">1</td>
</tr>
</tbody>
</table>
<p>= 2^9 + 2^8 + 2^5 + 2^4 + 2^1 + 2^0&nbsp;<br />= 819<br /><br />Resultat:</p>
<table style="height: 40px;" border="1" width="291">
<tbody>
<tr>
<td style="text-align: center;">Bin&aelig;rtall</td>
<td style="text-align: center;">100</td>
<td style="text-align: center;">1001</td>
<td style="text-align: center;">1100110011</td>
</tr>
<tr>
<td style="text-align: center;">&nbsp;Desimaltall</td>
<td style="text-align: center;">&nbsp;4</td>
<td style="text-align: center;">9&nbsp;</td>
<td style="text-align: center;">&nbsp;819</td>
</tr>
</tbody>
</table>
<p>&nbsp;</p>
<p>&nbsp;</p>
<p><em>Informasjonskilde: Tek.no</em></p>

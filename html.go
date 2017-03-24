package main

const (
	html = `<HTML>
<HEAD>
	<TITLE></TITLE>
	<STYLE TYPE="text/css">
	<!--
		@page { margin: 2cm }
		P { margin-bottom: 0.21cm }
		TD P { margin-bottom: 0cm }
	-->
	</STYLE>
</HEAD>
<BODY LANG="en-IN" DIR="LTR">
<TABLE WIDTH=100% BORDER=1 BORDERCOLOR="#000000" CELLPADDING=4 CELLSPACING=0>
	<COL WIDTH=256*>
	<TR>
		<TD WIDTH=100% VALIGN=TOP>
			<P STYLE="background: SPRINGGREEN">I'm running on pod {{.Pod}} with version {{.Version}} at {{.TS}} </P>
		</TD>
	</TR>
</TABLE>
<P STYLE="margin-bottom: 0cm"><BR>
</P>
</BODY>
</HTML>`
)


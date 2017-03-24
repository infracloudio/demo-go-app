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
		<TD WIDTH=50% VALIGN=TOP>
			<P STYLE="background: SPRINGGREEN"><b>Pod ID</b> </P>
		</TD>
                <TD WIDTH=50% VALIGN=TOP>
                        <P STYLE="background: SPRINGGREEN">{{.Pod}}</P>
                </TD>
	</TR>
	<TR>
	        <TD WIDTH=50% VALIGN=TOP>
                        <P STYLE="background: SPRINGGREEN"><b>Version</b> </P>
                </TD>
                <TD WIDTH=50% VALIGN=TOP>
                        <P STYLE="background: SPRINGGREEN">{{.Version}}</P>
                </TD>
	</TR>
	<TR>
                <TD WIDTH=50% VALIGN=TOP>
                        <P STYLE="background: SPRINGGREEN"><b>Timestamp</b> </P>
                </TD>
                <TD WIDTH=50% VALIGN=TOP>
                        <P STYLE="background: SPRINGGREEN">{{.TS}}</P>
                </TD>
	</TR>
</TABLE>
<P STYLE="margin-bottom: 0cm"><BR>
</P>
</BODY>
</HTML>`
)


package main

import "html/template"

var mainTmpl = template.Must(template.New("index").Parse(mainTmplStr))

const mainTmplStr = `
<html>
<head>
<title>overtime tracker</title>
</head>
<body>
<h2>Time tracker</h2>
<h3>Last session</h3>
<form action={{if .S.Active}}"/stop"{{else}}"/start"{{end}}>
	Running:<br>
	<input type="text" name="active" disabled value="{{.S.Active}}"><br>
	Start time:<br>
	<input type="text" name="start time" disabled value="{{.S.Start}}" ><br>
	{{if not .S.Active}}
	End time:<br>
	<input type="text" name="end time" disabled value="{{.S.End}}"><br>
	{{end}}
	Current time:<br>
	<input type="text" name="end time" disabled value="{{.S.D}}"><br>
	<input type="submit" value={{if .S.Active}}"Stop"{{else}}"Start"{{end}}>
</form>

<h3>Past results</h3>
<a href="/download"> download </a>
<br>
<style>
	table, th, td {
	  border: 1px solid black;
	  border-collapse: collapse;
	}
	th, td {
	  padding: 5px;
	}
</style>
<table>
	<tr>
		<th>Start</th>
		<th>End</th>
		<th>Duration</th>
		<th>Duration, hours</th>
	</tr>

	{{ with .Ss }}
		{{ range . }}
			<tr>
				<td>{{ .Start }}</td>
				<td>{{ .End }}</td>
				<td>{{ .D }}</td>
				<td>{{ .Hours }}</td>
			</tr>
		{{ end }}
	{{ end }}

</body>
</html>
`

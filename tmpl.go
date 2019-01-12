package main

import "html/template"

var mainTmpl = template.Must(template.New("index").Parse(mainTmplStr))

const mainTmplStr = `
<html lang="en">

<head>
	<title>overtime tracker</title>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">
</head>

<body>
<div class="container">
	<h2>Time tracker</h2>
	<h3>Last session</h3>
	<form class="col-md-6" action={{if .S.Active}}"/stop"{{else}}"/start"{{end}}>
		<div class="form-group">
			<label for="active">Active</label>
			<input class="form-control" type="text" name="active" disabled value="{{.S.Active}}">
		</div>
		<div class="form-group">
			<lable for="startTime">Start time</label>
			<input class="form-control" type="text" name="startTime" disabled value="{{.S.Start}}">
		</div>
		{{if not .S.Active}}
		<div class="form-group">
			<lable for="endTime">Start time</label>
			<input class="form-control" type="text" name="endTime" disabled value="{{.S.End}}">
		</div>
		{{end}}
		<div class="form-group">
			<lable for="d">Duration</label>
			<input class="form-control" type="text" name="d" disabled value="{{.S.D}}">
		</div>
		<button type="submit"{{if .S.Active}} class="btn btn-warning">Stop{{else}} class="btn btn-primary">Start{{end}}</button>
	</form>

	<h3>Past results</h3>
	<a href="/download"> download </a>
	<br>
	<table class="table table-sm">
		<thead class="thead-light">
			<tr>
				<th scope="col">Duration, hours</th>
				<th scope="col">Start</th>
				<th scope="col">End</th>
				<th scope="col">Duration</th>
			</tr>
		</thead>

		<tbody>
		{{ with .Ss }}
			{{ range . }}
				<tr>
					<td>{{ .Hours }}</td>
					<td>{{ .Start }}</td>
					<td>{{ .End }}</td>
					<td>{{ .D }}</td>
				</tr>
			{{ end }}
		{{ end }}
		</tbody>
</div>

<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.6/umd/popper.min.js" integrity="sha384-wHAiFfRlMFy6i5SRaxvfOCifBUQy1xHdJ/yoi7FRNXMRBu5WHdZYu1hA6ZOblgut" crossorigin="anonymous"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/js/bootstrap.min.js" integrity="sha384-B0UglyR+jN6CkvvICOB2joaf5I4l3gm9GU6Hc1og6Ls7i6U/mkkaduKaBhlAXv9k" crossorigin="anonymous"></script>

</body>
</html>
`

package main

const tplHome = `
<!DOCTYPE html>
<html>
	<body>
		<ul>
			{{range .Items}}<li>{{.Key}} - {{.Value}}</li>
			{{else}}<li>No Keys</li>
			{{end}}
		</ul>
	</body>
</html>
`

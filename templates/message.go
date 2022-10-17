package templates

var RaspMessage = `
{{ range .Segments -}}
{{- if eq .Thread.Transport_subtype.Title "«Ласточка»" }}
Поезд 			:	{{ .Thread.Transport_subtype.Title }}
Отправление в	:	{{ .Departure }}
От 				:	{{ .From.Title }}
До				:	{{ .To.Title }}
Время в пути 	:	{{ Hours .Duration }}
{{ end -}}
{{ end -}}
`

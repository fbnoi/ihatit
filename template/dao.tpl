// Code generated by ihatit. DO NOT EDIT.
// source: {{ .source }}

package db

{{ range .mapping }}
func {{ .name }}({{range .args}}{{.name}} {{type}},{{end}}) {
    
}

{{ end }}

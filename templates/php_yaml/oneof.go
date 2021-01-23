package php_yaml

const oneOfConstTpl = `
{{ range .Fields }}{{ renderConstants (context .) }}{{ end }}
`

const oneOfTpl = `
      - TODOOneOf: ~
`

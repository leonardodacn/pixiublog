{{.config.siteUrl}}{{range .blogTypes}}{{$.config.siteUrl}}/type/{{.Id}}
{{end}}{{range .blogTags}}{{$.config.siteUrl}}/tag/{{.Id}}
{{end}}{{range .blogs}}{{$.config.siteUrl}}/blog/{{.Code}}
{{end}}
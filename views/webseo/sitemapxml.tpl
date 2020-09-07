<urlset xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
	<url>
		<loc>{{.config.siteUrl}}</loc>
		<priority>1.0</priority>
		<lastmod>{{.currentDate}}</lastmod>
		<changefreq>daily</changefreq>
	</url>
	{{range .blogs}}
		<url>
			<loc>{{$.config.siteUrl}}/blog/{{.Code}}</loc>
			<priority>0.6</priority>
			<lastmod>{{DateStr .CreateTime}}</lastmod>
			<changefreq>daily</changefreq>
		</url>
	{{end}}
	{{range .blogTypes}}
	   <url>
			<loc>{{$.config.siteUrl}}/type/{{.Id}}</loc>
			<priority>0.6</priority>
			<lastmod>{{DateStr .CreateTime}}</lastmod>
			<changefreq>daily</changefreq>
		</url>
    {{end}}
    
	{{range .blogTags}}
	   <url>
			<loc>{{$.config.siteUrl}}/tag/{{.Id}}</loc>
			<priority>0.6</priority>
			<lastmod>{{DateStr .CreateTime}}</lastmod>
			<changefreq>daily</changefreq>
		</url>
	{{end}}
</urlset>
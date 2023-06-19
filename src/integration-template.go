package src

var IntegrationPageTemplate = `<h2>{{ int.name }}</h2>


<ul>
   {% for item in int.dashboards %}
      <li>
	  	{{ item.title }}
		<p>{{item.description}}</p>
	  </li>
   {% endfor %}
</ul>
`

func NewIntegrationTemplate(integration string) string {
	template := "---\n\n---\n\n"
	template += "{% assign int = site.data.integration_" + integration + " %}\n\n"
	template += IntegrationPageTemplate

	return template
}

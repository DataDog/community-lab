package src

import "fmt"

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
	template := fmt.Sprintf("---\ncategory: integration\ntitle: %s\n---\n\n", integration)
	template += "{% assign int = site.data.integration_" + integration + " %}\n\n"
	template += IntegrationPageTemplate

	return template
}

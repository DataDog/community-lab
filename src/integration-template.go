package src

import "fmt"

var IntegrationPageTemplate = `
<div class="integration-title">
	<h1>{{ int.name }}</h1>
	<div class="line"></div>
</div>


<div class="dash-list">
   {% for item in int.dashboards %}
      <div class="dash">
	  <figure>
	  	{% if item.image != "" %}
	  		<img src="https://github.com/{{site.repository}}/blob/main/{{item.image}}?raw=true" alt="{{ item.title }}" >
		{% else %}
			<img src="{{ site.baseurl }}/assets/dashboard.png" alt="{{ item.title }}" >
		{% endif %}
			<figcaption><a href="https://github.com/{{site.repository}}/blob/main/{{item.path}}">{{ item.title }}</a></figcaption>
		</figure>
		<div class="middle">
    		<div class="text">{{item.description}}</div>
  		</div>
	  </div>
   {% endfor %}
</div>
`

func NewIntegrationTemplate(integration string) string {
	template := fmt.Sprintf("---\ncategory: integration\ntitle: %s\n---\n\n", integration)
	template += "{% assign int = site.data.integration_" + integration + " %}\n\n"
	template += IntegrationPageTemplate

	return template
}

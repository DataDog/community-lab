---

---

{% assign int = site.data.integration_azure %}

<h2>{{ int.name }}</h2>


<ul>
   {% for item in int.dashboards %}
      <li>
	  	{{ item.title }}
		<p>{{item.description}}</p>
	  </li>
   {% endfor %}
</ul>

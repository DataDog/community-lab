---
category: integration
title: {{ .Name }}
---

<div class="integration-title">
	<h1>{{ .Name }}</h1>
	<div class="line"></div>
</div>

<div>
	<input id="tab1" type="radio" name="tabs" checked>
	<label for="tab1">Dashboards</label>
	<input id="tab2" type="radio" name="tabs">
	<label for="tab2">Monitors</label>
	<section id="content1">
		<div class="dash-list">
			{{ range .Dashboards }}
			<div class="dash">
				<figure>
					{{ if .Image }}
					<img src="https://github.com/{{`{{site.repository}}`}}/blob/main/{{ .Image }}?raw=true"
						alt="{{ .Title }}">
					{{ else }}
					<img src="{{`{{ site.baseurl }}`}}/assets/dashboard.png" alt="{{ .Title }}">
					{{ end }}
					<figcaption><a href="https://github.com/{{`{{site.repository}}`}}/blob/main/{{ .Path }}">{{ .Title
							}}</a>
					</figcaption>
				</figure>
				<div class="middle">
					<div class="text">{{ .Description }}</div>
				</div>
			</div>
			{{ end }}
		</div>
	</section>
	<section id="content2">
		<ul class="mon-list">
			{{ range .Monitors }}
			<li class="mon">
				<a href="https://github.com/{{`{{site.repository}}`}}/blob/main/{{ .Path }}">{{ .Title }}</a>
				<p>{{ .Description }}</p>
			</li>
			{{ end }}
		</ul>
	</section>
</div>
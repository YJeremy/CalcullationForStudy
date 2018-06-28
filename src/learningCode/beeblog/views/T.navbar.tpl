{{define "navbar"}}
	<a class="navbar-brand" href="/">我的测试页</a>
	<div>
	<u1 class="nav navbar-nav">
	  <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
	  <li {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
	  <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a><li>
	</u1>
	</div>


<div class="pull-right">
     <ul class="nav navbar-nav">
     {{if .IsLogin}}
     <li><a href="/login?exit=123">退出</a></li>
     {{else}}
     <li><a href="/login? exit=321">管理员登录</a></li>
     {{end}}
     </ul>
</div>

{{end}}
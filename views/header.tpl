<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="/static/css/style.css">
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css">
     {{ block "css" . }}{{ end }}
</head>

<body>
    
<div id="container">
    <nav>
        <ul>
            <li><a href="#">Master</a>
            <!-- First Tier Drop Down -->
            <ul>
                <li><a href="/company">Company</a></li>
            </ul>        
            </li>
			<li><a href="#">Billing</a>
            <!-- First Tier Drop Down -->
            <ul>
                <li><a href="/bill">Company's Bill</a></li>
            </ul>        
            </li>

        </ul>
    </nav>
</div>
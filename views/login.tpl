<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Sign in - Beego Billing</title>
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
<div class="app-loader" aria-hidden="true"></div>
<div class="container" style="max-width:460px; margin-top:80px;">
    <h1>Beego Billing</h1>
    <p>Sign in to continue.</p>
    {{if .Error}}<div class="alert alert-danger">{{.Error}}</div>{{end}}
    <form method="post" action="/login">
        <div class="form-group">
            <label for="username">Username</label>
            <input class="form-control" id="username" name="username" required autofocus>
        </div>
        <div class="form-group">
            <label for="password">Password</label>
            <input class="form-control" id="password" name="password" type="password" required>
        </div>
        <button class="btn btn-primary" type="submit"><span class="spinner" aria-hidden="true">&#9679;</span>Sign in</button>
    </form>
</div>
<script>
document.querySelector('form').addEventListener('submit', function () {
    this.classList.add('is-submitting');
    this.querySelector('button').classList.add('is-loading');
    document.querySelector('.app-loader').classList.add('is-loading');
});
</script>
</body>
</html>

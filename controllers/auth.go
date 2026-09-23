package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	models "Learning-Beego/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	beego.Controller
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth *AuthController) Login() {
	if auth.Ctx.Input.Method() == http.MethodGet {
		auth.TplName = "login.tpl"
		return
	}

	var input credentials
	if isJSONRequest(auth.Ctx.Input.Header("Content-Type")) && len(auth.Ctx.Input.RequestBody) > 0 {
		if err := json.Unmarshal(auth.Ctx.Input.RequestBody, &input); err != nil {
			auth.loginError(http.StatusBadRequest, "Invalid JSON request")
			return
		}
	} else {
		input.Username = auth.GetString("username")
		input.Password = auth.GetString("password")
	}
	if input.Username == "" || input.Password == "" {
		auth.loginError(http.StatusBadRequest, "Username and password are required")
		return
	}

	var user models.User
	o := orm.NewOrm()
	err := o.QueryTable("users").Filter("Username", input.Username).One(&user)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)) != nil {
		auth.loginError(http.StatusUnauthorized, "Invalid username or password")
		return
	}

	auth.SetSession("authenticated", true)
	auth.SetSession("user_id", user.Id)
	auth.SetSession("username", user.Username)
	if strings.HasPrefix(auth.Ctx.Input.Header("Content-Type"), "application/json") {
		auth.Data["json"] = map[string]interface{}{"success": true, "message": "Successfully logged in", "username": user.Username}
		auth.ServeJSON()
		return
	}
	auth.Redirect("/company", http.StatusFound)
}

func (auth *AuthController) Logout() {
	auth.DestroySession()
	if strings.HasPrefix(auth.Ctx.Input.Header("Accept"), "application/json") {
		auth.Data["json"] = map[string]interface{}{"success": true, "message": "Successfully signed out"}
		auth.ServeJSON()
		return
	}
	auth.Redirect("/login", http.StatusFound)
}

func (auth *AuthController) User_Save() {
	var input credentials
	if isJSONRequest(auth.Ctx.Input.Header("Content-Type")) && len(auth.Ctx.Input.RequestBody) > 0 {
		if err := json.Unmarshal(auth.Ctx.Input.RequestBody, &input); err != nil {
			auth.writeJSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "Invalid JSON request"})
			return
		}
	} else {
		input.Username = auth.GetString("username")
		input.Password = auth.GetString("password")
	}
	if len(input.Username) < 3 || len(input.Password) < 8 {
		auth.writeJSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "Username must have 3 characters and password must have 8 characters"})
		return
	}

	o := orm.NewOrm()
	var existing models.User
	if err := o.QueryTable("users").Filter("Username", input.Username).One(&existing); err == nil {
		auth.writeJSON(http.StatusConflict, map[string]interface{}{"success": false, "message": "Username already exists"})
		return
	} else if err != orm.ErrNoRows {
		auth.writeJSON(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": "Failed to check username"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		auth.writeJSON(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": "Failed to create user"})
		return
	}
	if _, err = o.Insert(&models.User{Username: input.Username, PasswordHash: string(hash)}); err != nil {
		auth.writeJSON(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": "Failed to create user"})
		return
	}
	auth.writeJSON(http.StatusCreated, map[string]interface{}{"success": true, "message": "Successfully created user", "username": input.Username})
}

func (auth *AuthController) CurrentUser() {
	auth.Data["json"] = map[string]interface{}{"success": true, "username": auth.GetSession("username")}
	auth.ServeJSON()
}

func (auth *AuthController) loginError(status int, message string) {
	if strings.HasPrefix(auth.Ctx.Input.Header("Content-Type"), "application/json") {
		auth.writeJSON(status, map[string]interface{}{"success": false, "message": message})
		return
	}
	auth.Ctx.ResponseWriter.WriteHeader(status)
	auth.Data["Error"] = message
	auth.TplName = "login.tpl"
}

func (auth *AuthController) writeJSON(status int, data map[string]interface{}) {
	auth.Ctx.ResponseWriter.WriteHeader(status)
	auth.Data["json"] = data
	auth.ServeJSON()
}

func isJSONRequest(contentType string) bool {
	return strings.HasPrefix(strings.ToLower(contentType), "application/json")
}

func AuthFilter(ctx *context.Context) {
	path := ctx.Input.URL()
	method := ctx.Input.Method()
	if path == "/login" || path == "/logout" || (path == "/user" && method == http.MethodPost) || strings.HasPrefix(path, "/static/") {
		return
	}
	if ctx.Input.CruSession != nil && ctx.Input.CruSession.Get("authenticated") == true {
		return
	}
	if strings.HasPrefix(ctx.Input.Header("Accept"), "application/json") || strings.HasPrefix(ctx.Input.Header("Content-Type"), "application/json") {
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		_ = ctx.Output.JSON(map[string]interface{}{"success": false, "message": "Authentication required"}, false, false)
		return
	}
	ctx.Redirect(http.StatusFound, "/login")
}

package controllers

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/revel/revel"

	"github.com/fasilz/co-memo/app/models"
)

type User struct {
	App
}

func (u User) AddUser() revel.Result {
	if user := u.connected(); user != nil {
		u.ViewArgs["user"] = user
	}
	return nil
}

func (u User) connected() *models.User {
	if u.ViewArgs["user"] != nil {
		return u.ViewArgs["user"].(*models.User)
	}
	if username, ok := u.Session["user"]; ok {
		return u.getUser(username.(string))
	}
	return nil
}

func (u User) getUser(username string) (user *models.User) {
	user = &models.User{}
	_, err := u.Session.GetInto("fulluser", user, false)
	if user.Username == username {
		return user
	}

	err = u.Txn.First(&user, "Username=?", username).Error
	if err != nil {
		revel.AppLog.Info("failed to retrieve user info from db")
		return nil
	}

	u.Session["fulluser"] = user
	return
}

func (u User) Index() revel.Result {
	if u.connected() != nil {
		return u.Redirect(u.Index)
	}
	u.Flash.Error("Please log in first")
	return u.Render()
}

func (u User) Register() revel.Result {
	return u.Render()
}

func (u User) SaveUser(user models.User, verifyPassword string) revel.Result {
	u.Validation.Required(verifyPassword)
	u.Validation.Required(verifyPassword == user.Password).
		MessageKey("Password does not match")
	user.Validate(u.Validation)

	if u.Validation.HasErrors() {
		u.Validation.Keep()
		u.FlashParams()
		return u.Redirect(User.Register)
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	err := u.Txn.Create(&user)
	if err != nil {
		revel.AppLog.Info("failed to create a user")
		return nil
	}

	u.Session["user"] = user.Username
	u.Flash.Success("Welcome, " + user.Name)
	return u.Redirect(u.Index)
}

func (u User) Login(username, password string, remember bool) revel.Result {
	user := u.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			u.Session["user"] = username
			if remember {
				u.Session.SetDefaultExpiration()
			} else {
				u.Session.SetNoExpiration()
			}
			u.Flash.Success("Welcome, " + username)
			return u.Redirect(u.Index)
		}
	}

	u.Flash.Out["username"] = username
	u.Flash.Error("Login failed")
	return u.Redirect(User.Index)
}

func (u User) Logout() revel.Result {
	for k := range u.Session {
		delete(u.Session, k)
	}
	return u.Redirect(User.Index)
}
func (u User) About() revel.Result {
	u.ViewArgs["Msg"] = "Revel Speaks"
	return u.Render()
}

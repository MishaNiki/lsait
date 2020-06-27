package auth

import (
	"encoding/json"
	"net/http"

	"github.com/MishaNiki/lsait/backend/internal/app/mail"
	"github.com/MishaNiki/lsait/backend/internal/app/model"
	"github.com/MishaNiki/lsait/backend/internal/app/security"
	"github.com/MishaNiki/lsait/backend/internal/app/servers"
	"github.com/MishaNiki/lsait/backend/internal/app/storage"
	"github.com/MishaNiki/lsait/backend/internal/app/storage/combatstorage"
	"github.com/MishaNiki/lsait/backend/internal/app/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	server   *http.Server
	logger   *logrus.Logger
	mail     *mail.Mail
	security *security.Security
	storage  storage.Storage
	frontURL string
}

type Config struct {
	BindAddr string                `json:"bindAddr"`
	LogLevel string                `json:"logLevel"`
	Mail     *mail.Config          `json:"mail"`
	Security *security.Config      `json:"security"`
	Storage  *combatstorage.Config `json:"storage"`
	FrontURL string                `json:"frontURL"`
}

// New ...
func New() *Auth {
	return &Auth{}
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		Mail:     &mail.Config{},
		Security: &security.Config{},
		Storage:  &combatstorage.Config{},
	}
}

// Configure ...
func (a *Auth) Configure(config *Config) error {

	var err error
	a.logger, err = a.configureLogger(config.LogLevel)
	if err != nil {
		return err
	}

	a.mail, err = mail.New(config.Mail)
	if err != nil {
		return err
	}

	if err = a.mail.Ping(); err != nil {
		return err
	}

	a.security = security.New(config.Security)

	stor := combatstorage.New()
	err = stor.Configure(config.Storage)
	if err != nil {
		return err
	}
	a.storage = stor
	a.frontURL = config.FrontURL

	a.server = &http.Server{
		Addr:    config.BindAddr,
		Handler: a.configureRouter(),
	}
	return nil
}

// Start ...
func (a *Auth) Start() error {
	a.logger.Info("start auth server on port ", a.server.Addr)
	return a.server.ListenAndServe()
}

// configureRouter ...
func (a *Auth) configureRouter() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/auth/signup", a.handleSignup()).Methods("POST")
	router.HandleFunc("/auth/signup/confirm", a.handleSignupConfirm()).Methods("POST")
	router.HandleFunc("/auth/forgot", a.handleForgot()).Methods("POST")
	router.HandleFunc("/auth/forgot/confirm", a.handleForgotConfirm()).Methods("POST")
	router.HandleFunc("/auth/login", a.handleLogin()).Methods("POST")
	router.HandleFunc("/auth/logout", a.handleLogout()).Methods("POST")
	router.HandleFunc("/auth/refrash", a.handleRefrash()).Methods("POST")

	// CORS
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	method := handlers.AllowedMethods([]string{"POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(header, method, origins)(router)

}

// configureLogger ...
func (a *Auth) configureLogger(logLevel string) (*logrus.Logger, error) {
	logger := logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)
	return logger, nil
}

// handleSignup ...
func (a *Auth) handleSignup() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleSignup")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}

		modAuth := &model.Auth{Email: req.Email, Password: req.Password}
		if err := modAuth.Validate(); err != nil {
			servers.ResponseError(w, r, http.StatusUnauthorized, err)
			return
		}
		if err := modAuth.BeforeCreate(); err != nil {
			servers.ResponseError(w, r, http.StatusUnauthorized, err)
			return
		}
		if err := a.storage.Auth().Create(modAuth); err != nil {
			servers.ResponseError(w, r, http.StatusUnauthorized, err)
			return
		}
		key := utils.RandString(24)
		if err := a.storage.ConfirmKey().Create(key, modAuth.ID); err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}

		if err := a.mail.SendingConfirmURL(
			modAuth.Email,
			a.frontURL+"/confirm?key="+key,
		); err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}

		servers.Response(w, r, http.StatusOK, nil)
	}
}

// handleSignupConfirm ...
func (a *Auth) handleSignupConfirm() http.HandlerFunc {

	type request struct {
		Key string `json:"key"`
	}

	type response struct {
		AccessToken  string `json:"accessToken"`
		RefrashToken string `json:"refrashToken"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleSignupConfirm")

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}

		id, ok := a.storage.ConfirmKey().Get(req.Key)
		if !ok {
			servers.Response(w, r, http.StatusUnauthorized, nil)
			return
		}
		modAuth, err := a.storage.Auth().GetByID(id)
		if err != nil {
			servers.ResponseError(w, r, http.StatusUnauthorized, err)
			return
		}

		res := &response{}
		if t, err := a.security.NewTokens(modAuth.ID); err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		} else {
			if err := a.storage.Token().Create(t.RefrashToken, modAuth.ID); err != nil {
				servers.ResponseError(w, r, http.StatusNotAcceptable, err)
				return
			}
			res.AccessToken = t.AccessToken
			res.RefrashToken = t.RefrashToken
		}
		servers.Response(w, r, http.StatusCreated, res)
	}
}

// handleForgot ...
func (a *Auth) handleForgot() http.HandlerFunc {

	type request struct {
		Email string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleForgot")
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}
		modAuth, err := a.storage.Auth().GetByEmail(req.Email)
		if err != nil {
			servers.ResponseError(w, r, http.StatusUnauthorized, err)
			return
		}
		key := utils.RandString(24)
		if err := a.storage.ConfirmKey().Create(key, modAuth.ID); err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}
		if err := a.mail.SendingForgotURL(
			modAuth.Email,
			a.frontURL+"/forgot?key="+key,
		); err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}
		servers.Response(w, r, http.StatusOK, nil)
	}
}

// handleForgotConfirm ...
func (a *Auth) handleForgotConfirm() http.HandlerFunc {

	type request struct {
		Key      string `json:"key"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleForgotConfirm")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}

		id, ok := a.storage.ConfirmKey().Get(req.Key)
		if !ok {
			servers.Response(w, r, http.StatusUnauthorized, nil)
			return
		}
		modAuth, err := a.storage.Auth().GetByID(id)
		if err != nil {
			servers.ResponseError(w, r, http.StatusUnauthorized, err)
			return
		}
		modAuth.EncryptedPassword = ""
		modAuth.Password = req.Password
		if err := modAuth.Validate(); err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}
		if err := modAuth.BeforeCreate(); err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}
		a.storage.Auth().Update(modAuth)

		servers.Response(w, r, http.StatusOK, nil)
	}
}

// handleLogin ...
func (a *Auth) handleLogin() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		AccessToken  string `json:"accessToken"`
		RefrashToken string `json:"refrashToken"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleLogin")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}

		modAuth, err := a.storage.Auth().GetByEmail(req.Email)
		if err != nil {
			servers.ResponseError(w, r, http.StatusUnauthorized, err)
			return
		}
		if err := modAuth.ComparePassword(req.Password); err != nil {
			servers.ResponseError(w, r, http.StatusUnauthorized, err)
			return
		}
		res := &response{}
		if t, err := a.security.NewTokens(modAuth.ID); err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		} else {
			if err := a.storage.Token().Create(t.RefrashToken, modAuth.ID); err != nil {
				servers.ResponseError(w, r, http.StatusNotAcceptable, err)
				return
			}
			res.AccessToken = t.AccessToken
			res.RefrashToken = t.RefrashToken
		}
		servers.Response(w, r, http.StatusOK, res)
	}
}

// handleLogout ...
func (a *Auth) handleLogout() http.HandlerFunc {

	type request struct {
		RefrashToken string `json:"refrashToken"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleLogout")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}
		a.storage.Token().Delete(req.RefrashToken)
		servers.Response(w, r, http.StatusOK, nil)
	}
}

// handleRefrash ...
func (a *Auth) handleRefrash() http.HandlerFunc {

	type request struct {
		RefrashToken string `json:"refrashToken"`
	}

	type response struct {
		AccessToken  string `json:"accessToken"`
		RefrashToken string `json:"refrashToken"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleRefrash")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}
		id, ok := a.storage.Token().Get(req.RefrashToken)
		if !ok {
			servers.Response(w, r, http.StatusUnauthorized, nil)
			return
		}
		res := &response{}
		a.storage.Token().Delete(req.RefrashToken)

		if t, err := a.security.NewTokens(id); err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		} else {
			if err := a.storage.Token().Create(t.RefrashToken, id); err != nil {
				servers.ResponseError(w, r, http.StatusNotAcceptable, err)
				return
			}
			res.AccessToken = t.AccessToken
			res.RefrashToken = t.RefrashToken
		}
		servers.Response(w, r, http.StatusOK, res)
	}
}

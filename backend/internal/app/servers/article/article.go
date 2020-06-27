package article

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MishaNiki/lsait/backend/internal/app/model"
	"github.com/MishaNiki/lsait/backend/internal/app/servers"
	"github.com/MishaNiki/lsait/backend/internal/app/storage/combatstorage"
	"github.com/MishaNiki/lsait/backend/internal/app/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Article ...
type Article struct {
	server    *http.Server
	logger    *logrus.Logger
	secretKey []byte
	storage   *combatstorage.CombatStorage
}

// Config ...
type Config struct {
	BindAddr  string                `json:"bindAddr"`
	LogLevel  string                `json:"logLevel"`
	SecretKey string                `json:"secretKey"`
	Storage   *combatstorage.Config `json:"storage"`
}

func New() *Article {
	return &Article{}
}

func NewConfig() *Config {
	return &Config{
		Storage: &combatstorage.Config{},
	}
}

// Configure ...
func (a *Article) Configure(config *Config) error {

	var err error
	a.logger, err = a.configureLogger(config.LogLevel)
	if err != nil {
		return err
	}

	stor := combatstorage.New()
	err = stor.Configure(config.Storage)
	if err != nil {
		return err
	}
	a.storage = stor
	a.secretKey = []byte(config.SecretKey)

	a.server = &http.Server{
		Addr:    config.BindAddr,
		Handler: a.configureRouter(),
	}
	return nil
}

// Start ...
func (a *Article) Start() error {
	a.logger.Info("start auth server on port ", a.server.Addr)
	return a.server.ListenAndServe()
}

// configureRouter ...
func (a *Article) configureRouter() http.Handler {

	mainRouter := mux.NewRouter()

	publicRouter := mainRouter.PathPrefix("/article").Subrouter()
	privateRouter := mainRouter.PathPrefix("/article").Subrouter()

	// public router
	publicRouter.HandleFunc("/section", a.handleSectionsGET()).Methods("GET")
	publicRouter.HandleFunc("/section/theme", a.handleSectionsAndThemeGET()).Methods("GET")
	publicRouter.HandleFunc("/section/{uuid}", a.handleSectionGET()).Methods("GET")
	publicRouter.HandleFunc("/article/{id}", a.handleArticleGET()).Methods("GET")
	publicRouter.HandleFunc("/profile/{id}", nil).Methods("GET")

	// Security
	privateRouter.HandleFunc("/article", a.handleArticlePOST()).Methods("POST")
	privateRouter.HandleFunc("/article", a.handleArticlePUT()).Methods("PUT")
	privateRouter.HandleFunc("/article", a.handleArticleDELETE()).Methods("DELETE")
	privateRouter.HandleFunc("/edit/{id}", a.handleEditGET()).Methods("GET")
	privateRouter.HandleFunc("/draft", a.handleDraftPOST()).Methods("POST")
	/*
		Для создания статьи сначала создаётся черновик [POST]/draft
		Потом [POST]|[PUT]/artile
		Для получения статьи на редактирование [GET]/edit/{id}
	*/
	//privateRouter.HandleFunc("/draft", a.handleDraftPUT()).Methods("PUT")
	//privateRouter.HandleFunc("/draft", a.handleDraftDELETE()).Methods("DELETE")

	privateRouter.HandleFunc("/profile", a.handleFullProfileGET()).Methods("GET")
	privateRouter.HandleFunc("/profile", a.handleProfilePUT()).Methods("PUT")

	// connecting middleware and combining routers into one
	privateRouter.Use(servers.SecurityMiddleware(a.secretKey))
	mainRouter.Handle("/article/", publicRouter)
	mainRouter.Handle("/article/", privateRouter)

	// CORS
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Token"})
	method := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(header, method, origins)(mainRouter)
}

// configureLogger ...
func (a *Article) configureLogger(logLevel string) (*logrus.Logger, error) {
	logger := logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)
	return logger, nil
}

// HANDLERS
// PUBLIC HANDLERS
// handleSectionsGET ...
func (a *Article) handleSectionsGET() http.HandlerFunc {

	type response struct {
		Sections []*model.Section `json:"sections"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleSectionsGET")

		sections, err := a.storage.Section().GetAll()
		if err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}
		res := response{Sections: sections}
		servers.Response(w, r, http.StatusOK, res)
	}
}

// handleSectionsAndThemeGET
func (a *Article) handleSectionsAndThemeGET() http.HandlerFunc {
	type response struct {
		Sections []*model.Section `json:"sections"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleSectionsAndThemeGET")
		sections, err := a.storage.Section().GetSectionsAndThemes()
		if err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}
		res := response{Sections: sections}
		servers.Response(w, r, http.StatusOK, res)
	}
}

// handleSectionGET ...
func (a *Article) handleSectionGET() http.HandlerFunc {
	type response struct {
		Section *model.Section `json:"section"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleSectionGET")
		vars := mux.Vars(r)
		uuid := vars["uuid"]
		section, err := a.storage.Section().GetByUUID(uuid)
		if err != nil {
			servers.ResponseError(w, r, http.StatusNotFound, err)
			return
		}
		res := &response{Section: section}
		servers.Response(w, r, http.StatusOK, res)
	}
}

// handleArticleGET ...
func (a *Article) handleArticleGET() http.HandlerFunc {

	type response struct {
		Article *model.Article `json:"article"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleArticleGET")
		vars := mux.Vars(r)
		param := vars["id"]
		id, err := strconv.Atoi(param)
		if err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}
		article, err := a.storage.Article().GetByID(id)
		if err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}
		res := &response{Article: article}
		servers.Response(w, r, http.StatusOK, res)
	}
}

// PRIVATE HANDLERS
// handleArticlePOST ...
func (a *Article) handleArticlePOST() http.HandlerFunc {

	type request struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Text        string `json:"text"`
		Theme       int    `json:"idtheme"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleArticlePOST")

		// Checking context
		idUser := r.Context().Value(servers.ContexKey("id"))
		if idUser == nil {
			servers.Response(w, r, http.StatusUnauthorized, nil)
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}
		art := &model.Article{
			ID:          req.ID,
			Title:       req.Title,
			Description: req.Description,
			Text:        req.Text,
			Theme:       req.Theme,
			Auth:        &model.Profile{ID: idUser.(int)},
		}
		a.storage.Article().Create(art)
		servers.Response(w, r, http.StatusCreated, nil)
	}
}

// handleArticlePUT ...
func (a *Article) handleArticlePUT() http.HandlerFunc {

	type request struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Text        string `json:"text"`
		Theme       int    `json:"idtheme"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleArtiarticleclePUT")
		idUser := r.Context().Value(servers.ContexKey("id"))
		if idUser == nil {
			servers.Response(w, r, http.StatusUnauthorized, nil)
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}

		art := &model.Article{
			ID:          req.ID,
			Title:       req.Title,
			Description: req.Description,
			Text:        req.Text,
			Theme:       utils.Ternary(req.Theme == 0, 1, req.Theme).(int),
			Auth:        &model.Profile{ID: idUser.(int)},
		}
		a.storage.Article().Update(art)
		servers.Response(w, r, http.StatusOK, nil)
	}
}

// handleArticlePUT ...
func (a *Article) handleArticleDELETE() http.HandlerFunc {

	type request struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleArticleDELETE")

		idUser := r.Context().Value(servers.ContexKey("id"))
		if idUser == nil {
			servers.Response(w, r, http.StatusUnauthorized, nil)
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}
		a.storage.Article().DeleteByID(req.ID, idUser.(int))
		servers.Response(w, r, http.StatusOK, nil)
	}
}

// handleEditGET ...
func (a *Article) handleEditGET() http.HandlerFunc {

	type response struct {
		Article *model.Article `json:"article"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleEditGET")

		idUser := r.Context().Value(servers.ContexKey("id"))
		if idUser == nil {
			servers.Response(w, r, http.StatusUnauthorized, nil)
			return
		}
		vars := mux.Vars(r)
		param := vars["id"]
		idDraft, err := strconv.Atoi(param)
		if err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}
		article, err := a.storage.Article().GetEditByID(idDraft, idUser.(int))
		if err != nil {
			servers.ResponseError(w, r, http.StatusNotFound, err)
			return
		}
		res := &response{Article: article}
		servers.Response(w, r, http.StatusOK, res)
	}
}

// handleDraftPOST ...
func (a *Article) handleDraftPOST() http.HandlerFunc {

	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	type response struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleDraftPOST")
		idUser := r.Context().Value(servers.ContexKey("id"))
		if idUser == nil {
			servers.Response(w, r, http.StatusUnauthorized, nil)
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}
		draft := &model.Article{
			Title:       req.Title,
			Description: req.Description,
			Auth:        &model.Profile{ID: idUser.(int)},
		}
		err := a.storage.Article().CreateDraft(draft)
		if err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}
		res := &response{ID: draft.ID}
		servers.Response(w, r, http.StatusCreated, res)
	}
}

// handleDraftPUT ...
// func (a *Article) handleDraftPUT() http.HandlerFunc {

// 	type request struct {
// 		ID    int    `json:"id"`
// 		Title string `json:"title"`
// 		Text  string `json:"text"`
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		a.logger.Info("handleDraftPUT")

// 		idUser := r.Context().Value(servers.ContexKey("id"))
// 		if idUser == nil {
// 			servers.Response(w, r, http.StatusUnauthorized, nil)
// 			return
// 		}

// 		req := &request{}
// 		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
// 			servers.ResponseError(w, r, http.StatusBadRequest, err)
// 			return
// 		}
// 		draft := &model.Article{
// 			Title: req.Title,
// 			Text:  req.Text,
// 			Auth:  &model.Profile{ID: idUser.(int)},
// 		}
// 		a.storage.Draft().Update(draft)
// 		servers.Response(w, r, http.StatusOK, nil)
// 	}
// }

// handleDraftDELETE ...
// func (a *Article) handleDraftDELETE() http.HandlerFunc {

// 	type request struct {
// 		ID int `json:"id"`
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		a.logger.Info("handleDraftDELETE")
// 		idUser := r.Context().Value(servers.ContexKey("id"))
// 		if idUser == nil {
// 			servers.Response(w, r, http.StatusUnauthorized, nil)
// 			return
// 		}
// 		req := &request{}
// 		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
// 			servers.ResponseError(w, r, http.StatusBadRequest, err)
// 			return
// 		}
// 		a.storage.Draft().DeleteByID(req.ID, idUser.(int))
// 		servers.Response(w, r, http.StatusOK, nil)
// 	}
// }

func (a *Article) handleFullProfileGET() http.HandlerFunc {

	type response struct {
		Profile *model.Profile `json:"profile"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleFullProfileGET")
		idUser := r.Context().Value(servers.ContexKey("id"))
		if idUser == nil {
			servers.Response(w, r, http.StatusUnauthorized, nil)
			return
		}
		var err error
		res := &response{}
		res.Profile, err = a.storage.Profile().GetFullByID(idUser.(int))
		if err != nil {
			servers.ResponseError(w, r, http.StatusNotAcceptable, err)
			return
		}
		servers.Response(w, r, http.StatusOK, res)
	}
}

// ...
func (a *Article) handleProfilePUT() http.HandlerFunc {

	type request struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		Position     string `json:"position,omitempty"`
		Descriptions string `json:"description,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("handleProfilePUT")
		idUser := r.Context().Value(servers.ContexKey("id"))
		if idUser == nil {
			servers.Response(w, r, http.StatusUnauthorized, nil)
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			servers.ResponseError(w, r, http.StatusBadRequest, err)
			return
		}

		profile := &model.Profile{
			ID:          idUser.(int),
			Name:        req.Name,
			Surname:     req.Surname,
			Position:    req.Position,
			Description: req.Descriptions,
		}
		a.storage.Profile().Update(profile)
		servers.Response(w, r, http.StatusOK, nil)
	}
}

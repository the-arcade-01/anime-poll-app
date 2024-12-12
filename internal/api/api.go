package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/the-arcade-01/anime-poll-app/internal/service"
)

type Server struct {
	Router *chi.Mux
}

func (s *Server) mountMiddlewares() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Heartbeat("/ping"))
	s.Router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

func (s *Server) mountHandlers() {
	apiService := service.NewApiService()
	s.Router.Get("/greet", apiService.Greet)
	s.Router.Post("/db/ingestion", apiService.StartDBAnimeIngestion)
	s.Router.Delete("/db/flush", apiService.FlushAnimeDB)
	s.Router.Delete("/db/{id}", apiService.DeleteAnimeById)
	s.Router.Get("/db/animes", apiService.FetchAllAnimes)
	s.Router.Get("/anime/fight", apiService.GetAnimesForFaceOff)
	s.Router.Post("/anime/vote/{id}", apiService.VoteAnime)
	s.Router.Get("/anime/leaderboard/{count}", apiService.GetLeaderBoard)
}

func NewServer() *Server {
	server := &Server{
		Router: chi.NewRouter(),
	}
	server.mountMiddlewares()
	server.mountHandlers()
	return server
}

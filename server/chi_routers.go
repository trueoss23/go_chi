package server

h := handlers.NewHandler(usecase)

R := chi.NewRouter()
R.Use(middleware.Logger)
R.Get("/books", h.GetAllBooks)
R.Post("/book", h.CreateBook)
R.Delete("/book/{id}", h.DeleteBook)
R.Get("/book/{id}", h.GetBook)
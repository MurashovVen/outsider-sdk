package app

type Option func(*App)

func AppendWorks(works ...Work) Option {
	return func(app *App) {
		app.works = append(app.works, works...)
	}
}

package app

type Option func(*App)

func AppendWorks(works ...Work) Option {
	return func(app *App) {
		for _, work := range works {
			app.works = append(app.works, work)
		}
	}
}

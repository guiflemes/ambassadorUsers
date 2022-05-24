package config

type Config struct {
	App      App
	Database Database
}

func Parser() Config {
	app := &App{}
	db := &Database{}

	app.Parse()
	db.Parse()

	config := Config{
		App:      *app,
		Database: *db,
	}

	return config

}

package database

func StartDB() {
	Connect()
	AutoMigrate()
}

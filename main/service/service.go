package service

func InitServices() error {
	// Initialize PostgreSQL
	if err := InitPostgres(); err != nil {
		return err
	}

	// Run database migrations
	if err := RunMigrations(); err != nil {
		return err
	}

	// Initialize Redis
	if err := InitRedis(); err != nil {
		return err
	}

	return nil
}

func CloseServices() {
	ClosePostgres()
	CloseRedis()

}

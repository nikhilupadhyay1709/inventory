package database

func Prepare() {
	Init()
	Migrate()
}

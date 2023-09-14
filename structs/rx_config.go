package structs

type ReindexerConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Namespaces Namespaces
}

type Namespaces struct {
	Teachers string
	Groups   string
	Names    string
}

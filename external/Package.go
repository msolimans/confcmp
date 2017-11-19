package external

//added here as a seprate package
type Package struct {
	Authors []Author // treated like a set
	Kind  string     // case insensitive from the set "app", "lib", "builder"
	Version  [3]int
	VersionString  string // if present, split on '.' and use as Version
}

type Author struct {
	Name, Email string
}

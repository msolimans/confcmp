package config2

//kept in config2, not moved to contracts package (AS IT IS ONLY USED OR INTENDED FOR config2 package)
type IManager interface{
	Compare(p1 []byte, p2 []byte) ( error, bool)
	Diff(p1 []byte, p2 []byte) ( error, string)
}

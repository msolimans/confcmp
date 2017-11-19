package contracts

type IManager interface{
	Compare(p1 []byte, p2 []byte) ( error, bool)
	Diff(p1 []byte, p2 []byte) ( error, string)
}

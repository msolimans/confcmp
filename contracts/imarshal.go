package contracts

//contract or behaviour for marshaling algorithms we may use
type IMarshal interface{


	Compare(p1 []byte, p2 []byte) ( error, bool)//can be separated to different behaviour (to say for example IComparable)
}

package contracts

//contract or behaviour for marshaling algorithms we may use
type IConfig interface{
	Unmarshal(in []byte) (chan error)
	Marshal() (out []byte, err error)
}

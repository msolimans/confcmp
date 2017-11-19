package contracts

//contract or behaviour for marshaling algorithms we may use
type IConfig interface{
	UpdateVersion(done chan bool)
	Unmarshal(in []byte, ch chan error)
	Marshal() (out []byte, err error)
}

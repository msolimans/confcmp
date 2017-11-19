package contracts

//contract or behaviour for marshaling algorithms we may use
type IMarshal interface{


	Compare() ( error, bool)//can be separated to different behaviour (to say for example IComparable)
}

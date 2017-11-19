package contracts

type IManager interface{
	Compare() ( error, bool)
	Diff() ( error, string)
}

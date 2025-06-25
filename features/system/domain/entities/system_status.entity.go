package entities

type SystemStatus struct {
	OS      string
	CPU     CPU
	Memory  Memory
	GPU     GPU
	Storage Storage
	Server  Server
}

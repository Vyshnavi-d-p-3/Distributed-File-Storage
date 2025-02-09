package p2p

// Message represents any arbitrary date that is being sent over each transport between two node in the network

type Message struct {
	Payload []byte
}

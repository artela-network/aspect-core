package api

func wrapNilByte(resp []byte) []byte {
	if resp == nil {
		return []byte{}
	}
	return resp

}

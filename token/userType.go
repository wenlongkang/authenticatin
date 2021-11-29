package token

type UserType int

const (
	//IM 不托管用户信息，但需要管理设备信息
	IM UserType = iota
	//simple 不托管用户信息，也不需要管理设备信息
	SIMPLE
	//trust 托管用户信息，并需要管理设备信息
	TRUST
)

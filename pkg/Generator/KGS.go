package generator

const (
	char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	base = 62
)

func Create(num int) string {
	var key []byte
	for num > 0 {
		key = append(key, char[num%base])
		num /= base
	}
	return string(key)
}

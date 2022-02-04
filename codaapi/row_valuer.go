package codaapi

func (row *Row) GetValue(key string) (value interface{}, ok bool) {
	return row.Values.Get(key)
}

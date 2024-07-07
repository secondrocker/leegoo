package data

type KvData struct {
	data *Data
}

func NewKvData(data *Data) *KvData {
	return &KvData{data: data}
}

func (kv *KvData) SetKv(key string, data string) error {
	return kv.data.LevelDb.Put([]byte(key), []byte(data), nil)
}

func (kv *KvData) GetKv(key string) (data string, err error) {
	var val []byte
	val, err = kv.data.LevelDb.Get([]byte(key), nil)
	if err == nil {
		data = string(val)
	}
	return data, err
}

func (kv *KvData) DelKv(key string) error {
	return kv.data.LevelDb.Delete([]byte(key), nil)
}

package cloud

type CloudDB struct {
	Url string
}

func NewCloudDB(url string) *CloudDB {
	return &CloudDB{
		Url: url,
	}
}

func (db *CloudDB) Write(content []byte) {}

func (db *CloudDB) Read() ([]byte, error) {
	return []byte{}, nil
}

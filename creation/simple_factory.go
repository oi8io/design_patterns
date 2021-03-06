package creation

/**
工厂模式
1. 抽象工厂方法
2. 分别实现工厂方法
3. 采用工厂方法创建实例
*/

type SimpleFactoryStorage interface {
	Create(data interface{})
	Read() interface{}
	Update(data interface{})
	Delete(data interface{})
}

type SimpleFactoryMySQLFactory struct {
	data interface{}
}

func (m *SimpleFactoryMySQLFactory) Create(data interface{}) {
	m.data = data
}

func (m *SimpleFactoryMySQLFactory) Read() interface{} {
	return m.data
}

func (m *SimpleFactoryMySQLFactory) Update(data interface{}) {
	m.data = data
}

func (m *SimpleFactoryMySQLFactory) Delete(data interface{}) {
	m.data = nil
}

type SimpleFactoryMgoFactory struct {
}

func (m SimpleFactoryMgoFactory) Create(data interface{}) {
	panic("implement me")
}

func (m SimpleFactoryMgoFactory) Read() interface{} {
	panic("implement me")
}

func (m SimpleFactoryMgoFactory) Update(data interface{}) {
	panic("implement me")
}

func (m SimpleFactoryMgoFactory) Delete(data interface{}) {
	panic("implement me")
}

const (
	StorageTypeMySQL   = "MySQL"
	StorageTypeRedis   = "Redis"
	StorageTypeMongoDB = "Mgo"
)

func NewSimpleFactoryStorage(storageType string) SimpleFactoryStorage {
	switch storageType {
	case StorageTypeMongoDB:
		return new(SimpleFactoryMgoFactory)
	case StorageTypeMySQL:
		return new(SimpleFactoryMySQLFactory)
	case StorageTypeRedis:
		return new(SimpleFactoryMySQLFactory)
	}
	return new(SimpleFactoryMySQLFactory)
}

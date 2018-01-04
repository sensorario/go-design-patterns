package flyweight

const (
	TYPE_ONE = iota
	TYPE_TWO
)

type Object struct {
	ID   uint64
	Name int
}

type objectFactory struct {
	createdObjects map[int]*Object
}

func (f *objectFactory) GetObject(objectID int) *Object {
	if f.createdObjects[objectID] != nil {
		return f.createdObjects[objectID]
	}
	object := getObjectFactory(objectID)
	f.createdObjects[objectID] = &object
	return f.createdObjects[objectID]
}

func (f *objectFactory) GetNumberOfObjects() int {
	return len(f.createdObjects)
}

func getObjectFactory(object int) Object {
	switch object {
	case TYPE_TWO:
		return Object{
			ID:   2,
			Name: TYPE_TWO,
		}
	default:
		return Object{
			ID:   1,
			Name: TYPE_ONE,
		}
	}
}

func NewObjectFactory() objectFactory {
	return objectFactory{
		createdObjects: make(map[int]*Object),
	}
}

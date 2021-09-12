package common

type entityMap map[Entity]Entity

type EntityManager struct {
	Entities     entityMap
	removalQueue entityMap
}

func (em *EntityManager) IsInit() bool {
	return em.Entities != nil && em.removalQueue != nil
}

func (em *EntityManager) Init() {
	em.Entities = make(entityMap)
	em.removalQueue = make(entityMap)
}

func (em *EntityManager) AddEntity(e Entity) {
	if !em.IsInit() {
		em.Init()
	}

	em.Entities[e] = e
}

func (em *EntityManager) RemoveEntity(e Entity) {
	em.removalQueue[e] = e
}

func (em *EntityManager) ProcessRemovalQueue() {
	for e := range em.removalQueue {
		delete(em.removalQueue, e)
		delete(em.Entities, e)
	}
}

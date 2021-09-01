package common

import "github.com/gravestench/akara"

type entityMap map[akara.EID]akara.EID

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

func (em *EntityManager) AddEntity(e akara.EID) {
	if !em.IsInit() {
		em.Init()
	}

	em.Entities[e] = e
}

func (em *EntityManager) RemoveEntity(e akara.EID) {
	em.removalQueue[e] = e
}

func (em *EntityManager) ProcessRemovalQueue() {
	for e := range em.removalQueue {
		delete(em.removalQueue, e)
		delete(em.Entities, e)
	}
}

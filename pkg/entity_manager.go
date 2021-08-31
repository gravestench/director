package pkg

import "github.com/gravestench/akara"

type entityMap map[akara.EID]akara.EID

type entityManager struct {
	entities     entityMap
	removalQueue entityMap
}

func (em *entityManager) entityManagerIsInit() bool {
	return em.entities != nil && em.removalQueue != nil
}

func (em *entityManager) entityManagerInit() {
	em.entities = make(entityMap)
	em.removalQueue = make(entityMap)
}

func (em *entityManager) addEntity(e akara.EID) {
	if !em.entityManagerIsInit() {
		em.entityManagerInit()
	}

	em.entities[e] = e
}

func (em *entityManager) removeEntity(e akara.EID) {
	em.removalQueue[e] = e
}

func (em *entityManager) processRemovalQueue() {
	for e := range em.removalQueue {
		delete(em.removalQueue, e)
		delete(em.entities, e)
	}
}

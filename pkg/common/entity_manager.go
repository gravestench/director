package common

import (
	"sync"

	"github.com/gravestench/akara"
)

type entityMap map[akara.EID]akara.EID

type EntityManager struct {
	EntitiesMutex     sync.Mutex
	Entities          entityMap
	removalQueueMutex sync.Mutex
	removalQueue      entityMap
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

	em.EntitiesMutex.Lock()
	em.Entities[e] = e
	em.EntitiesMutex.Unlock()
}

func (em *EntityManager) RemoveEntity(e akara.EID) {
	em.removalQueueMutex.Lock()
	em.removalQueue[e] = e
	em.removalQueueMutex.Unlock()
}

func (em *EntityManager) ProcessRemovalQueue() {
	em.EntitiesMutex.Lock()
	em.removalQueueMutex.Lock()

	for e := range em.removalQueue {
		delete(em.removalQueue, e)
		delete(em.Entities, e)
	}

	em.EntitiesMutex.Unlock()
	em.removalQueueMutex.Unlock()
}

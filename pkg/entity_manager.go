package pkg

import "github.com/gravestench/akara"

type entityManager struct {
	entities []akara.EID
	removalQueue []akara.EID
}

func (em *entityManager) entityManagerIsInit() bool {
	return em.entities == nil || em.removalQueue == nil
}

func (em *entityManager) entityManagerInit() {
	em.entities = make([]akara.EID, 0)
	em.removalQueue = make([]akara.EID, 0)
}

func (em *entityManager) processRemovalQueue() {

}


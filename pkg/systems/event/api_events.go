package event

type ApiEvent = int

const (
	OnBlindfold ApiEvent = iota
	OnChat
	OnExternalMessage
	OnFixedUpdate
	OnLoad
	OnObjectCollisionEnter
	OnObjectCollisionExit
	OnObjectCollisionStay
	OnObjectDestroy
	OnObjectDrop
	OnObjectEnterContainer
	OnObjectEnterScriptingZone
	OnObjectEnterZone
	OnObjectFlick
	OnObjectHover
	OnObjectLeaveContainer
	OnObjectLeaveScriptingZone
	OnObjectLeaveZone
	OnObjectLoopingEffect
	OnObjectNumberTyped
	OnObjectPageChange
	OnObjectPeek
	OnObjectPickUp
	OnObjectRandomize
	OnObjectRotate
	OnObjectSearchEnd
	OnObjectSearchStart
	OnObjectSpawn
	OnObjectStateChange
	OnObjectTriggerEffect
	OnPlayerAction
	OnPlayerChangeColor
	OnPlayerChangeTeam
	OnPlayerConnect
	OnPlayerDisconnect
	OnPlayerPing
	OnPlayerTurn
	OnSave
	OnScriptingButtonDown
	OnScriptingButtonUp
	OnUpdate
)
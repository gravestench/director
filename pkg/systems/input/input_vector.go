package input

import (
	"github.com/gravestench/bitset"
)

// NewInputVector creates a new input vector
func NewInputVector() *Vector {
	v := &Vector{
		KeyVector:         bitset.NewBitSet(),
		ModifierVector:    bitset.NewBitSet(),
		MouseButtonVector: bitset.NewBitSet(),
	}

	return v.Clear()
}

// Vector represents the state of keys, modifiers, and mouse buttons.
// It can be used to compare input states, and is intended to be used as such:
// 		* whatever manages system input keeps a "current" input vector and updates it
//  	* things that are listening for certain inputs will be compared using `Contains` and `Intersects` methods
type Vector struct {
	KeyVector         *bitset.BitSet
	ModifierVector    *bitset.BitSet
	MouseButtonVector *bitset.BitSet
}

// SetKey sets the corresponding key bit in the keys bitset
func (iv *Vector) SetKey(key Key) *Vector {
	return iv.SetKeys([]Key{key})
}

// SetKeys sets multiple key bits in the keys bitset
func (iv *Vector) SetKeys(keys []Key) *Vector {
	if len(keys) == 0 {
		return iv
	}

	for _, key := range keys {
		iv.KeyVector.Set(int(key), true)
	}

	return iv
}

// SetModifier sets the corresponding modifier bit in the modifier bitset
func (iv *Vector) SetModifier(mod Modifier) *Vector {
	return iv.SetModifiers([]Modifier{mod})
}

// SetModifiers sets multiple modifier bits in the modifier bitset
func (iv *Vector) SetModifiers(mods []Modifier) *Vector {
	if len(mods) == 0 {
		return iv
	}

	for _, key := range mods {
		iv.ModifierVector.Set(int(key), true)
	}

	return iv
}

// SetMouseButton sets the corresponding mouse button bit in the mouse button bitset
func (iv *Vector) SetMouseButton(button MouseButton) *Vector {
	return iv.SetMouseButtons([]MouseButton{button})
}

// SetMouseButtons sets multiple mouse button bits in the mouse button bitset
func (iv *Vector) SetMouseButtons(buttons []MouseButton) *Vector {
	if len(buttons) == 0 {
		return iv
	}

	for _, key := range buttons {
		iv.MouseButtonVector.Set(int(key), true)
	}

	return iv
}

// Contains returns true if this input vector is a superset of the given input vector
func (iv *Vector) Contains(other *Vector) bool {
	keys := iv.KeyVector.ContainsAll(other.KeyVector)

	buttons := iv.MouseButtonVector.ContainsAll(other.MouseButtonVector)

	// We do Equals here, because we dont want CTRL+X and CTRL+ALT+X to fire at the same time
	mods := iv.ModifierVector.Equals(other.ModifierVector)

	return keys && mods && buttons
}

// Intersects returns true if this input vector shares any bits with the given input vector
func (iv *Vector) Intersects(other *Vector) bool {
	keys := iv.KeyVector.Intersects(other.KeyVector)
	mods := iv.ModifierVector.Intersects(other.ModifierVector)
	buttons := iv.MouseButtonVector.Intersects(other.MouseButtonVector)

	return keys || mods || buttons
}

// Clear sets all bits in this input vector to 0
func (iv *Vector) Clear() *Vector {
	iv.KeyVector.Clear()
	iv.ModifierVector.Clear()
	iv.MouseButtonVector.Clear()

	return iv
}

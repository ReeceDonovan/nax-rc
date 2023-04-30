package graph

// GenericSet is a simple set implementation that can store elements of any type.
type GenericSet map[interface{}]interface{}

// HashableElement is an interface that represents any type that can be hashed.
type HashableElement interface {
	Hashcode() interface{}
}

// GetElementHashcode returns the hashcode of an element if it implements the HashableElement interface.
// Otherwise, it returns the element itself.
func GetElementHashcode(element interface{}) interface{} {
	// If the element implements the HashableElement interface, use that to get a hashcode.
	if hashableElement, ok := element.(HashableElement); ok {
		return hashableElement.Hashcode()
	}
	// Otherwise, just use the element itself as the hashcode.
	return element
}

// Contains checks if the set contains a specific element.
func (set GenericSet) Contains(element interface{}) bool {
	elementHash := GetElementHashcode(element)
	_, ok := set[elementHash]
	return ok
}

// Add inserts an element into the set.
func (set GenericSet) Add(element interface{}) {
	hashcode := GetElementHashcode(element)
	set[hashcode] = element
}

// Delete removes an element from the set.
func (set GenericSet) Delete(element interface{}) {
	hashcode := GetElementHashcode(element)
	delete(set, hashcode)
}

// Intersection returns a new set containing the elements that are common to both sets.
func (set GenericSet) Intersection(otherSet GenericSet) GenericSet {
	intersection := make(GenericSet)
	if set == nil || otherSet == nil {
		return intersection
	}
	// For efficiency, iterate over the smaller set.
	if otherSet.Length() < set.Length() {
		set, otherSet = otherSet, set
	}
	// Add the elements of the first set to the intersection if they are also in the second set
	for _, element := range set {
		if otherSet.Contains(element) {
			intersection.Add(element)
		}
	}
	return intersection
}

// Difference returns a new set containing the elements that are present in the current set but not in the other set.
func (set GenericSet) Difference(otherSet GenericSet) GenericSet {
	// If the other set is nil or empty, then return a copy of this set.
	if otherSet == nil || otherSet.Length() == 0 {
		return set.Copy()
	}
	// Otherwise, return a set with the elements of this set that are not in the other set.
	difference := make(GenericSet)
	for key, element := range set {
		if _, ok := otherSet[key]; !ok {
			difference.Add(element)
		}
	}
	return difference
}

// Filter returns a new set containing the elements that satisfy the provided filter function.
func (set GenericSet) Filter(filterFn func(interface{}) bool) GenericSet {
	filteredSet := make(GenericSet)
	// Add the elements that satisfy the filter function to the new set.
	for element := range set {
		if filterFn(element) {
			filteredSet.Add(element)
		}
	}
	return filteredSet
}

// Copy creates a new set with the same elements as the current set.
func (set GenericSet) Copy() GenericSet {
	copiedSet := make(GenericSet, len(set))
	// Copy each element from the original set into the new set
	for key, element := range set {
		copiedSet[key] = element
	}
	return copiedSet
}

// Length returns the number of elements in the set.
func (set GenericSet) Length() int {
	return len(set)
}

// ToList converts the set to a list of elements.
func (set GenericSet) List() []interface{} {
	if set == nil {
		return nil
	}
	// Convert this set to a list, preserving order.
	list := make([]interface{}, 0, len(set))
	for _, element := range set {
		list = append(list, element)
	}
	return list
}

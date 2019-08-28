package protolock

import (
	"reflect"
)

// Check whether one lockfile is equal to another.
func (p *Protolock) Equal(q *Protolock) bool {
	// Check whether the two lockfiles have the same list of
	// definitions, ignoring order.
	return isPermutation(p.Definitions, q.Definitions, equalDefinitions)
}

// Check whether two slices are equal, ignoring ordering.
// Uses the provided comparator function to determine equality.
func isPermutation(as, bs interface{}, cmp func(x, y interface{}) bool) bool {
	aKind := reflect.TypeOf(as).Kind()
	bKind := reflect.TypeOf(bs).Kind()
	if aKind != reflect.Array && aKind != reflect.Slice {
		panic("isPermutation was given an argument that isn't an array or slice")
	}
	if bKind != reflect.Array && bKind != reflect.Slice {
		panic("isPermutation was given an argument that isn't an array or slice")
	}

	// Get the lengths of the slices via reflection
	aList := reflect.ValueOf(as)
	bList := reflect.ValueOf(bs)
	aLen := aList.Len()
	bLen := bList.Len()

	// Slices of different lengths are trivially inequal
	if aLen != bLen {
		return false
	}

	// Empty slices are trivially equal
	if aLen == 0 {
		return true
	}

	// Try to match each element in A to an element in B
	// Keep track of which elements in B we've already matched
	used := make([]bool, bLen)
	for i := 0; i < aLen; i++ {
		current := aList.Index(i).Interface()
		found := false
		for j := 0; j < bLen; j++ {
			if used[j] {
				continue
			}
			candidate := bList.Index(j).Interface()

			if cmp(current, candidate) {
				// Found a match, mark it as used
				found = true
				used[j] = true
				break
			}
		}

		if !found {
			// Nothing in B (that was not already matched)
			// matches the current element, slices are
			// inequal
			return false
		}
	}

	return true
}

// Helper functions to determine equality of subparts of a lockfile.
// Some functions take interface{} because they're used with
// isPermutation.

func equalDefinitions(i, j interface{}) bool {
	a := i.(Definition)
	b := j.(Definition)
	return a.Filepath == b.Filepath && equalEntries(a.Def, b.Def)
}

func equalEntries(a, b Entry) bool {
	if a.Package != b.Package {
		return false
	}
	if !isPermutation(a.Enums, b.Enums, equalEnums) {
		return false
	}
	if !isPermutation(a.Messages, b.Messages, equalMessages) {
		return false
	}
	if !isPermutation(a.Services, b.Services, equalServices) {
		return false
	}
	if !isPermutation(a.Imports, b.Imports, equalImports) {
		return false
	}
	return isPermutation(a.Options, b.Options, equalOptions)
}

func equalImports(i, j interface{}) bool {
	// Struct has only primitive fields and no slice fields, fall
	// back to default equality
	a := i.(Import)
	b := j.(Import)
	return a == b
}

func equalPackage(i, j interface{}) bool {
	// Struct has only primitive fields and no slice fields, fall
	// back to default equality
	a := i.(Package)
	b := j.(Package)
	return a == b
}

func equalOptions(i, j interface{}) bool {
	a := i.(Option)
	b := j.(Option)

	if a.Name != b.Name || a.Value != b.Value {
		return false
	}
	return isPermutation(a.Aggregated, b.Aggregated, equalOptions)
}

func equalMessages(i, j interface{}) bool {
	a := i.(Message)
	b := j.(Message)

	if a.Name != b.Name || a.Filepath != b.Filepath {
		return false
	}
	if !isPermutation(a.Fields, b.Fields, equalFields) {
		return false
	}
	if !isPermutation(a.Maps, b.Maps, equalMaps) {
		return false
	}
	if !isPermutation(a.ReservedIDs, b.ReservedIDs, equalPrimitives) {
		return false
	}
	if !isPermutation(a.ReservedNames, b.ReservedNames, equalPrimitives) {
		return false
	}
	if !isPermutation(a.Messages, b.Messages, equalMessages) {
		return false
	}
	return isPermutation(a.Options, b.Options, equalOptions)
}

func equalEnumFields(i, j interface{}) bool {
	a := i.(EnumField)
	b := j.(EnumField)

	if a.Name != b.Name || a.Integer != b.Integer {
		return false
	}
	return isPermutation(a.Options, b.Options, equalOptions)
}

func equalEnums(i, j interface{}) bool {
	a := i.(Enum)
	b := j.(Enum)

	if a.Name != b.Name || a.AllowAlias != b.AllowAlias {
		return false
	}
	if !isPermutation(a.ReservedIDs, b.ReservedIDs, equalPrimitives) {
		return false
	}
	if !isPermutation(a.ReservedNames, b.ReservedNames, equalPrimitives) {
		return false
	}
	return isPermutation(a.EnumFields, b.EnumFields, equalEnumFields)
}

func equalMaps(i, j interface{}) bool {
	a := i.(Map)
	b := j.(Map)

	return a.KeyType == b.KeyType && equalFields(a.Field, b.Field)
}

func equalFields(i, j interface{}) bool {
	a := i.(Field)
	b := j.(Field)

	if a.ID != b.ID || a.Name != b.Name {
		return false
	}
	if a.Type != b.Type || a.IsRepeated != b.IsRepeated {
		return false
	}
	return isPermutation(a.Options, b.Options, equalOptions)
}

func equalServices(i, j interface{}) bool {
	a := i.(Service)
	b := j.(Service)

	if a.Name != b.Name || a.Filepath != b.Filepath {
		return false
	}
	return isPermutation(a.RPCs, b.RPCs, equalRPCs)
}

func equalRPCs(i, j interface{}) bool {
	a := i.(RPC)
	b := j.(RPC)

	if a.Name != b.Name || a.InType != b.InType || a.OutType != b.OutType {
		return false
	}
	if a.InStreamed != b.InStreamed || a.OutStreamed != b.OutStreamed {
		return false
	}

	return isPermutation(a.Options, b.Options, equalOptions)
}

// Helper to compare primitive types in isPermutation
func equalPrimitives(i, j interface{}) bool {
	return i == j
}

package app

// Access represents app Access
type Access string

// List of accesses
const (
	AccessRead   Access = "r"
	AccessWrite  Access = "w"
	AccessEdit   Access = "e"
	AccessDelete Access = "d"
)

// AccessList provides list of Accesses
type AccessList []Access

// AllAccesses returns access list with all access
func AllAccesses() AccessList {
	return AccessList{AccessRead, AccessWrite, AccessEdit, AccessDelete}
}

// Has looks for specifice Access
func (a AccessList) Has(ac Access) bool {
	for _, v := range a {
		if v == ac {
			return true
		}
	}

	return false
}

// CanRead checks read access
func (a AccessList) CanRead() bool {
	return a.Has(AccessRead)
}

// CanWrite checks write access
func (a AccessList) CanWrite() bool {
	return a.Has(AccessRead)
}

// CanEdit checks edit access
func (a AccessList) CanEdit() bool {
	return a.Has(AccessEdit)
}

// CanDelete checks delete access
func (a AccessList) CanDelete() bool {
	return a.Has(AccessDelete)
}

// Remove removes specific Access
func (a *AccessList) Remove(ac Access) {
	acList := *a

	for i, v := range *a {
		if v == ac {
			acList = append(acList[:i], acList[i+1:]...)
			*a = acList
			return
		}
	}
}

// Add adds Access
func (a *AccessList) Add(ac Access) {
	if !a.Has(ac) {
		*a = append(*a, ac)
	}
}

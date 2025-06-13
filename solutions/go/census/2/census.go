// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	var resident Resident = Resident{Name:name,Age:age,Address:address}
    return &resident
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if v, ok := r.Address["street"]; !ok {
        return false
    } else if v == "" {
    	return false
    }
	if r.Name == "" {
        return false
    }
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
    r.Name = ""
    r.Age = 0
    for k,_ := range r.Address {
        delete(r.Address, k)
    }
    r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var c int = 0
    for i := 0; i < len(residents); i++ {
    	if residents[i].HasRequiredInfo() {
            c += 1
        }
    }
	return c
}

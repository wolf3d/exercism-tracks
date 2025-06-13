package sorting

import "fmt"
import "strconv"

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if val, ok := fnb.(FancyNumber); ok {
        if v, err := strconv.Atoi(val.Value()); err == nil {
            return v
        } else { 
        	return 0
        }
    }    
	return 0    
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {    
	var message string
	switch t := i.(type) {
        case string:
        	if val, err := strconv.Atoi(t); err == nil {
                fmt.Printf("wtf %v",val)
    			message = DescribeNumber(float64(val))                        
            } else {
            	message = "Return to sender"
            }
    
        case int: 
    		message = DescribeNumber(float64(t))
        case float64:
    		message = DescribeNumber(t)        
    	case NumberBox:
			v, _ := t.(NumberBox)
    		message = DescribeNumberBox(v)
    	case FancyNumberBox:
			v, _ := t.(FancyNumber)    
    		message = DescribeFancyNumberBox(v)
    	default: message = "Return to sender"
    }
	return message
}

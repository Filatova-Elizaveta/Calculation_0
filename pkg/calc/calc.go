package calc

import (
    "fmt"
    "strconv"
    "strings"
)
func Operation (operands []float64, operations []rune) ([]float64, []rune) {
	b := operands[len(operands)-1]
	a := operands[len(operands)-2]
	operands = operands[:len(operands)-2]
    
	switch operations[len(operations)-1] {
	case '+':
		operands = append(operands, a+b)
	case '-':
		operands = append(operands, a-b)
	case '*':
		operands = append(operands, a*b)
	case '/':
		operands = append(operands, a/b)
	}
	operations = operations[:len(operations)-1]
	return operands, operations
}

func Calc(expression string) (float64, error) {
    count := 0
    expression = strings.ReplaceAll(expression, " ", "")
    if expression == "" {
        return 0, fmt.Errorf("Error")
    }
    runes := []rune(expression)
    var operands []float64
    var operations []rune
    if runes[len(runes)-1] == '+' || runes[len(runes)-1] == '*' || runes[len(runes)-1] == '/' || runes[len(runes)-1] == '-' || runes[len(runes)-1] == '(' {
        return 0, fmt.Errorf("Error")
}
    if runes[0] == '+' || runes[0] == '*' || runes[0] == '/' || runes[0] == '-' || runes[0] == ')' {
        return 0, fmt.Errorf("Error") }
    for i := 0; i < len(runes); i++ {
	    if runes[i] >= '0' && runes[i] <= '9' {
		    start := i
		    for i < len(runes) && ((runes[i] >= '0' && runes[i] <= '9') || runes[i] == '.') {
			    i++
		}
    		number, err := strconv.ParseFloat(string(runes[start:i]), 64)
	    	if err != nil {
		    	return 0, fmt.Errorf("Error")
		}
    		operands = append(operands, number)
	    	i--
    	} else if runes[i] == '(' {
	    	operations = append(operations, runes[i])
	    	count++
    	} else if runes[i] == ')' {
    	    if count == 0 {
    	        return 0, fmt.Errorf("Error") }
    	    count--
	    	for len(operations) > 0 && operations[len(operations)-1] != '(' {
		    	operands, operations = Operation(operands, operations)
		}
		    operations = operations[:len(operations)-1]
    	} else if runes[i] == '+' || runes[i] == '-' || runes[i] == '*' || runes[i] == '/' {
    	    if runes[i-1] == '+' || runes[i-1] == '-' || runes[i-1] == '*' || runes[i-1] == '/' {
	    	    return 0, fmt.Errorf("Error") }
	    	for len(operations) > 0 && (operations[len(operations)-1] == '*' || operations[len(operations)-1] == '/') {
		    	if (operations[len(operations)-1] == '/') && (operands[len(operands)-1] == 0) {
		    	    return 0, fmt.Errorf("Error")
		    	}
		    	operands, operations = Operation(operands, operations)
		}
	    	operations = append(operations, runes[i])
	    } else {
		    return 0, fmt.Errorf("Error")
	}
}
    if count != 0 {
        return 0, fmt.Errorf("Error")
    }
    for len(operations) > 0 {
	    operands, operations = Operation(operands, operations)
}

    if len(operands) != 1 || len(operations) != 0 {
	    return 0, fmt.Errorf("Error")
}

    return operands[0], nil
}

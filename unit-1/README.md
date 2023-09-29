## Unit 0. Getting started

- Know what sets Go apart
- Visit the Go Playground
- Print text to the screen
- Experiment with text in any natural language


## Unit 1. Imperative programming
### LESSON 1
### 2.1. Performing calculations
- Mars is 37.83% of the weight on Earth
- Mars takes 687 Earth days to travel around the sun


### 2.2. Formatted print
- Unlike Print and Println, the first argument to Printf is always text. 
- The text contains the format verb %v, which is substituted for the value of the expression provided by the second argument.
- Printf can help you align text eg. padding


### Quick check 2.5
Write the shortest line of code to subtract two pounds from a variable named weight.
soln: weight -= 2


### Summary
- The Print, Println, and Printf functions display text and numbers on the screen.
- With Printf and the %v format verb, values can be placed anywhere in the displayed text.
- Constants are declared with the const keyword and can’t be changed.
- Variables are declared with var and can be assigned new values while a program is running.
- The math/rand import path refers to the rand package.
- The Intn function in the rand package generates pseudorandom numbers.
- You used 5 of the 25 Go keywords: package, import, func, const, and var.


### LESSON 3: Loops and branches
- Have your computer make choices with if and switch
- Repeat code with for loops
- Use conditions for looping and branching

In Go, the only true value is true and the only false value is false.

The rules for determining a leap year are as follows:
- Any year that is evenly divisible by 4 but not evenly divisible by 100
- Or any year that is evenly divisible by 

#### Go uses short-circuit logic;
- If the first condition is true (the year is evenly divisible by 400), there’s no need to evaluate
what follows the || operator, so it is ignored.

### Branching with switch
- One unique feature of switch is the fallthrough keyword, which is used to execute the body of the next case. Go takes a safer approach, requiring an explicit fallthrough keyword.

### Summary
1. Booleans are the only values that can be used in conditions.
2. Go provides branching and repetition with if, switch, and for.
2. You used 12 of the 25 Go keywords: package, import, func, var, if,
else, switch, case, default, fallthrough, for, and break.
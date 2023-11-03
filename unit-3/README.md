## Unit 3. Building blocks

Programming is the breaking of one big impossible task into several very
small possible tasks. ~ Jazzwant

#### Functions
Parameter and argument are terms from mathematics, with a subtle distinction.
A function accepts parameters and is invoked with arguments, though at times
people may use the terms interchangeably.

Lowercase functions can only be used by the package they are declared in, 
whereas capitalized functions are exported for use anywhere.

The ellipsis (...) in a function indicate that the  function is variadic. 
You can pass it a variable number of arguments

side-effect-free functions:
Its only input is the parameter it accepts, and its only output is theresult it returns. 
It makes no modifications to external state. 

(CHECK): Splitting code into functions;
- Functions are reusable, 
- they provide isolation for variables through function scope, and 
- they provide a name for the action they perform which makes code easier to follow and understand.


### Methods
Methods are like functions that enhance types with additional behavior. 
Methods provide another way to organize code, an arguably nicer way

Functions have a place, but types and methods provide another useful way to organize code and represent the world around you.

Being able to define your own types can be incredibly useful for improving both readability and reliability.
## Unit 2. Types
This covers the primitive types that Go provides for text, characters, numbers and other simple values.

-> Go has two floating-point types;
- The default floating-point type is float64, a 64-bit floating-point type that uses eight bytes of memory.
- The float32 type uses half the memory of float64 (4 bytes or 32 bits) but offers less precision. This type is sometimes called single precision.

#### 6.2. Displaying floating-point types
Printf with the %f formatting verb to specify the number of digits

Quick check 6.4: 
1. What is the width and precision of 0015.1021?
The width is 9 and the precision is 4, with zero padding "%09.4f".

- floating-point isn’t the best choice for representing money
- To minimize rounding errors, we recommend that you perform multiplication
before division.

- To take the absolute value of a float64, the math package provides an ```Abs``` function

```fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)```

#### Summary
- Go can infer types for you. In particular, Go will infer float64 for variables initialized with real numbers.
- Floating-point types are versatile but not always accurate.
- You used 2 of Go’s 15 numeric types (float64, float32).


#### Lesson 7. Whole numbers
- Go offers 10 different types for whole numbers, collectively called integers: (can’t store fractional numbers and they have a limited range)

- A single hexadecimal digit consumes four bits, called a nibble.
- To distinguish between decimal and hexadecimal, Go requires a 0x prefix for hexadecimal.

- The %b format verb will show you the bits for an integer value.
- The simplest way to avoid wrapping is to use an integer type large enough to hold the values you expect to store

- Only the int64 and uint64 integer
types are able to store numbers well beyond two billion on all platforms.

#### Summary
1. The most common integer types are int and uint, but some situations
call for smaller or larger types.
2. Integer types need to be chosen carefully to avoid wrapping around, unless
wrapping is what you want.
3. You looked at 10 more of the 15 numeric types in Go (int, int8, int16,
int32, int64, uint, uint8, uint16, uint32, uint64).


#### Lesson 8: Big numbers

The big package provides three types:
1. big.Int is for big integers, when 18 quintillion isn’t enough.

2. big.Float is for arbitrary-precision floating-point numbers.

3. big.Rat is for fractions like ⅓.

- The NewInt function takes an int64 and returns a big.Int:

#### Summary
- When the native types can’t go the distance, the big package has you
covered.
- Big things are possible with constants that are untyped, and all numeric
literals are untyped constants too.
- Untyped constants must be converted to typed variables when passed to
functions.
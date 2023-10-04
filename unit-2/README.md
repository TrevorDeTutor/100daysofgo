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

#### Summary
- Go can infer types for you. In particular, Go will infer float64 for variables initialized with real numbers.
- Floating-point types are versatile but not always accurate.
- You used 2 of Go’s 15 numeric types (float64, float32).
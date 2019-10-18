#### Names
Keys:
```
break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var
```
Constants: 
```
true false iota nil
```
Types: 
```
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64
bool byte rune string error
```
Functions:
```
make len cap new append copy close delete
complex real imag
panic recover
```

#### Declarations
```
var, const, type, and func
```
##### Variables
"var name type = expression"  
The **type** or the **= expression** part maybe omitted, but not both. 
  
##### zero value
If the expression is omitted the initial value is the **zero value** for the type:
- 0 for numbers.
- false for booleans.
- "" for strings.
- nil for interfaces and reference types (slice, point er, map, channel, function). 
- The **zero value** of anaggregate type（array,struct) has the zero value of all of its elements or fields.
```go
var name type = expression 
var boiling float64 = 100 // a float64
var boiling = 100 // type part omitted
var boiling float64 // = expression part omitted
```
单行声明多个变量
- 用type声明多个相同类型的变量
```go
var i, j, k int // int, int, int
```
- 省略类型可以单行声明多种类型的变量
```go
// Omitting the type allows declaration of multiple variables of different types
var b, f, s = true, 2.3, "four" 
```

```go
var f, err = os.Open(name) // os.Open returns a file and an error
```

```
Short Variable Declarations
```go
i := 100 // an int
i, j := 0, 1
i, j = j, i // swap values of i and j
```

### pointer

#### 变量
>A variable is a piece of storage containing a value. Variables created by declarations are identified by a name, such as x, but many variables are identified only by expressions like x[i] or x.f. All these expressions read the value of a variable, except when the y appear on the left-hand side of an assignment, in which case a new  alueisassigned to the variable.
>>变量是包含一个值的存储器。由声明创建的变量由名称（如x）标识，但许多变量仅由表达式（如x[i]或x.f）标识。所有这些表达式都读取变量的值，除非y出现在赋值的左侧，在这种情况下，将新值分配给变量。
#### pointer
>A pointer value is the address of a variable. A pointer is thu s the location at which a value is stored. Not every value has an address, but every variable does. With a pointer, we can read or update the value of a variable in directly, without using or even knowing the name of the variable, if indeed it has a name.
>>指针值是变量的地址。指针是指存储值的位置。不是每个值都有地址，但每个变量都有地址。有了指针，我们可以直接读取或更新变量的值，而不必使用甚至不知道变量的名称（如果变量确实有名称的话）。

>If a variable is declared “var x int”, the expression &x (‘‘address of x’’) yields a pointer to an integer variable, that is, a value of type *int, which is pronounced ‘‘pointer to int.’’ If this value is called “p”, we say ‘‘p points to x,’’ or equivalently ‘‘p contains the address of x.’’ The variable to which p points is written *p. The expression *p yields the value of that variable, an int, but since \*p denotes a variable, it may also app ear on the left-hand side of an assignment, in which case the assignment updates the variable.
>>如果一个变量声明为“var x int”,表达式&x(x的地址)产生一个指向整数变量的指针，指针的数据类型是\*int，读作“指向int”，如果&x的值叫做“p”，我们读作“p指向x”或者说“p包含x的地址(p contains the address of x)”，指针变量写作\*p(The variable to which p points is written \*p),表达式\*p是变量的值即数据存储块（The expression \*p yields the value of that variable）,这样\*p表示一个变量,它也可以出现在赋值的左侧，在这种情况下赋值会更新变量。


```
```
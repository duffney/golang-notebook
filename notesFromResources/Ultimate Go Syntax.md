Author(s): William Kennedy   
Url: https://courses.ardanlabs.com/  
Date published: 2022  
tags:: course software development   

- ### Variables  
	- Any value constructed in memory is at least initialized to its `zero value`. zero value golang/variables  
		- Float are set to `0`, strings are empty `" ",` bools are `false`.  
	- zero value golang/idioms golang/variables  
		- Use `var` to denote a variables set to their zero value state.  
			- var time int  
	- type conversion golang/variables  
		- Go uses `type conversion` instead of casting.  
			- //To perform type conversions you must be explicit.  
			  var1 := 10  
			  var2 := int32(var1) // explicit conversion  
			    
			  fmt.Printf("var1 = %Tn", var1) // type int  
			  fmt.Printf("var2 = %Tn", var2) // type int32  
- ### Struct  
	- Links:  
		- https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/struct_types  
	- Go isn't really an object-oriented language, but a data oriented langauge  
	- literal construction is when you define a variable of a specific type and define the values of that type to something other than their zero value. golang/types literal construction  
	- Avoid partially constructed types golang/types golang/idioms  
		- https://go.dev/play/p/djzGT1JtSwy  
- ### Pointers  
	- Value and pointer semantics. golang/types/pointers  
		- value semantic, everyone working on their own copy of the data.  
			- Trade-offs: cost is extra complexity in the code, maintains high integrity.  
				- example: https://go.dev/play/p/9kxh18hd_BT  
		- Pointer semantic, everyone shares a single copy  
			- Trade-offs: increased efficiency and reduce complexity, but can lose integrity.  
				- example: https://go.dev/play/p/mJz5RINaimn  
	- Variables contain two types of data; the `value of` and the `address`. Pointers use the address to determine the location, where, the value of the variable is stored. golang/types/pointers  
		-  
		  ``` go
		  			  // Declare variable of type int with a value of 10.
		  			  count := 10
		  			  
		  			  // Display the "value of" and "address of" count.
		  			  println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
		  ```
			- Calling `count`, displays the value of the variable while `&` displays the address  
	- Pointers have three levels. A `value` which is the address in memory to a value. But a pointer is a variable, so it also has an address that can be accessed with `&value`. And `*value`, which is the value of what the pointer points to. golang/types/pointers  
		- ![Screen_Shot_2022-04-27_at_1.10.22_PM_1651083026731_0.png](../assets/Screen_Shot_2022-04-27_at_1.10.22_PM_1651083026731_0_1652896003772_0.png)  
			- `inc` is the value of  
			- `&inc` is the address of  
			- `*inc` is the value that the pointer points to  
- ### Literal structs  
	- A type only needs to be named if it's going to be used more than once. Unnecessarily naming types leads to "type pollution", use anonymous types when naming isn't required. golang/types/structs  
		-  
		  ``` go
		  			  //Delcare a variable of an anonymous type set to its zero value
		  			  var e1 struct {
		  			    flag bool
		  			    counter int16
		  			    pi float32
		  			  }
		  ```
		-  
		  ``` go
		  			  //Declare a variable of an anonymous type and init using a literal struct
		  			  e2 := struct {
		  			    flag bool
		  			    coutner int16
		  			    pi float32
		  			  }{
		  			    flat: true,
		  			    counter: 10,
		  			    pi: 3.141592,
		  			  }
		  ```
- ### Constants  
	- Constants in Go aren't variables and have their own type system. golang/types constants  
		-  
		  ``` go
		  			  // Untyped costant
		  			  const myUnit8 = 12345 // Kind: integer, no error (implicit compiler conversion)
		  			  
		  			  // Type constants restricts percision
		  			  const myUnit8 unit8 == 1000 // overflow error cannot exceed 255 
		  ```
	- `const` kind promotion system golang/types constants  
		- example `ints` can be promoted to `float32`, this eliminates the need to do constant type conversion with constants.  
	- `iota` gives you a way to automatically increment golang/types constants  
		-  
		  ``` go
		  			  const (
		  			    A2 = iota // 0
		  			    B2        // 1
		  			  )
		  ```
	- Understand the concept of enumerators gap  
	- Don't create aliases to create compiler protections. golang/types # kind promotion golang/idioms  
		-  
		  > Go's "kind promotion" while allows you to write really nice code in terms of literal values destroys your ability to have enumerations.  

		- Good  
			-  
			  ``` go
			  				  type color int
			  				  
			  				  const (
			  				  	red = 10
			  				      blue = 20
			  				  )
			  				  
			  				  func PrintColor(color int){
			  				    fmt.Println(color)
			  				  }
			  ```
		- Bad  
			-  
			  ``` go
			  				  type color int
			  				  
			  				  const (
			  				  	Red color = 10
			  				      Blue color = 20
			  				  )
			  				  
			  				  func PrintColor(c color){
			  				    fmt.Println(c)
			  				  }
			  				  
			  				  // because of kind promtion any int can be passed in, which bypassed
			  				  // compiler protection.
			  				  PrintColor(30)
			  ```
- ### Function idioms  
	- Returning multiple values is a unique feature of Go. The beauty of multiple returns is that you can return a value and an error. golang/functions  
		- The Go standard library return a max of three values.  
	- Avoid else at all costs golang/functions golang/idioms  
		- Use the column within a function to outline the happy path.  
		- Else statements drastically reduce code readability by disrupts the vertical visualization of the happy path.  
	- Return out of a function as soon as possible. golang/functions golang/idioms  
		-  
		  ``` go
		  			  r, err := getUser(name)
		  			  if err != nil {
		  			    return nil, err // "idiom" in Go, return the zero value literal when returning an error 
		  			  }
		  ```
	- When returning values from a function with multiple returns, only return the value or the error. Use the zero value literal instead of both at once. golang/functions golang/idioms  
		- Improves readability  
	- Scope an error variable within an `if` statement.  golang/control structures/if golang/error handling  
		-  
		  ``` go
		  			  if err := json.Unmarshal([]byte(r), &u); err != nil {
		  			    return nil, err
		  			  }
		  ```
- ### Literal functions  
	- Go supports first-class functions, which means functions can be named or literal (unnamed). golang/functions  
		- What is a first-class function?  
	- Go supports closures  
		- What is closure?  
	- When you use `defer`, the last one you define is the first one to execute. defer golang/functions/defer  
- ### Blank identifiers  
	- A blank identifier ignores unused values returned by functions. blank identifiers golang/functions/  
		- `_, err := function()`, used when the syntax requires a variable name but the program logic does not.  
- ### Short variable declaration operator  
	- When the short variable declaration operator is being used, it only needs to construct one new variable and can reuse anything else. short variable declaration golang/variables  
	- Unique error variables should be avoided. Seeing an `error` variable not named `err` is a code smell. Most of the time `err` should be usable throughout an entire function. golang/error handling golang/idioms  
	- To maintain Go's no unique error variables use an `if` statement to limit the scope of an `err` variable.  
		- Scope an error variable within an `if` statement.  golang/control structures/if golang/error handling  
- ### Arrays  
	- Go's compiler needs to know the length of an array at compile time.   golang/data structures/array  
	- Go only uses the keyword `for` for all iterations.  golang/control structures/for  
	- while loop  golang/control structures/for  
		-  
		  ``` go
		  			  for i < len(numbers) {
		  			    fmt.Println(i, numbers[i])
		  			  }
		  ```
	-  
-  
  > Author(s): William Kennedy with Brian Ketelsen and Erik St. Martin  
  Title: Go in Action  
  Edition:  
  Source code: https://github.com/goinaction/code   
  Publisher: Manning Publication  
  Date published: 2015  
  ISBN: 9781617291784  
  tags:: book  software development golang  

- Chapter 2. Go quick-start  
	- For a package to produce an executable, the `main` function must be declared. If it's not declared the `go build` command won't produce an executable. golang/packages  
	- Variables located outside the scope of any function are consider `package-level variables`, and are accessible anywhere within the package. golang/variables  
	- An `anonymous function` is a function that's declared without a name. golang/functions  
	- If the `interface` type contains only one method, the name of the interface should end with the `er` suffix. But, when declaring multiple methods, the name of the interface should be related to its general behavior.  golang/types/interfaces  
	- Use the blank identifier as an alias import name to prevents the compiler from producing an error for declaring an import. golang/variables blank identifier  
		-  
		  ``` go
		  			  import (
		  			  	_ "github.com/goinaction/code/package"
		  			  )
		  ```
- Chapter 3. Packaging and tooling  
	- All Go programs are organized into groups of files called packages. And every `.go` file must declare what package it belongs to on the first line. golang/packages  
		-  
		  ``` go
		  			  package main // 'main' is the package name
		  ```
	- By naming convention, packages are named after the directory that contains it. Keeping names unique isn't required because you import packages using its full path. golang/packages  
		-  
		  ``` bash
		  			  -- nameOfPackage
		  			  	-- file.go // package nameOfPackage
		  ```
	- In Go's documentation a command is used to refer to any executable program and a package as an importable unit of functionality. Both consist of and make use of Go packages, but the executable command contains the `main()` function, that compiles an executable. Non-executable packages are sometimes called libraries. golang/packages  
	- Remote imports use distributed version control systems (DVCS) such as GitHub to manage external dependencies.  golang/packages remote import  
		-  
		  ``` go
		  			  import "github.com/spf13/viper"
		  ```
	- When you need to import two packages of the same name you use a `named import`, which allows you to specify and alias for the package within the code. golang/packages named import  
		-  
		  ``` go
		  			  import (
		  			  	"fmt"
		  			    	myfmt "mylib/fmt"
		  			  )
		  ```
	- All packages in Go have the ability to run an `init` function that is scheduled by the complier to be run before the `main` function is executed. Init is used to initialize or bootstrap anything your program needs prior to execution. golang/functions/init init  
		-  
		  ``` go
		  			  // example, load db driver
		  			  func init() {
		  			    sql.Register("postgress", new(PostgresDriver))
		  			  }
		  ```
- Chapter 5. Go's type system  
	- 5.1 User-defined types  
		-  
		  > Go is a statically typed language. That means that the compiler always wants to know what the type for every value in the program. Knowing that information a head of time reduces potential memory corruption and bugs. golang  

		- A value's type provides the compiler with two pieces of information; how much memory to allocate and what that memory represents. golang/types  
		- The `struct` keyword in Go allows you to create a composite type. golang/types/structs  
			-  
			  ``` go
			  				  type user struct {
			  				    name string
			  				    email string
			  				    ext int
			  				    privileged bool
			  				  }
			  ```
		- "Any time a variable is created and initialized to it's `zero value`, it's idiomatic to use the key word `var`" golang/variables zero value  
			-  
			  ``` go
			  				  var bill user // variable name is bill with the type 'user'
			  ```
		- If the variable is initialized to something other than its zero value, then use the short variable declaration with a struct literal. golang/types/structs  
			- Most common form:  
				-  
				  ``` go
				  					  lisa := user {
				  					    name: "Lisa",
				  					    email: "lisa@email.com",
				  					    ext: 123,
				  					    privileged: true,
				  					  }
				  ```
			- Less common form:  
				-  
				  ``` go
				  					  lisa := user{"Lisa","lisa@email.com",123,true}
				  ```
		- Go's compiler doesn't implicitly convert values of different types. golang/types  
	- 5.2 Methods  
		- Methods add behavior to user-defined types. Think of methods has functions that are declared using the attached type's name as an input type. golang/types/methods  
			-  
			  ``` go
			  				  type user struct {
			  				    name string
			  				    email string
			  				  }
			  				  
			  				  // associates the notify method with the type 'user'
			  				  func (u user) notify() { 
			  				    fmt.Printf("Sending User Email to %s,%s>\n", u.name, u.email)
			  				  }
			  ```
		- A receiver is the parameter between the `func` keyword and the function name. Functions that have a receiver are called methods. receiver golang/types/methods  
			-  
			  ``` go
			  				  func (u user) notify()
			  ```
				- `(u user)` is the receiver that makes the function a method. `notify()` is the name of the function. Well, it's the name of the method because the func has a receiver.  
		- Go has two types of receivers; value and pointer.  receiver golang/types/methods  
			-  
			  ``` go
			  				  func (u user) notify() {} // value receiver
			  				  
			  				  func (u *user) changeEmail(email string) {} // pointer receiver
			  ```
	- 5.3 The nature of types  
		- Reference types  
			- String, slice, map, channel, interface, and function are reference types, sometimes called built-in types. golang/types reference types  
			- When you declare a variable from  a reference type, the value that's created is called a header value. All header values contain a pointer to an underlying data structure. You never share reference type values with a pointer. Instead, you should pass a copy of the reference type by value to share the underlying data structure.  golang/types reference types  
				-  
				  ``` go
				  					  type IP []byte //ref type
				  					  
				  					  func ipEmptyString(ip IP) string { // passed by value
				  					      if len(ip) == 0 {
				  					          return ""
				  					      }
				  					      return ip.String()
				  					  }
				  ```
				- reference type values are treated like primitive data values.  
					- What is a primitive data value?  
		- Struct types  
			- ?? No clue  
	- 5.4 Interfaces  
		- Interfaces are types that declare behavior. golang/types/interfaces  
		- The type of A receiver is the parameter between the `func` keyword and the function name. Functions that have a receiver are called methods. receiver golang/types/methods used for a method determines if it accepts an input by value or by pointer.  receiver golang/types/methods  
			-  
			  ``` go
			  				  package main
			  				  
			  				  import (
			  				    "fmt"
			  				  )
			  				  
			  				  type notifier interface {
			  				    notify()
			  				  }
			  				  
			  				  type user struct {
			  				    name string
			  				    email string
			  				  }
			  				  
			  				  func (u *user) notify() {
			  				    fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
			  				  }
			  				  
			  				  func sendNotification(n notifier){
			  				    n.notify()
			  				  }
			  				  
			  				  func main() {
			  				    u := user{"Bill", "bill@email.com"}
			  				    sendNotification(u) 
			  				    // err: “Cannot use u (type user) as type notifier in argument
			  				    //       to sendNotification: user does not implement notifier”
			  				  }
			  ```
				- Because the notify method uses a pointer receiver and not a value receiver only pointers can be passed to the notify interface of type user.  
					- What if another method for admin used a value receiver? Would it work for that type?  
						- Yes, it does work. Interfaces can be associated with methods that have different receivers.  
							-  
							  ``` go
							  								  func (a admin) notify() {
							  								  	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
							  								  }
							  								  
							  								  func main(){
							  								  	a := admin{"John", "john@email.com"}
							  								  	sendNotification(a) // pass a value of type admin
							  								  }
							  ```
					-  
		- Receivers are what determine the method set of an interface.   golang/types/methods  
			-  
			  ``` go
			  				  func (u *user) notify() { // associates the method with a pointer
			  				    fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
			  				  }
			  				  
			  				  func sendNotification(n notifer)
			  ```
		-  
		-  
		  > Polymorphism is the ability to write code that can take on different behavior through the implementations of types. programming/terms  

		- Interfaces are how you implement   
		  > Polymorphism is the ability to write code that can take on different behavior through the implementations of types. programming/terms  

		  in Go.  golang/types/interfaces
			- Both type `user` and type `admin` use the `notifier` interface by declaring a `notify` method.  
				-  
				  ``` go
				  					  package main
				  					  
				  					  import (
				  					  	"fmt"
				  					  )
				  					  
				  					  type notifier interface {
				  					    notify() // any type with the method notify() is also of type notifier
				  					  }
				  					  
				  					  type user struct {
				  					  	name  string
				  					  	email string
				  					  }
				  					  
				  					  func (u *user) notify() { //receiver accepts pointer values
				  					  	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
				  					  }
				  					  
				  					  type admin struct {
				  					  	name  string
				  					  	email string
				  					  }
				  					  
				  					  func (a admin) notify() { // receiver accpets pointers and values
				  					  	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
				  					  }
				  					  
				  					  func sendNotification(n notifier) {
				  					  	n.notify()
				  					  }
				  					  
				  					  func main() {
				  					  	u := user{"Bill", "bill@email.com"}
				  					  	sendNotification(&u) // pass a pointer of type user
				  					  	// sendNotification(u)
				  					  	// err: “Cannot use u (type user) as type notifier in argument
				  					  	//       to sendNotification: user does not implement notifier”
				  					  	a := admin{"John", "john@email.com"}
				  					  	sendNotification(a) // pass a value of type admin
				  					      sendNotification(&a) //pass a pointer value of type admin
				  					  }
				  ```
- Chapter 8. Standard library  
	- `Marshaling` is the process of transforming data into a JSON string. `Unmarshaling` is the process of processing a JSON string into a struct type.   golang/encoding  
	- `tags` provide metadata about the field mappings between JSON and struct types.  golang/encoding  
		-  
		  ``` go
		  			  type Contact struct {
		  			    Name string `json:"name"` // string is the 'tag'
		  			  }
		  ```
	- Create a The `struct` keyword in Go allows you to create a composite type. golang/types/structs and use `tags` to map the JSON to custom fields.   golang/encoding  
		-  
		  ``` go
		  			  type Contact struct {
		  			  	Name    string `json:"name"`
		  			  	Title   string `json:"title"`
		  			  	Contact struct {
		  			  		Home string `json:"home"`
		  			  		Cell string `json:"cell"`
		  			  	} `json:"contact"`
		  			  }
		  			  
		  			  var JSON = `{
		  			  	"name": "Gopher",
		  			  	"title": "programmer",
		  			  	"contact": {
		  			  		"home":"415.333.3333",
		  			  		"cell":"415.555.5555"
		  			        }
		  			  	}`
		  ```
			- Go playground example: https://go.dev/play/p/r8iMoFJ_OtY  
	- Marshal JSON into a string   golang/encoding  
		- When using a `map` with an `interface{}` type, there's no need to declare new types.  
			-  
			  ``` go
			  				  c := make(map[string]interface{})
			  				  c["name"] = "Gopher"
			  				  c["title"] = "programmer"
			  				  c["contact"] = map[string]interface{}{
			  				    "home": "415.333.3333",
			  				    "cell": "415.555.5555",
			  				  }
			  				  
			  				  data, err := json.MarshalIndent(c, "", "    ")
			  ```
				- [Go Playground](https://go.dev/play/p/u-u5l5rGwKV)  
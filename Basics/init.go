package basics

import "fmt"

func init() {
	fmt.Println("Init function executed")
}

func main() {
	//Init functions are special functions in Go that are automatically executed when a package is initialized.
	//They are typically used to set up package-level variables, perform initialization tasks, or register components before the main program starts executing.
	//The init function does not take any parameters and does not return any values.
	//A package can have multiple init functions, and they will be executed in the order they are defined within the package.
	//If a package imports other packages, the init functions of the imported packages will be executed before the init function of the importing package.
	//The init function is called only once per package, regardless of how many times the package is imported.
	fmt.Println("Main function executed")

	//Practical use cases of init functions:
	//1. Configuration Setup: Init functions can be used to load configuration settings from files or environment variables and initialize package-level variables accordingly.
	//2. Resource Initialization: Init functions can be used to set up connections to databases, initialize caches, or allocate other resources needed by the package.
	//3. Registration of Components: Init functions can be used to register components, such as plugins or handlers, with a central registry or framework.
	//4. Validation: Init functions can perform validation checks on package-level variables or configurations to ensure they meet certain criteria before the package is used.
	//5. Database  Initialization: Init functions can be used to set up database connections, create necessary tables, or perform migrations when the package is imported.
	//Best Practices:
	//1. Keep init functions simple
	//2. Avoid side effects,order of execution should be predictable
	//3. Use for essential setup only
}

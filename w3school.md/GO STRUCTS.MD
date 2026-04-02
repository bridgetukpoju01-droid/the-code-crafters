1. **What is a struct in Go?**
   A struct is a way to group different types of data into a single variable. It is useful for creating records with various pieces of information.

2. **Arrays vs Structs:**
   While arrays store multiple values of the same type, structs store multiple values of different types together.

3. **Declaring a Struct:**
   To declare a struct, use `type struct_name struct {}` and define members inside it.
   Example:

   ```go
   type Person struct {
     name string
     age int
     job string
     salary int
   }
   ```

4. **Accessing Struct Members:**
   To access a member of a struct, use the dot operator (`.`).
   Example:

   ```go
   pers1.name = "Hege"
   ```

5. **Example of Struct with Data:**
   You can create multiple instances of a struct and assign different values to each.
   Example:

   ```go
   var pers1 Person
   pers1.name = "Hege"
   ```

6. **Passing Struct as Function Argument:**
   You can pass structs to functions to access their data.
   Example:

   ```go
   func printPerson(pers Person) {
     fmt.Println(pers.name)
   }
   ```

7. **Result of Accessing Struct Members:**
   When accessing the members of the struct, it will print each member's value.
   Example Output:

   ```
   Name: Hege
   Age: 45
   Job: Teacher
   Salary: 6000
   ```

8. **Function with Struct Argument:**
   You can use a function to display a struct's data.
   Example:

   ```go
   printPerson(pers1)
   ```

9. **Why use structs?**
   Structs help organize data, making it easier to manage and use in your programs.

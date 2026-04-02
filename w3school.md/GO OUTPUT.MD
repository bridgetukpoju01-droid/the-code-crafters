In Go, there are three functions used to print or output text: `Print()`, `Println()`, and `Printf()`.

1. **Print()**: This function prints its arguments without any added formatting. It prints the values exactly as they are provided, without automatically adding spaces or newlines. If you want a new line, you need to manually include `\n` in the string. For example:

   ```go
   fmt.Print("Hello", "\n", "World")
   ```

   This will print:

   ```
   Hello
   World
   ```

   If the arguments are not strings, `Print()` automatically adds a space between them.

2. **Println()**: This function works like `Print()`, but it automatically adds a space between arguments and a newline at the end of the output. For example:

   ```go
   fmt.Println("Hello", "World")
   ```

   This will print:

   ```
   Hello World
   ```

3. **Printf()**: This function is used when you want to format the output. You use **formatting verbs** to specify how the values should be displayed. Two common formatting verbs are `%v` (to print the value) and `%T` (to print the type). For example:

   ```go
   fmt.Printf("i has value: %v and type: %T\n", "Hello", "Hello")
   ```

   This will print:

   ```
   i has value: Hello and type: string
   ```

These functions help you control how you display values in Go, whether you need simple printing, formatting, or spacing between values.

Base: https://github.com/quii/learn-go-with-tests

# Conteudo

- Install Go - Set up environment for productivity.
- [Hello, world - Declaring variables, constants, if/else statements, switch, write your first go program and write your first test. Sub-test syntax and closures.](hello)
- [Integers - Further Explore function declaration syntax and learn new ways to improve the documentation of your code.](integers)
- [Iteration - Learn about for and benchmarking.](iteration)
- [Arrays and slices - Learn about arrays, slices, len, varargs, range and test coverage.](arrays-and-slices)
- [Structs, methods & interfaces - Learn about struct, methods, interface and table driven tests.](structs-methods-and-interfaces)

- Pointers & errors - Learn about pointers and errors.
- Maps - Learn about storing values in the map data structure.
- Dependency Injection - Learn about dependency injection, how it relates to using interfaces and a primer on io.
- Mocking - Take some existing untested code and use DI with mocking to test it.
- Concurrency - Learn how to write concurrent code to make your software faster.
- Select - Learn how to synchronise asynchronous processes elegantly.
- Reflection - Learn about reflection
- Sync - Learn some functionality from the sync package including WaitGroup and Mutex
- Context - Use the context package to manage and cancel long-running processes
- Intro to property based tests - Practice some TDD with the Roman Numerals kata and get a brief intro to property based tests
- Maths - Use the math package to draw an SVG clock

# Pontos Legais

- GO já vem com um  coverage tool: `go test -cover`
- In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

# Duvidas / Leituras futuras

- Com ou sem o ```t.Helper()``` tive a mesma saida nos testes
- Ver Mais sobre 'Examples' nos testes
- Ver Mais sobre 'Benchmark' nos testes
- Porque temos que usar Slices, por causa dos 
  - 'if you try to pass an [4]int into a function that expects [5]int, it won't compile.' 
  - 'An interesting property of arrays is that the size is encoded in its type'
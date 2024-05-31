# blbuild

`blbuild` is a simple build tool for managing projects in a compiled language!

## Purpose

`blbuild` aims to simplify the compilation process using a minimalistic approach. It reads a TOML file, parses the values, assembles them into a string, and runs it as a command. That's all. 

## Features
 - Minimal configuration: only what is absolutely essential for the project.
 - No automation or scripting: the focus is purely compilation.
 - Easy to use: straightforward syntax and configuration. 


## Fields

 - **compiler** (string)
     - The compiler to use.
 - **files** (array of strings)
     - The files to be compiled.
 - **flags** (array of strings)
     - Compiler flags. 
     - Certain flags will require additional fields (e.g. -o would require an output file name).
 - **options** (array of arrays of strings; one array for each flag)
     - Options for the compiler flags. 
     - Array containing an array of strings for each flag.
     - *These must be presented in the same order as the options in the 'flags' field.*

## Examples

```toml
# blbuild
compiler = "gcc"
files = ["hello.c"]
flags = ["-o"]
options = [["hello"]]

# blbuild -> gcc hello.c -o hello
```

```toml
# blbuild
compiler = "g++"
files = ["one.cpp", "two.cpp", "three.cpp"]
flags = ["-o"]
options = [["run"]]

# blbuild -> g++ one.cpp two.cpp three.cpp -o run
```

```toml
# blbuild
compiler = "g++"
files = ["*.cpp"]
flags = ["--std=c++20", "-o"]
options = [[], ["run"]]

# blbuild -> g++ --std=c++20 *.cpp -o run
```
# blbuild

*blbuild is in its early stages and is currently non-functional.*

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
 - **path** (string)
     - The path to prepend to each file name.
 - **files** (array of strings)
     - The files to be compiled.
 - **extras** (array of strings)
     - Anything else to include in the command.


## Examples

```toml
# blbuild.toml
compiler = "gcc"
path = "" 
files = ["hello.c"]
extras = ["-o hello"]

# Result: gcc hello.c -o hello
```

```toml
# blbuild.toml
compiler = "g++"
path = "ex"
files = ["one.cpp", "two.cpp", "three.cpp"]
extras = ["-o run"]

# Result: g++ ex/one.cpp ex/two.cpp ex/three.cpp -o run
```

```toml
# blbuild.toml
compiler = "g++ --std=c++20"
path = ""
files = ["*.cpp"]
extras = ["-o run"]

# Result: g++ --std=c++20 *.cpp -o run
```
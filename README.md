# Bland Build Tool (blbld)

A simple build tool for C/C++ projects!

## Usage

The intended usage for this tool is to simplify the process of building projects that contain many files, but only a few that are changing at a given moment. 

This tool relies on a configuration file called `blbld.toml`. If this file does not exist, the program will prompt you for all of the necessary information, then generate the file for you. 

Currently, there are a few ways to interact with the CLI (For these examples '`blbld`' is going to act as the CLI command.):

### Build Project Directly
 - `blbld`: You can run the tool without any arguments.
   - This will parse the config file (or give you the prompts to generate one), construct a direct build command (e.g. `g++ -std=c++20 one.cpp two.cpp -o main`), and run it. 

### Compile All Files
 - `blbld compile`: You can also compile all of the files to their corresponding '.o' files, but does not link them.
   - This step is necessary to be able to link individual files back to the rest of the project. While the purpose of this tool is to avoid compiling the entire project every time, everything will still need to be compiled once at the start to be available to use. Luckily, this step is only necessary once. 
   - The output of this command will be all of the corresponding object files (e.g. `one.cpp` and `two.cpp` would become `one.o` and `two.o`)

### Compile Individual Files
 - `blbld compile two.cpp`: This is the important step; you can compile an individual file to its corresponding object file without touching the others.
   - This will compile a singular .cpp file to a singular .o file. 
 
### Build Compiled Files
 - `blbld build`: This is the step you take after compiling a single file.
   - This will link all of the object files and compile them to the final binary.

## Examples

> This project contains a `test` directory with some simple C++ files. To reset this directory, run `rm test/*.o test/main`.

The simplest way to run the project is with Go's toolchain:

No arguments
```bash
go run .
# will run:
g++ -std=c++20 test/one.cpp test/two.cpp test/three.cpp -o main
# and produce:
# test/main (an executable binary)
```

Compiling all files:
```bash
go run . compile
# will run:
g++ -std=c++20 -c test/one.cpp test/two.cpp test/three.cpp
# and produce:
# test/one.o test/two.o test/three.o
```

Compiling single file:
```bash
go run . compile two.cpp
# will print:
# Compiling 'two.cpp'...
# and produce:
# test/two.o
```

Building compiled files:
```bash
go run . build
# will run: 
g++ -std=c++20 test/*.o -o main
# and produce:
# main (an executable binary)
```

## Future Plans

I would like to be able to automate more of these steps. Ideally, I would like a single command to be able to rebuild the project when editing a single file (out of many). 

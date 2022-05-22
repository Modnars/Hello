# Step 1: A Basic Starting Point

The most basic project is an executable built from source code files.
For simple projects, a three line ``CMakeLists.txt`` file is all that is
required. This will be the starting point for our tutorial. Create a
``CMakeLists.txt`` file in the ``Step1`` directory that looks like:

```cmake
cmake_minimum_required(VERSION 3.10)

# set the project name
project(Tutorial)

# add the executable
add_executable(Tutorial tutorial.cxx)
```


Note that this example uses lower case commands in the ``CMakeLists.txt`` file.
Upper, lower, and mixed case commands are supported by CMake. The source
code for ``tutorial.cxx`` is provided in the ``Step1`` directory and can be
used to compute the square root of a number.

## Build and Run

That's all that is needed - we can build and run our project now! First, run
the [`cmake`](https://cmake.org/cmake/help/v3.23/manual/cmake.1.html#manual:cmake(1)) executable or the
[`cmake-gui`](https://cmake.org/cmake/help/v3.23/manual/cmake-gui.1.html#manual:cmake-gui(1)) to configure the project and then build it
with your chosen build tool.

For example, from the command line we could navigate to the
``Help/guide/tutorial`` directory of the CMake source code tree and create a
build directory:

```bash
mkdir Step1_build
```

Next, navigate to the build directory and run CMake to configure the project
and generate a native build system:

```bash
cd Step1_build
cmake ..  # 这里应该是源文档存在笔误，已按此修正
```

Then call that build system to actually compile/link the project:

```bash
cmake --build .
```

Finally, try to use the newly built ``Tutorial`` with these commands:

```bash
Tutorial 4294967296
Tutorial 10
Tutorial
```


## Adding a Version Number and Configured Header File

The first feature we will add is to provide our executable and project with a
version number. While we could do this exclusively in the source code, using
``CMakeLists.txt`` provides more flexibility.

First, modify the ``CMakeLists.txt`` file to use the [`project`](https://cmake.org/cmake/help/v3.23/command/project.html#command:project) command
to set the project name and version number.

```cmake
cmake_minimum_required(VERSION 3.10)

# set the project name and version
project(Tutorial VERSION 1.0)
```

Then, configure a header file to pass the version number to the source
code:

```cmake
configure_file(TutorialConfig.h.in TutorialConfig.h)
```

Since the configured file will be written into the binary tree, we
must add that directory to the list of paths to search for include
files. Add the following lines to the end of the ``CMakeLists.txt`` file:

```cmake
target_include_directories(Tutorial PUBLIC
                           "${PROJECT_BINARY_DIR}"
                           )
```

Using your favorite editor, create ``TutorialConfig.h.in`` in the source
directory with the following contents:

```cpp
// the configured options and settings for Tutorial
#define Tutorial_VERSION_MAJOR @Tutorial_VERSION_MAJOR@
#define Tutorial_VERSION_MINOR @Tutorial_VERSION_MINOR@
```

When CMake configures this header file the values for
``@Tutorial_VERSION_MAJOR@`` and ``@Tutorial_VERSION_MINOR@`` will be
replaced.

Next modify ``tutorial.cxx`` to include the configured header file,
``TutorialConfig.h``.

Finally, let's print out the executable name and version number by updating
``tutorial.cxx`` as follows:

```cpp
  if (argc < 2) {
    // report version
    std::cout << argv[0] << " Version " << Tutorial_VERSION_MAJOR << "."
              << Tutorial_VERSION_MINOR << std::endl;
    std::cout << "Usage: " << argv[0] << " number" << std::endl;
    return 1;
  }
```

## Specify the C++ Standard

Next let's add some C++11 features to our project by replacing ``atof`` with
``std::stod`` in ``tutorial.cxx``.  At the same time, remove
``#include <cstdlib>``.

```cpp
  const double inputValue = std::stod(argv[1]);
```

We will need to explicitly state in the CMake code that it should use the
correct flags. The easiest way to enable support for a specific C++ standard
in CMake is by using the [`CMAKE_CXX_STANDARD`](https://cmake.org/cmake/help/v3.23/variable/CMAKE_CXX_STANDARD.html#variable:CMAKE_CXX_STANDARD) variable. For this
tutorial, set the [`CMAKE_CXX_STANDARD`](https://cmake.org/cmake/help/v3.23/variable/CMAKE_CXX_STANDARD.html#variable:CMAKE_CXX_STANDARD) variable in the
``CMakeLists.txt`` file to ``11`` and [`CMAKE_CXX_STANDARD_REQUIRED`](https://cmake.org/cmake/help/v3.23/variable/CMAKE_CXX_STANDARD_REQUIRED.html#variable:CMAKE_CXX_STANDARD_REQUIRED)
to ``True``. Make sure to add the ``CMAKE_CXX_STANDARD`` declarations above the
call to ``add_executable``.

```cmake
cmake_minimum_required(VERSION 3.10)

# set the project name and version
project(Tutorial VERSION 1.0)

# specify the C++ standard
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)
```

## Rebuild

Let's build our project again. We already created a build directory and ran
CMake, so we can skip to the build step:

```bash
cd Step1_build
cmake --build .
```

Now we can try to use the newly built ``Tutorial`` with same commands as before:

```bash
Tutorial 4294967296
Tutorial 10
Tutorial
```

Check that the version number is now reported when running the executable without
any arguments.

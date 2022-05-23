# Step 2: Adding a Library

Now we will add a library to our project. This library will contain our own
implementation for computing the square root of a number. The executable can
then use this library instead of the standard square root function provided by
the compiler.

For this tutorial we will put the library into a subdirectory
called ``MathFunctions``. This directory already contains a header file,
``MathFunctions.h``, and a source file ``mysqrt.cxx``. The source file has one
function called ``mysqrt`` that provides similar functionality to the
compiler's ``sqrt`` function.

Add the following one line ``CMakeLists.txt`` file to the ``MathFunctions``
directory:

```cmake
add_library(MathFunctions mysqrt.cxx)
```

To make use of the new library we will add an [`add_subdirectory`](https://cmake.org/cmake/help/v3.23/command/add_subdirectory.html#command:add_subdirectory)
call in the top-level ``CMakeLists.txt`` file so that the library will get
built. We add the new library to the executable, and add ``MathFunctions`` as
an include directory so that the ``MathFunctions.h`` header file can be found.
The last few lines of the top-level ``CMakeLists.txt`` file should now look
like:

```cmake
# add the MathFunctions library
add_subdirectory(MathFunctions)

# add the executable
add_executable(Tutorial tutorial.cxx)

target_link_libraries(Tutorial PUBLIC MathFunctions)

# add the binary tree to the search path for include files
# so that we will find TutorialConfig.h
target_include_directories(Tutorial PUBLIC
                          "${PROJECT_BINARY_DIR}"
                          "${PROJECT_SOURCE_DIR}/MathFunctions"
                          )
```

Now let us make the ``MathFunctions`` library optional. While for the tutorial
there really isn't any need to do so, for larger projects this is a common
occurrence. The first step is to add an option to the top-level
``CMakeLists.txt`` file.

```cmake
option(USE_MYMATH "Use tutorial provided math implementation" ON)

# configure a header file to pass some of the CMake settings
# to the source code
configure_file(TutorialConfig.h.in TutorialConfig.h)
```

This option will be displayed in the [`cmake-gui`](https://cmake.org/cmake/help/v3.23/manual/cmake-gui.1.html#manual:cmake-gui(1)) and
[`ccmake`](https://cmake.org/cmake/help/v3.23/manual/ccmake.1.html#manual:ccmake(1))
with a default value of ``ON`` that can be changed by the user. This setting
will be stored in the cache so that the user does not need to set the value
each time they run CMake on a build directory.

The next change is to make building and linking the ``MathFunctions`` library
conditional. To do this,  we will create an ``if`` statement which checks the
value of the option.  Inside the ``if`` block, put the
[`add_subdirectory`](https://cmake.org/cmake/help/v3.23/command/add_subdirectory.html#command:add_subdirectory) command from above with some additional list
commands to store information needed to link to the library and add the
subdirectory as an include directory in the ``Tutorial`` target.
The end of the top-level ``CMakeLists.txt`` file will now look like the
following:

```cmake
if(USE_MYMATH)
  add_subdirectory(MathFunctions)
  list(APPEND EXTRA_LIBS MathFunctions)
  list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/MathFunctions")
endif()

# add the executable
add_executable(Tutorial tutorial.cxx)

target_link_libraries(Tutorial PUBLIC ${EXTRA_LIBS})

# add the binary tree to the search path for include files
# so that we will find TutorialConfig.h
target_include_directories(Tutorial PUBLIC
                           "${PROJECT_BINARY_DIR}"
                           ${EXTRA_INCLUDES}
                           )
```

Note the use of the variable ``EXTRA_LIBS`` to collect up any optional
libraries to later be linked into the executable. The variable
``EXTRA_INCLUDES`` is used similarly for optional header files. This is a
classic approach when dealing with many optional components, we will cover
the modern approach in the next step.

The corresponding changes to the source code are fairly straightforward.
First, in ``tutorial.cxx``, include the ``MathFunctions.h`` header if we
need it:

```cpp
#ifdef USE_MYMATH
#  include "MathFunctions.h"
#endif
```

Then, in the same file, make ``USE_MYMATH`` control which square root
function is used:

```cpp
#ifdef USE_MYMATH
  const double outputValue = mysqrt(inputValue);
#else
  const double outputValue = sqrt(inputValue);
#endif
```

Since the source code now requires ``USE_MYMATH`` we can add it to
``TutorialConfig.h.in`` with the following line:

```cpp
#cmakedefine USE_MYMATH
```

**Exercise**: Why is it important that we configure ``TutorialConfig.h.in``
after the option for ``USE_MYMATH``? What would happen if we inverted the two?

Run the [`cmake`](https://cmake.org/cmake/help/v3.23/manual/cmake.1.html#manual:cmake(1)) executable or the
[`cmake-gui`](https://cmake.org/cmake/help/v3.23/manual/cmake-gui.1.html#manual:cmake-gui(1)) to configure the project and then build it
with your chosen build tool. Then run the built Tutorial executable.

Now let's update the value of ``USE_MYMATH``. The easiest way is to use the
[`cmake-gui`](https://cmake.org/cmake/help/v3.23/manual/cmake-gui.1.html#manual:cmake-gui(1)) or [`ccmake`](https://cmake.org/cmake/help/v3.23/manual/ccmake.1.html#manual:ccmake(1)) if you're
in the terminal. Or, alternatively, if you want to change the option from the
command-line, try:

```bash
cmake ../Step2 -DUSE_MYMATH=OFF
```

Rebuild and run the tutorial again.

Which function gives better results, ``sqrt`` or ``mysqrt``?

# CMAKE generated file: DO NOT EDIT!
# Generated by "NMake Makefiles" Generator, CMake Version 3.19

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


.SUFFIXES: .hpux_make_needs_suffix_list


# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

!IF "$(OS)" == "Windows_NT"
NULL=
!ELSE
NULL=nul
!ENDIF
SHELL = cmd.exe

# The CMake executable.
CMAKE_COMMAND = "C:\Program Files\JetBrains\CLion 2021.1.2\bin\cmake\win\bin\cmake.exe"

# The command to remove a file.
RM = "C:\Program Files\JetBrains\CLion 2021.1.2\bin\cmake\win\bin\cmake.exe" -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = D:\Study\lcRevision\c

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = D:\Study\lcRevision\c\cmake-build-debug

# Include any dependencies generated for this target.
include CMakeFiles\c.dir\depend.make

# Include the progress variables for this target.
include CMakeFiles\c.dir\progress.make

# Include the compile flags for this target's objects.
include CMakeFiles\c.dir\flags.make

CMakeFiles\c.dir\main.cpp.obj: CMakeFiles\c.dir\flags.make
CMakeFiles\c.dir\main.cpp.obj: ..\main.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=D:\Study\lcRevision\c\cmake-build-debug\CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/c.dir/main.cpp.obj"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) /FoCMakeFiles\c.dir\main.cpp.obj /FdCMakeFiles\c.dir\ /FS -c D:\Study\lcRevision\c\main.cpp
<<

CMakeFiles\c.dir\main.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/c.dir/main.cpp.i"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe > CMakeFiles\c.dir\main.cpp.i @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E D:\Study\lcRevision\c\main.cpp
<<

CMakeFiles\c.dir\main.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/c.dir/main.cpp.s"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) /FoNUL /FAs /FaCMakeFiles\c.dir\main.cpp.s /c D:\Study\lcRevision\c\main.cpp
<<

CMakeFiles\c.dir\dfs.cpp.obj: CMakeFiles\c.dir\flags.make
CMakeFiles\c.dir\dfs.cpp.obj: ..\dfs.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=D:\Study\lcRevision\c\cmake-build-debug\CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/c.dir/dfs.cpp.obj"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) /FoCMakeFiles\c.dir\dfs.cpp.obj /FdCMakeFiles\c.dir\ /FS -c D:\Study\lcRevision\c\dfs.cpp
<<

CMakeFiles\c.dir\dfs.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/c.dir/dfs.cpp.i"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe > CMakeFiles\c.dir\dfs.cpp.i @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E D:\Study\lcRevision\c\dfs.cpp
<<

CMakeFiles\c.dir\dfs.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/c.dir/dfs.cpp.s"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) /FoNUL /FAs /FaCMakeFiles\c.dir\dfs.cpp.s /c D:\Study\lcRevision\c\dfs.cpp
<<

CMakeFiles\c.dir\recursive.cpp.obj: CMakeFiles\c.dir\flags.make
CMakeFiles\c.dir\recursive.cpp.obj: ..\recursive.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=D:\Study\lcRevision\c\cmake-build-debug\CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building CXX object CMakeFiles/c.dir/recursive.cpp.obj"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) /FoCMakeFiles\c.dir\recursive.cpp.obj /FdCMakeFiles\c.dir\ /FS -c D:\Study\lcRevision\c\recursive.cpp
<<

CMakeFiles\c.dir\recursive.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/c.dir/recursive.cpp.i"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe > CMakeFiles\c.dir\recursive.cpp.i @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E D:\Study\lcRevision\c\recursive.cpp
<<

CMakeFiles\c.dir\recursive.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/c.dir/recursive.cpp.s"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) /FoNUL /FAs /FaCMakeFiles\c.dir\recursive.cpp.s /c D:\Study\lcRevision\c\recursive.cpp
<<

CMakeFiles\c.dir\maths.cpp.obj: CMakeFiles\c.dir\flags.make
CMakeFiles\c.dir\maths.cpp.obj: ..\maths.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=D:\Study\lcRevision\c\cmake-build-debug\CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Building CXX object CMakeFiles/c.dir/maths.cpp.obj"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) /FoCMakeFiles\c.dir\maths.cpp.obj /FdCMakeFiles\c.dir\ /FS -c D:\Study\lcRevision\c\maths.cpp
<<

CMakeFiles\c.dir\maths.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/c.dir/maths.cpp.i"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe > CMakeFiles\c.dir\maths.cpp.i @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E D:\Study\lcRevision\c\maths.cpp
<<

CMakeFiles\c.dir\maths.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/c.dir/maths.cpp.s"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) /FoNUL /FAs /FaCMakeFiles\c.dir\maths.cpp.s /c D:\Study\lcRevision\c\maths.cpp
<<

CMakeFiles\c.dir\binarySearch.cpp.obj: CMakeFiles\c.dir\flags.make
CMakeFiles\c.dir\binarySearch.cpp.obj: ..\binarySearch.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=D:\Study\lcRevision\c\cmake-build-debug\CMakeFiles --progress-num=$(CMAKE_PROGRESS_5) "Building CXX object CMakeFiles/c.dir/binarySearch.cpp.obj"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) /FoCMakeFiles\c.dir\binarySearch.cpp.obj /FdCMakeFiles\c.dir\ /FS -c D:\Study\lcRevision\c\binarySearch.cpp
<<

CMakeFiles\c.dir\binarySearch.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/c.dir/binarySearch.cpp.i"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe > CMakeFiles\c.dir\binarySearch.cpp.i @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E D:\Study\lcRevision\c\binarySearch.cpp
<<

CMakeFiles\c.dir\binarySearch.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/c.dir/binarySearch.cpp.s"
	C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\cl.exe @<<
 /nologo /TP $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) /FoNUL /FAs /FaCMakeFiles\c.dir\binarySearch.cpp.s /c D:\Study\lcRevision\c\binarySearch.cpp
<<

# Object files for target c
c_OBJECTS = \
"CMakeFiles\c.dir\main.cpp.obj" \
"CMakeFiles\c.dir\dfs.cpp.obj" \
"CMakeFiles\c.dir\recursive.cpp.obj" \
"CMakeFiles\c.dir\maths.cpp.obj" \
"CMakeFiles\c.dir\binarySearch.cpp.obj"

# External object files for target c
c_EXTERNAL_OBJECTS =

c.exe: CMakeFiles\c.dir\main.cpp.obj
c.exe: CMakeFiles\c.dir\dfs.cpp.obj
c.exe: CMakeFiles\c.dir\recursive.cpp.obj
c.exe: CMakeFiles\c.dir\maths.cpp.obj
c.exe: CMakeFiles\c.dir\binarySearch.cpp.obj
c.exe: CMakeFiles\c.dir\build.make
c.exe: CMakeFiles\c.dir\objects1.rsp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=D:\Study\lcRevision\c\cmake-build-debug\CMakeFiles --progress-num=$(CMAKE_PROGRESS_6) "Linking CXX executable c.exe"
	"C:\Program Files\JetBrains\CLion 2021.1.2\bin\cmake\win\bin\cmake.exe" -E vs_link_exe --intdir=CMakeFiles\c.dir --rc=C:\PROGRA~2\WI3CF2~1\10\bin\100190~1.0\x64\rc.exe --mt=C:\PROGRA~2\WI3CF2~1\10\bin\100190~1.0\x64\mt.exe --manifests -- C:\PROGRA~2\MICROS~2\2019\COMMUN~1\VC\Tools\MSVC\1429~1.300\bin\Hostx64\x64\link.exe /nologo @CMakeFiles\c.dir\objects1.rsp @<<
 /out:c.exe /implib:c.lib /pdb:D:\Study\lcRevision\c\cmake-build-debug\c.pdb /version:0.0 /machine:x64 /debug /INCREMENTAL /subsystem:console  kernel32.lib user32.lib gdi32.lib winspool.lib shell32.lib ole32.lib oleaut32.lib uuid.lib comdlg32.lib advapi32.lib 
<<

# Rule to build all files generated by this target.
CMakeFiles\c.dir\build: c.exe

.PHONY : CMakeFiles\c.dir\build

CMakeFiles\c.dir\clean:
	$(CMAKE_COMMAND) -P CMakeFiles\c.dir\cmake_clean.cmake
.PHONY : CMakeFiles\c.dir\clean

CMakeFiles\c.dir\depend:
	$(CMAKE_COMMAND) -E cmake_depends "NMake Makefiles" D:\Study\lcRevision\c D:\Study\lcRevision\c D:\Study\lcRevision\c\cmake-build-debug D:\Study\lcRevision\c\cmake-build-debug D:\Study\lcRevision\c\cmake-build-debug\CMakeFiles\c.dir\DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles\c.dir\depend

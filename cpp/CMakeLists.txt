cmake_minimum_required(VERSION 3.20)
project(interview_setup LANGUAGES CXX)
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)
find_package(SQLite3 REQUIRED)
include(FetchContent)
FetchContent_Declare(nlohmann_json
    GIT_REPOSITORY https://github.com/nlohmann/json.git
    GIT_TAG v3.11.3
    GIT_SHALLOW TRUE)
FetchContent_MakeAvailable(nlohmann_json)
add_executable(setup setup.cpp)
target_link_libraries(setup PRIVATE SQLite::SQLite3 nlohmann_json::nlohmann_json)
target_compile_options(setup PRIVATE -Wall -Wextra -Werror -pedantic)
enable_testing()
add_test(NAME setup COMMAND setup)
set_tests_properties(setup PROPERTIES TIMEOUT 10)

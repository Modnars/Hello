#!/usr/bin/env bash

rm -fr ./debug/ ./release/

mkdir -p debug && cd debug
cmake -DCMAKE_BUILD_TYPE=Debug ..
cmake --build .
cd ..

mkdir  -p release && cd release
cmake -DCMAKE_BUILD_TYPE=release ..
cmake --build .
cd ..

cpack --config MultiCPackConfig.cmake

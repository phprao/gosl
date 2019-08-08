#!/bin/bash

set -e

platform='unknown'
unamestr=`uname`
if [[ "$unamestr" == 'Linux' ]]; then
   platform='linux'
elif [[ "$unamestr" == 'MINGW32_NT-6.2' ]]; then
   platform='windows'
elif [[ "$unamestr" == 'MINGW64_NT-10.0' ]]; then
   platform='windows'
elif [[ "$unamestr" == 'Darwin' ]]; then
   platform='darwin'
fi

echo "   platform = $platform"

if [[ $platform != 'darwin' ]]; then
    echo "You don't need to use this script in your platform. It's just for macOS."
    exit 1
fi

# this line will fail if the path does not exist (OK)
GCC_PATH=`ls -d /usr/local/Cellar/gcc/*`

echo "   GCC_PATH = $GCC_PATH"

if [[ -z $GCC_PATH ]]; then
    echo "cannot find /usr/local/Cellar/gcc path"
    exit 1
fi

SPECIFIC_VERSION=`basename $GCC_PATH`
VERSION="${SPECIFIC_VERSION:0:1}"

echo "   SPECIFIC_VERSION = $SPECIFIC_VERSION"
echo "   VERSION = $VERSION"

FLAGS_FILE="xautogencgoflags.go"

echo "// Copyright 2016 The Gosl Authors. All rights reserved." > $FLAGS_FILE
echo "// Use of this source code is governed by a BSD-style" >> $FLAGS_FILE
echo "// license that can be found in the LICENSE file." >> $FLAGS_FILE
echo "" >> $FLAGS_FILE
echo "// *** NOTE: this file was auto generated by all.bash ***" >> $FLAGS_FILE
echo "// ***       and should be ignored                    ***" >> $FLAGS_FILE
echo "" >> $FLAGS_FILE
echo "// +build darwin" >> $FLAGS_FILE
echo "" >> $FLAGS_FILE
echo "package qpck" >> $FLAGS_FILE
echo "" >> $FLAGS_FILE
echo "/*" >> $FLAGS_FILE
echo "#cgo darwin LDFLAGS: -L/usr/local/opt/openblas/lib -L/usr/local/Cellar/gcc/$SPECIFIC_VERSION/lib/gcc/$VERSION/ -lopenblas -lgfortran -lm" >> $FLAGS_FILE
echo "*/" >> $FLAGS_FILE
echo "import \"C\"" >> $FLAGS_FILE

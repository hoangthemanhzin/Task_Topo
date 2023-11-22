#!/bin/bash
rm -rf b5gc
mkdir -p b5gc
cd ..
cp -r common mesh nfs util pfcp nas config sbi sctp logctx docker/b5gc/
cp go.sum go.mod Makefile docker/b5gc/

#! /bin/sh

# verbose procedure
set -v

# alias for antlr jar
# JDK 11 required for latest antlr v4.12.0 tool, here is jdk-11.0.17 downloaded from Oracle source website
# Replace your jdk11's absolute path to configure it on your own machine for developing
alias java11='/Library/Java/JavaVirtualMachines/jdk-11.0.17.jdk/Contents/Home/bin/java'
alias antlr4='java11 -Xmx500M -cp "./antlr-4.12.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'

# generate antlr files
antlr4 -Dlanguage=Go -visitor -no-listener arishem.g4
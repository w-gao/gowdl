#!/bin/bash
PACKAGE='parsers'
ANTLR4_JAR='antlr-4.11.1-complete.jar'


if [ ! -f $ANTLR4_JAR ]; then
    echo "Downloading ANTLR4 tools..."
    curl "https://www.antlr.org/download/$ANTLR4_JAR" > $ANTLR4_JAR
else
    echo "ANTLR4 already exists. Skipping download."
fi


# Generate parsers
java -jar $ANTLR4_JAR -Dlanguage=Go -no-visitor -package $PACKAGE *.g4
echo "Parsers generated!"

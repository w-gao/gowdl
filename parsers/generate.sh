#!/bin/bash
ANTLR4_JAR='antlr-4.11.1-complete.jar'


if [ ! -f $ANTLR4_JAR ]; then
    echo "Downloading ANTLR4 tools..."
    curl -O "https://www.antlr.org/download/$ANTLR4_JAR"
else
    echo "ANTLR4 already exists. Skipping download."
fi


# Generate parsers
for f in v*; do
    if [ -d "$f" ]; then
        echo "Generating parsers for $f..."
        java -jar $ANTLR4_JAR -Dlanguage=Go -no-listener -visitor -package $f $f/*.g4
    fi
done

echo "Parsers generated!"

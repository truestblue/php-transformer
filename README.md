Transformer for Polyscripted PHP.

./php-transformer [-dump -phar -dump -usePhp5 -replace] [path]

-dump		dumps the AST to stdout

-phar		processes .inc and .phar files rather than copying

-test		processess .phpt files rather than copying

-usePhp5	parses php5 rather than 7

-replace	rather than creating a new directory from the given path, replaces files and directories in place so the output directory is the same but with transformed files



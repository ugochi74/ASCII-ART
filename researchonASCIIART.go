nput

A single string passed as a command-line argument
The string can contain letters, numbers, spaces, and special characters
The string can contain literal \n which should be treated as a newline

Output

The string rendered in large ASCII characters
Each character is represented by a glyph that is 8 lines tall
Glyphs are printed side by side on the same 8 lines
Each \n in the input produces a blank line between rendered blocks

The font file (banner.txt)

Contains glyphs for all printable ASCII characters (32 to 126)
Each glyph is 8 lines tall
Glyphs are separated by one blank line (so 9 lines per character block)
Characters are in sequential ASCII order, starting from space

Core logic

Read and parse the font file into a character-to-glyph mapping
Split the input on \n
For each segment, print 8 lines, each built by concatenating the matching glyph line for every character
An empty segment (from \n) prints a single blank line

Edge cases

Empty string input → output nothing
Input with only \n characters → only blank lines
Unknown characters → skip or treat as space
No argument given → exit with a usage message

Key constraints

Output must match the font file exactly — spacing, alignment, and character width all matter
The font file must be read at runtime, not hardcoded
Works only with monospace font rendering (terminal output)

That's the full picture. What would you like to tackle first?

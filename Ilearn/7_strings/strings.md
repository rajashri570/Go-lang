1. What is character encoding?

Character encoding is a set of rules for converting characters into numeric codes that computers can understand. Common character encodings include ASCII, ISO-8859-1, and UTF-8.

UTF-8 is a variable-length Unicode character encoding method that uses 1 to 4 bytes to represent a character. It is the official recommended encoding of the Unicode standard.


2. UTF-8 

Strings are stored and represented in computers through character encoding. In Go, strings use UTF-8 encoding by default, which means it can easily represent any Unicode character.

How does Go represent Unicode characters internally?

    Go uses UTF-8 to represent Unicode characters internally. Each Unicode character is encoded as one to four bytes.


3. Unicode:
The Unicode Standard is the specification of an encoding scheme for written characters and text. It is a universal standard that enables consistent encoding of multilingual text and allows text data to be interchanged internationally without conflict. The ISO standards for C and C++ refer to Information technolog

Unicode and UTF-8 in the Go programming language. Below are some common questions related to Unicode and UTF-8 in Go, along with brief answers:

  
How can I iterate over Unicode code points in a Go string?
Ans:-Use the for range loop to iterate over Unicode code points in a Go string. The loop automatically handles multi-byte characters.

What is the difference between string and []byte in Go?
Ans:- A string in Go is immutable, while a []byte is mutable. Use string for text data and []byte for binary data.

How can I check the length of a string in terms of Unicode characters?
Ans:- Use the len function to get the byte length of a string. To get the number of Unicode characters, use the utf8.RuneCountInString function.

How do I convert a Unicode character to its UTF-8 representation in Go?
Ans:- You don't need to explicitly convert a Unicode character to UTF-8 in Go. The language handles Unicode encoding automatically.

4. Byte:
A byte is the basic unit of digital information storage and processing. It is typically composed of 8 bits and can represent values from 0 to 255. The byte type is used to represent a sequence of 8 bits.

For example, in Go, a byte is an alias for the uint8 type. You often encounter bytes when working with binary data, files, or network communication.

ex: var myByte byte = 65 //here it will represent character 'A'

5. Rune:
In Go, a "rune" is used to represent a Unicode code point. A rune is an alias for the int32 type and is used when dealing with characters in the Unicode character set. Unicode characters can have values larger than 255, so a rune allows you to work with a wider range of characters, including those from different languages and symbols.

6.Ascii : ASCII, which stands for American Standard Code for Information Interchange, is a character encoding standard used for representing text and control characters in binary format so computer can opearate.
It presented by 7-bits so allows total 128 characters.

7. code point
A code point is a numerical value that represents a specific character in a character set or encoding standard, such as Unicode. In Unicode, each character is assigned a unique code point, which is usually expressed as a hexadecimal number. For example, the code point for the letter "A" in Unicode is U+0041, where "U+" indicates that it's a Unicode code point, and "0041" is the hexadecimal representation of the code point.
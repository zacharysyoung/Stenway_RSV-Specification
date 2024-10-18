# Rows of String Values (RSV) format specification

## Introduction

The RSV format is a simple binary alternative to CSV that eliminates the problem of delimiters appearing as values and thus needing special handling of the value.  The format is binary, and so cannot be handled in an RSV-unaware text editor.

This format doesn't interfere with the values, and relies on software to handle the rules of the format so that people can focus on the values.

_(This specification differs from the original in that null values have been removed to remain as close to the spirit of CSV as possible.)_

## Definition of the RSV format

 1. A value is 0 or more sequences of valid UTF-8 encode bytes, terminated by the end-of-value byte `0xFE` (EOV).  For example (spaces included for clarity):

    ```
    aaa EOV
    ```

    or an empty value:

    ```
    EOV
    ```

 2. A row is 0 or more values, terminated by the end-of-row byte `0xFF` (EOR).  For example:

    ```
    aaa EOV EOV ccc EOV EOR
    ```

    or an emtpy row:

    ```
    EOR
    ```

 3. An RSV file is 0 or more rows, allowing for an empty file.  For example, one row of three values, followed by an empty row, followed by a row of two values:

    ```
    aaa EOV EOV ccc EOV EOR zzz EOV yyy EOV EOR
    ```

The ABNF grammar [^1] appears as follows:

file = *row

row = *value EOR

value = *([UTF-8_DATA] EOV)

EOR = %xFF

EOV = %xFE

UTF-8_DATA = as per section 4 of RFC 3629 [^2]

## Normative references

[^1]: Crocker, D. and P. Overell, "Augmented BNF for Syntax Specifications: ABNF", RFC 2234, November 1997.
[^2]: Yergeau, F., "UTF-8, a transformation format of ISO 10646", RFC 3629, November 2003

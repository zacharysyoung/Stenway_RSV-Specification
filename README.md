# Rows of String Values (RSV) format specification

## Introduction

The RSV format is a simple binary alternative to CSV that eliminates the problem of delimiters appearing as values and thus needing special handling of the value.  The format is binary, and so cannot be handled in any software that is RSV-unaware; it is intended that users not manually create or edit an RSV file.  This specification differs from RFC 4180 in that it doesn't describe how RSV might be used, e.g. trying to define a header.

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

 3. A file is 0 or more rows, allowing for an empty file.  For example, one row of three values, followed by an empty row, followed by a row of two values:

    ```
    aaa EOV EOV ccc EOV EOR EOR zzz EOV yyy EOV EOR
    ```

The ABNF grammar [^1] appears as follows:

```
file = *row

row = *value EOR

value = *( [UTF-8_DATA] EOV )

EOR = %xFF

EOV = %xFE

UTF-8_DATA = as per section 4 of RFC 3629
```

The EOR and EOV byte-values were chosen because by definition they cannot appear in valid UTF-8 data [^2].

##

## Normative references

[^1]: Crocker, D. and P. Overell, "Augmented BNF for Syntax Specifications: ABNF", RFC 2234, November 1997.
[^2]: Yergeau, F., "UTF-8, a transformation format of ISO 10646", RFC 3629, November 2003

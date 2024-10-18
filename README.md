# Rows of String Values (RSV) format specification

## Introduction

The RSV format is a simple binary alternative to CSV that eliminates the problem of delimiters appearing as values and thus needing special handling of the value.  The format is binary, and so cannot be handled by applications that are RSV-unaware; it is intended that users rely on RSV-aware applications to create, edit, import, or export an RSV file.

This specification differs from [the original](https://github.com/Stenway/RSV-Specification) in that this spec does not define null values.

This specification differs from RFC 4180 [^1] in that it doesn't describe how RSV might be used, e.g. trying to define a header.


## Definition of the RSV format

 1. A value is 0 or more sequences of valid UTF-8 encode bytes, terminated by the end-of-value byte 0xFE (EOV).  For example (spaces included for clarity):

    `aaa EOV`

    or an empty value:

    `EOV`

 2. A row is 0 or more values, terminated by the end-of-row byte 0xFF (EOR).  For example:

    `aaa EOV EOV ccc EOV EOR`

    or an emtpy row:

    `EOR`

 3. A file is 0 or more rows, allowing for an empty file.  For example, one row of three values, followed by an empty row, followed by a row of two values:

    `aaa EOV EOV ccc EOV EOR EOR zzz EOV yyy EOV EOR`

The ABNF grammar [^2] appears as follows:

`file = *row`

`row = *value EOR`

`value = *( [UTF8DATA] EOV )`

`EOR = %xFF`

`EOV = %xFE`

`UTF8DATA = as per section 4 of RFC 3629`

The EOR and EOV byte-values were chosen because by definition they cannot appear in valid UTF-8 data [^3].

## A trivial implementation

To decode bytes to rows of strings:

```python
EOV = b"\xEF"
EOR = b"\xFF"


def decode(rsv: bytes) -> list[list[str]]:
    rows: list[list[str]] = []
    for _row in rsv.split(EOR)[:-1]:
        row: list[str] = []
        for _value in _row.split(EOV)[:-1]:
            row.append(_value.decode())
        rows.append(row)
    return rows


rsv_bytes = (
    b"\x41\x6c\x6c\x20\x64\x6f\x6e\x65\x21\x20\xe2\x9c\xa8\x20\xf0\x9f\x8d\xb0\x20\xe2\x9c\xa8"
    + EOV
    + b"\x48\x6f\x6f\x72\x61\x79\x21"
    + EOV
    + EOR
    + b"\x41\x6c\x6c\x20\x64\x6f\x6e\x65\x2e"
    + EOV
    + EOR
)

rsv_rows = decode(rsv_bytes)

print(rsv_rows)
```

```python
[
    ['All done! ✨ 🍰 ✨', 'Hooray!'],
    ['All done.'],
]
```

To encode rows of strings to bytes:

```python
from io import BytesIO


def encode(rows: list[list[str]]) -> bytes:
    s = BytesIO()
    for row in rows:
        for value in row:
            s.write(value.encode())
            s.write(EOV)
        s.write(EOR)
    return s.getvalue()


assert encode(rsv_rows) == rsv_bytes
assert decode(encode(rsv_rows)) == rsv_rows
```

## Normative references

[^1]: Shafranovich Y., "Common Format and MIME Type for Comma-Separated Values (CSV) Files", RFC 4180, October 2005.
[^2]: Crocker, D. and P. Overell, "Augmented BNF for Syntax Specifications: ABNF", RFC 2234, November 1997.
[^3]: Yergeau, F., "UTF-8, a transformation format of ISO 10646", RFC 3629, November 2003.

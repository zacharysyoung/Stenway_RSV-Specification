# == Decode ==

EOV = b"\xEF"
EOR = b"\xFF"

Row = list[str]
Rows = list[Row]


def decode(rsv: bytes) -> Rows:
    rows: Rows = []
    for _row in rsv.split(EOR)[:-1]:
        row: Row = []
        for _value in _row.split(EOV)[:-1]:
            row.append(_value.decode())
        rows.append(row)
    return rows


RSV_BYTES = (b""
    + b"\x61\x61\x61" + EOV + EOV + b"\x63\x63\x63" + EOV + EOR
    + EOR
    + b"\x7a\x7a\x7a" + EOV + b"\x79\x79\x79" + EOV + EOR
)  # fmt: skip

rsv_rows = decode(RSV_BYTES)

print(rsv_rows)

# == Encode ==

from io import BytesIO


def encode(rows: Rows) -> bytes:
    b = BytesIO()
    for row in rows:
        for value in row:
            b.write(value.encode())
            b.write(EOV)
        b.write(EOR)
    return b.getvalue()


assert encode(rsv_rows) == RSV_BYTES
assert decode(encode(rsv_rows)) == rsv_rows

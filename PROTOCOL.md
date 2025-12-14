# RESP Protocol Support

## Supported Types

- Simple Strings (+)
- Errors (-)
- Integers (:)
- Bulk Strings ($)
- Arrays (*)

## Example

Client:
*1\r\n$4\r\nPING\r\n

Server:
+PONG\r\n

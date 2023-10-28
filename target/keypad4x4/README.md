# Keypad 4x4

## Wiring

Uses an [Adafruit 4x4 Matrix](https://www.adafruit.com/product/3844) on Raspberry Pi Pico:

| Matrix PIN | Pi Pico       |
| ---------- | ------------- |
| 1 (`NC`)   | Not connected |
| 2 (`C1`)   | `GPIO6`       |
| 3 (`C2`)   | `GPIO7`       |
| 4 (`C3`)   | `GPIO8`       |
| 5 (`C4`)   | `GPIO9`       |
| 6 (`R1`)   | `GPIO10`      |
| 7 (`R2`)   | `GPIO11`      |
| 8 (`R3`)   | `GPIO12`      |
| 9 (`R4`)   | `GPIO13`      |
| 10 (`NC`)  | Not connected |

## Building

Using TinyGo, put the Pico info flash-mode and run:

```bash
tinygo flash -target pico ./target/keypad4x4
```
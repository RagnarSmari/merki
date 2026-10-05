# Merki

A Go library and print server for designing and printing thermal labels with ZPL.

*Merki* is Icelandic for "mark" or "sign".

> **Status: early development.** The label format and API will change. Not ready for production use yet.

## Why

Commercial label SDKs such as Neodynamic's are capable but closed and paid. Merki aims to be a free, simple alternative for the common case: describe a label as data, preview it, and print it.

ZPL is the starting point because it is publicly documented and supported by Zebra printers and by many other manufacturers.

## How it works

1. A label is described in a predefined JSON structure.
2. Variables in the label are filled in with your data at print time.
3. Merki turns the result into ZPL, a PNG preview, or a print job sent to a printer.

Because the label is plain JSON, you can build your own label editor on top of it in any language or framework.

## Status

Checked items are implemented and tested.

### Label elements

- [ ] Text
- [ ] Fonts and font sizes
- [ ] Box
- [ ] Line
- [ ] Circle
- [ ] 1D barcodes
  - [ ] GS1-128
- [ ] QR codes
  - [ ] GS1 QR
- [ ] Images
- [ ] Sizes and positions in millimetres or dots

### Output

- [ ] JSON label format
- [ ] Variable injection
- [ ] ZPL generation
- [ ] PNG preview

### Print server

- [ ] Send labels to a ZPL printer over the network
- [ ] Print queue
- [ ] Queue monitoring
- [ ] Printer status

## Out of scope

- A bundled label editor. Merki defines the label format and leaves the editor to you.

## Installation

```bash
go get github.com/RagnarSmari/merki
```

## License

MIT. See [LICENSE](LICENSE).

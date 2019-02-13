ffbmparse
=========

Go library/spec for the Firefox bookmark export.

To create a backup of your bookmarks in JSON format, hit `CTRL-SHIFT-b` and
use `Import & Backup -> Backup`.

`main.go` serves as an example, it won't do anything remotely useful.
`bookmarks/spec.go` contains the actual structure and some references.

License
=======

IANAL, but I don't believe interop or format parsing requires me to license
this under the Mozilla Public License. Happy to hear otherwise from someone
who is, but until then, consider this repository under the MIT license,
specifically the terms below:

```
Copyright 2019 chrsm

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in the
Software without restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
```


# Change Log


## Unreleased

### Added

- **log:** Logger interfaces


## 0.2.1 - 2017-05-26

### Added

- **time:** `MySQLDateTime` format string


## 0.2.0 - 2017-05-23

### Added

- **net:** `NewVirtualAddress` to create virtual `net.Addr` instances
- **net:** Pipe Listener-Dialer pair for testing client-server applications (eg. gRPC)


## 0.1.0 - 2017-05-16

### Added

- **errors:** `Recover` function to create an error from a recovered panic
- **errors:** `Handler` interface to handle errors
- **errors:** `LogHandler` implementation (tested with [sirupsen/logrus](https://github.com/sirupsen/logrus))
- **errors:** `TestHandler` implementation for testing purposes
- **errors:** `NullHandler` implementation as a fallback
- **util:** `ShutdownManager` to register and execute shutdown handlers when an application exits


## 0.0.5 - 2017-05-10

### Added

- `Must` func to panic when error is passed
- **time:** `Clock` interface and implementations to make testing with time easy


## 0.0.4 - 2017-04-11

### Added

- **strings:** `ToSpinal` util to convert a string to *spinal-case*
- **strings:** `ToTrain` util to convert a string to *Train-Case*
- **strings:** `ToCamel` util to convert a string to *CamelCase*


## 0.0.3 - 2017-04-09

### Added

- **archive/tar:** A Reader which reads a certain file from a TAR archive, optionally decompressing it
- **archive/tar:** `NewTarGzFileReader` (returns a Reader which decompresses and unarchives a .tar.gz stream and returns a file from it)


## 0.0.2 - 2017-04-03

### Changed

- Matched `str` package with `strings` in the stdlib


## 0.0.1 - 2017-03-11

### Added

- **strings:** `ToSnake` util to convert a string to *snake_case*

# Introduction
Simple App CLI-Based Command to Generate NTHASH From the Plaintext Password

# Install
Download the binary according with your OS on Release section here

# Build
Install latest Go version in your machine, and type these command to build into a binary.

For Linux/Mac
```bash
go build -o texttontlm
```

For Windows
```bash
go build -o texttontlm.exe
```

# Usage 
## MacOS / Linux
- Open your terminal and type <code>./texttontlm -p 12345678</code>

## Windows
- Open your CMD and type <code>./texttontlm.exe -p 12345678</code>

# Result
- Result : 259745CB123A52AA2E693AAACCA2DB52
- If error, it will return "something_error"

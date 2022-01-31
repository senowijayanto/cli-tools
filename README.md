This is a tool that functions to convert log files in the linux /var/log folder into text files or JSON files.

# Overview

---

These tools are made using the Golang programming language and use the [Cobra library](https://github.com/spf13/cobra) as a library to create CLI Applications.

## Installation

---

If you to use this tools, please follow this steps:

1. Clone this repository
2. Run `go mod tidy` and make sure that you activated go modules first
3. Install this application with `go install mytools`
4. Read and follow the description with running `mytools -h`

### Convert log files into PlainText or JSON

---

Run this command to convert the log files into PlainText or JSON: `mytools [path of log files] [flag of type] [type convert]`

1. To convert log files into PlainText, e.g. `mytools /var/log/nginx/error.log -t text`
2. To convert log files into JSON, e.g. `mytools /var/log/nginx/error.log -t json`
3. If flags of type is not included, it will be converted into PlainText file by default, e.g. `mytools /var/log/nginx/error.log`

### Export converted log files

---

You can also export the converted files into PlainText or JSON file.
Run this command to exported the files: `mytools [path of log files] [flag of type] [type convert] [flag of output] [output destination]`

1. To export from log files into PlainText, e.g. `mytools /var/log/nginx/error.log -t text -o /Users/JohnDoe/Desktop/nginxlog.txt`
2. To export from log files into JSON, e.g. `mytools /var/log/nginx/error.log -t json -o /Users/JohnDoe/Desktop/nginxlog.json`
3. If flag of type file is not included, it will be convert PlainText by default, e.g. `mytools /var/log/nginx/error.log -o /Users/JohnDoe/Desktop/nginxlog.txt`

# csvtojl

This Go program converts housesInput.csv to Json Line file.

## Instruction

### Step 1:

Download the source code:

`git clone https://github.com/yuchenbi87/csvtojl.git -b master`

### Step 2:

Build the executable from source code:

`cd csvtojl`

`go build -o csvtojl.exe`

### Step 3:

Run the program:

`.\csvtojl housesInput.csv housesOutput.jl`

The command take two argument:
1. **First argument**: the address of input file.
1. **Second argument**: the address of output file.

## UT

Run the command to test the code:

`go test -v`


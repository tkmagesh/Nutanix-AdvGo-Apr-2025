
## Prerequisites
    - Data Types
    - Language Constructs
        var, if else, for, switch case
    - Functions
        - Variadic functions
        - Higher Order functions
        - Anonymous functions
    - OO Programming
        - Structs
            - Struct Composition
        - Methods
        - Abstraction using Interfaces
    - Error Handling
    - Panic & Recovery
    - Modules & Package
    - Concurrency
        - Goroutines
        - *WaitGroups
        - Channels
        - Streaming
        - *Signaling
        - Context
    
## Repo
- https://github.com/tkmagesh/Nutanix-AdvGo-Apr-2025

## Concurrency

### WaitGroup
    - Semaphore based counter
    - Capable of blocking the execution of a (any) function until the counter becomes 0

## GRPC
- Alternative to HTTP based restful services
- Apt for microservices communication
- Uses HTTP/2
- Communication Patterns
    - Request & Response
    - Server Streaming
    - Client Streaming
    - Bidirectional Streaming
- Protocol Buffers for serialization
    - Allows for sharing the schema well in advance
    - Enables ONLY the data (without any annotation) for communication
- Multilanguage support
    - Go
    - .Net
    - JS
    - Java
    - C++
### Steps
    - Define data contracts (protobuf)
    - Define service contract (protobuf)
        - Define operation contracts (protobuf)
    - Generate the Proxy & Stub
    - Implement the service and host it using the Stub
    - Create a client that communicates to the sevice using the Proxy

### Tools Installation 
    1. Protocol Buffers Compiler (protoc tool)
        Windows:
            Download the file, extract and keep in a folder (PATH) accessble through the command line
            https://github.com/protocolbuffers/protobuf/releases/download/v24.4/protoc-24.4-win64.zip
        Mac:
            brew install protobuf

        Verification:
            protoc --version

    2. Go plugins (installed in the GOPATH/bin folder)
        go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

        Verification:
            the binaries (protoc-gen-go, protoc-gen-go-grpc) must be present in $GOPATH/bin folder




## Feedback Survey
- https://www.surveymonkey.com/r/KNYGHTZ
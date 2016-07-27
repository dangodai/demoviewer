demoviewer
===================

demoviewer is a tool for managing and viewing details about TF2 demo files. 

### Contributing
I'm not likely to make many updates to this myself, but feel free to make pull requests. I'll accept and release any acceptable changes or fixes. 

### Building
Make sure you have Go configured properly ($GOPATH)

    go get github.com/dangodai/demoviewer
    
Follow the [instructions](https://github.com/visualfc/goqt/blob/master/doc/install.md) for setting up visualfc/goqt

#### Linux
From with the demoviewer folder

    ./build
This will compile the program to $GOPATH/bin/demoviewer. Make sure you have copied the .so files from goqt.



#### Windows
From with the demoviewer folder

    build.bat
This will compile the program to $GOPATH/bin/demoviewer. Make sure you have copied the .dll files from goqt.

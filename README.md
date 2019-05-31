# darkreigngo

This is an attempt to make the Qmegas HTML5 version of Dark Reign (https://github.com/Qmegas/Dark-Reign---HTML5-Version) independant.

In order to start the game you need to clone this repo. 
It also downloads the html5 version as submodule.

```
$ git clone --resursive https://github.com/siredmar/darkreigngo.git
```

Build it using
```
$ make
```

and run it by
```
./darkreigngo
```

# Future Development?

This can be a a step in the direction of a new Dark Reign engine.
I need to get in contact with Qmegas and ask for some sort of cooperation.
Networking between clients could be done using Go (gRPC, ...). 

The UI would need to communicate with the Go backend and synchronize all steps with the other clients.

There are several Go frameworks out there (e.g. https://github.com/asticode/go-astilectron) that does allow writing programs in Go using a Web based frontend. So this might be the next step - not at least when trying to make a communication working between the Go backend and the game frontend.



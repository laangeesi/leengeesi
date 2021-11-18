=========================
==== IPLIST PROGRAM ====
=========================

This "Net.go" program is created in Go(lang).

The objective of the this program is to list information about active network interfaces on the host.

Go Program Overview:

The Go program has a very simple syntax sequence that starts with a main package, followed by fmt package
then func action and finally print action.

1) Main package owns all of the code within the program. 

2) Import package ensures Go program only loads the requested package(s)only, for faster results.

3) The func action executes the written code within "func main" after it first runs "Main,fmt."    

4) Finally, the fmt.Println prints the results of the code that func main generated as an output.

=======================
NET.GO PROGRAM DETAILS:
=======================

1) I first ensure "package main" is delcared as the first line of the code, followed by spaces.

2) Then, I imported package "net" which allows us to list the network interfaces information.

3) Then, I imported package "fmt" which allows us to format the func main code and then prints output.

4) Then, I created for-range loop for "iplist" variable which holds interfaces values that we need.

5) And finally, we used the fmt.Println() to call and print the value of the fmt.Println inside the func main. 

============================
GIT TO COMMIT NET.GO PROJECT
============================

After I created the NET.GO program, I used my existing github.com:laangeesi/leengeesi repo to commit.

DETAILS:

On my IONOS Cloud Linux Server, I was able to do the following:

1) git init //Initialized my local directory for git to track changes

2) git clone https://github.com/laangeesi/leengeesi.git //cloned my github repo

3) git add * // added my files "Net.go and README.MD"

4) git commit -m "Dear Toni, this is my go ip project as requested. I got busy but wanted to complete this same day."

5) git commit -a //to ensure changes and added files are commited.

6) git push origin main //Sent the project to my main repo on github.


============
ATTRIBUTIONS
===========
1) https://www.golangprograms.com/go-language/interface.html
2) https://pkg.go.dev/net#LookupHost
3) And some googlingx


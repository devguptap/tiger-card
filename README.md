# tiger-card
This aplication processes the trip data given as an input and returns the total fare applicable throught the journey.

# Go installation and Setup
- Install the latest version of go using link : https://golang.org/doc/install
- Set Go Binary path in PATH variable. In linux and OS X use command : 
```sh
export PATH="$PATH:/usr/local/go/bin"
```
- For Windows OS use link : https://www.computerhope.com/issues/ch000549.htm
- Run command "go env" and check the value of GOPATH variable: (GOPATH="/root/go")
- Run below command to set GO111MODULE. This is require to support the go module integration.
```sh
go env -w GO111MODULE=on
```
- Check go installation and go variables using below command: 
```sh
go version
go env
```

# Unit test cases
Unit test cases are present inside package unittest.
To run the using test please use below command.
```sh
go test -v ./unittest/
```

The sample output for the Unit test is as below:
```sh
2021-03-14T01:40:19.754+0530    INFO    unittest/main_test.go:44        Starting test cases
=== RUN   TestTigerCard

  Given a test case with id : Test1
    When run the test case
          Then the expected result should be : 120 .


1 total assertion


  Given a test case with id : Test2
    When run the test case
      Then the expected result should be : 700 .


2 total assertions

--- PASS: TestTigerCard (0.00s)
PASS
2021-03-14T01:40:19.757+0530    INFO    unittest/main_test.go:46        Test case execution done. Exit code is : 0
ok      tiger-card/unittest     0.462s
```

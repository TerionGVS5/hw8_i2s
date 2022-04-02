# hw8_i2s

interface2struct function

i2s - interface to struct. A function that populates struct values from map[string]interface{} and the like - what happens when json is unpacked into interface{} (see json/dynamic.go for an example)

Reflection task.

Despite some sophistication at first glance, reflection is used very often. Understanding how it works and how you work with it will be very useful in the future.

Implementation takes 80-100 lines of code

Of the data types, it is enough to provide those that are in the test.

Run go test -v

Write the code in the i2s.go file

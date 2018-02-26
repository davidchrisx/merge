# merge

#input example

file a
````
a
x
b
````
file b
````
a
b
c
````
output
````
  a
- x
  b
+ c
````

file a
````
g
a
c
````
file b
````
a
g
c
a
t
````
output
````
+ a
  g
+ c
  a
+ t
- c
````
